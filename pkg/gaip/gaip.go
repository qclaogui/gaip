// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/felixge/fgprof"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/qclaogui/gaip/genproto/todo/apiv1/todopb"
	"github.com/qclaogui/gaip/pkg/protocol/grpc/interceptors"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/service/bookstore"
	"github.com/qclaogui/gaip/pkg/service/library"
	"github.com/qclaogui/gaip/pkg/service/project"
	"github.com/qclaogui/gaip/pkg/service/routeguide"
	"github.com/qclaogui/gaip/pkg/service/todo"
	"github.com/qclaogui/gaip/pkg/vault"
	lg "github.com/qclaogui/gaip/tools/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"

	// mysql driver
	_ "github.com/go-sql-driver/mysql"
)

type Gaip struct {
	Cfg        Config
	Server     *service.Server
	Vault      *vault.Vault
	Registerer prometheus.Registerer
}

func (g *Gaip) RegisterAPI() {
	g.RegisterRoute("/debug/fgprof", fgprof.Handler(), false, true, "GET")
}

// initVault init Vault
func (g *Gaip) initVault() error {
	if !g.Cfg.VaultCfg.Enabled {
		return nil
	}

	v, err := vault.New(g.Cfg.VaultCfg)
	if err != nil {
		return err
	}
	g.Vault = v

	// Update Configs - KVStore
	//g.Cfg.MemberlistKV.TCPTransport.TLS.Reader = g.VaultCfg

	// Update Configs - GRPCServer Clients
	//g.Cfg.Worker.GRPCClientConfig.TLS.Reader = g.VaultCfg

	return nil
}

// New makes a new Gaip.
func New(cfg Config, reg prometheus.Registerer) (*Gaip, error) {
	if cfg.PrintConfig {
		if err := yaml.NewEncoder(os.Stdout).Encode(&cfg); err != nil {
			fmt.Println("Error encoding config:", err)
		}
		os.Exit(0)
	}

	setUpGoRuntimeMetrics(cfg, reg)

	// Inject the registerer in the Server config too.
	cfg.ServerCfg.Registerer = reg

	app := &Gaip{
		Cfg:        cfg,
		Registerer: reg,
	}

	app.Cfg.ServerCfg.Router = mux.NewRouter()

	// TODO(qc) config gRPC and REST

	if err := app.initServices(); err != nil {
		return nil, err
	}

	return app, nil
}

func (g *Gaip) initServices() error {
	var err error
	if g.Server, err = service.NewServer(g.Cfg.ServerCfg); err != nil {
		return err
	}

	if err = todo.New(g.Cfg.TodoCfg, g.Server); err != nil {
		return err
	}

	if err = routeguide.New(g.Cfg.RouteGuideCfg, g.Server); err != nil {
		return err
	}

	if err = bookstore.New(g.Cfg.BookstoreCfg, g.Server); err != nil {
		return err
	}

	if err = library.New(g.Cfg.LibraryCfg, g.Server); err != nil {
		return err
	}

	if err = project.New(g.Cfg.ProjectCfg, g.Server); err != nil {
		return err
	}

	if err = g.initVault(); err != nil {
		return err
	}

	return nil
}

// Run gRPC server and HTTP gateway
func (g *Gaip) Run() error {
	ctx := context.Background()

	// before starting servers, register /healthz handler.
	g.Server.Router.Path("/healthz").Handler(g.healthzHandler())

	g.RegisterAPI()

	//g.API.RegisterServiceMapHandler(http.HandlerFunc(g.servicesHandler))

	// Initialize tracing and handle the tracer provider shutdown
	stopTracing := interceptors.InitTracing()
	defer stopTracing()

	// Start the REST server in goroutine
	//err = rest.RunRESTServer(ctx, g.Server)
	//lg.CheckFatal("running REST server", err)

	// Runs HTTP/REST gateway
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	gwmux := runtime.NewServeMux()
	// Register the gRPC server's handler with the Router gwmux
	err := todopb.RegisterToDoServiceHandlerFromEndpoint(ctx, gwmux, g.Server.GRPCListenAddr().String(), opts)
	if err != nil {
		slog.Error("failed to start Router gateway", "error", err)
		return err
	}

	// Set up the REST server and handle requests by proxying them to the gRPC server
	g.Server.Router.PathPrefix("/v1").Handler(gwmux)

	// Register reflection service on gRPC server.
	// Enable reflection to allow clients to query the server's services
	reflection.Register(g.Server.GRPCServer)

	return g.Server.Run()
}

func setUpGoRuntimeMetrics(cfg Config, reg prometheus.Registerer) {
	rules := []collectors.GoRuntimeMetricsRule{
		// Enable the mutex wait time metric.
		{Matcher: regexp.MustCompile(`^/sync/mutex/wait/total:seconds$`)},
	}

	if cfg.EnableGoRuntimeMetrics {
		// Enable all available runtime metrics.
		rules = append(rules, collectors.MetricsAll)
	}

	// Unregister the default Go collector...
	reg.Unregister(collectors.NewGoCollector())

	// ...and replace it with our own that adds our extra rules.
	reg.MustRegister(collectors.NewGoCollector(
		collectors.WithGoCollectorRuntimeMetrics(rules...),
	))
}

// healthzHandler for a liveness probe.
func (g *Gaip) healthzHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}
}

// RegisterRoute registers a single route enforcing HTTP methods. A single
// route is expected to be specific about which HTTP methods are supported.
func (g *Gaip) RegisterRoute(path string, handler http.Handler, auth, gzipEnabled bool, method string, methods ...string) {
	methods = append([]string{method}, methods...)
	_ = level.Debug(lg.Logger).Log("msg", "api: registering route", "methods", strings.Join(methods, ","), "path", path, "auth", auth, "gzip", gzipEnabled)
	g.newRoute(path, handler, false, auth, gzipEnabled, methods...)
}

func (g *Gaip) RegisterRoutesWithPrefix(path string, handler http.Handler, auth, gzipEnabled bool, method string, methods ...string) {
	methods = append([]string{method}, methods...)
	_ = level.Debug(lg.Logger).Log("msg", "api: registering route", "methods", strings.Join(methods, ","), "path", path, "auth", auth, "gzip", gzipEnabled)
	g.newRoute(path, handler, true, auth, gzipEnabled, methods...)
}

func (g *Gaip) newRoute(path string, handler http.Handler, isPrefix, auth, gzip bool, methods ...string) {
	var route *mux.Route
	//if auth {
	//	//handler = g.AuthMiddleware.Wrap(handler)
	//}
	//if gzip {
	//	//handler = gziphandler.GzipHandler(handler)
	//}

	_ = auth
	_ = gzip

	if isPrefix {
		route = g.Server.Router.PathPrefix(path)
	} else {
		route = g.Server.Router.Path(path)
	}

	if len(methods) > 0 {
		route = route.Methods(methods...)
	}

	route.Handler(handler)
}
