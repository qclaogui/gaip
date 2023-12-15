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
	"github.com/qclaogui/gaip/genproto/bookstore/apiv1alpha1/bookstorepb"
	"github.com/qclaogui/gaip/genproto/library/apiv1/librarypb"
	"github.com/qclaogui/gaip/genproto/project/apiv1/projectpb"
	"github.com/qclaogui/gaip/genproto/routeguide/apiv1/routeguidepb"
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
	Registerer prometheus.Registerer

	Server *service.Server

	TodoSrv       todo.Service
	RouteGuideSrv routeguide.Service
	BookstoreSrv  bookstore.Service
	LibrarySrv    library.Service
	ProjectSrv    project.Service

	Vault *vault.Vault
}

func (g *Gaip) RegisterAPI() {
	g.RegisterRoute("/debug/fgprof", fgprof.Handler(), false, true, "GET")
}

// initVault init Vault
func (g *Gaip) initVault() error {
	if !g.Cfg.Vault.Enabled {
		return nil
	}

	v, err := vault.New(g.Cfg.Vault)
	if err != nil {
		return err
	}
	g.Vault = v

	// Update Configs - KVStore
	//g.Cfg.MemberlistKV.TCPTransport.TLS.Reader = g.Vault

	// Update Configs - GRPCServer Clients
	//g.Cfg.Worker.GRPCClientConfig.TLS.Reader = g.Vault

	return nil
}

func (g *Gaip) initTodo() error {
	var err error
	if g.TodoSrv, err = todo.NewServiceServer(g.Cfg.Todo, lg.Logger, g.Registerer); err != nil {
		return err
	}

	todopb.RegisterToDoServiceServer(g.Server.GRPCServer, g.TodoSrv)

	// Expose HTTP endpoints.
	g.RegisterRoute("/todo/healthz", g.healthzHandler(), false, true, "GET", "POST")
	return nil
}

func (g *Gaip) initRouteguide() error {
	var err error
	if g.RouteGuideSrv, err = routeguide.NewRouteGuideService(g.Cfg.RouteGuide, lg.Logger, g.Registerer); err != nil {
		return err
	}

	routeguidepb.RegisterRouteGuideServiceServer(g.Server.GRPCServer, g.RouteGuideSrv)
	// Expose HTTP endpoints.
	g.RegisterRoute("/routeguide/healthz", g.healthzHandler(), false, true, "GET", "POST")
	return nil
}

func (g *Gaip) initBookstore() error {
	var err error
	if g.BookstoreSrv, err = bookstore.NewBookstoreServer(g.Cfg.Bookstore, lg.Logger, g.Registerer); err != nil {
		return err
	}

	bookstorepb.RegisterBookstoreServiceServer(g.Server.GRPCServer, g.BookstoreSrv)

	// Expose HTTP endpoints.
	g.RegisterRoute("/bookstore/healthz", g.healthzHandler(), false, true, "GET", "POST")
	return nil
}

func (g *Gaip) initLibrary() error {
	var err error
	if g.LibrarySrv, err = library.NewLibraryService(g.Cfg.Library, lg.Logger, g.Registerer); err != nil {
		return err
	}

	librarypb.RegisterLibraryServiceServer(g.Server.GRPCServer, g.LibrarySrv)

	// Expose HTTP endpoints.
	g.RegisterRoute("/library/healthz", g.healthzHandler(), false, true, "GET", "POST")

	return nil
}

func (g *Gaip) initProject() error {
	var err error
	if g.ProjectSrv, err = project.NewProjectService(g.Cfg.Project, lg.Logger, g.Registerer); err != nil {
		return err
	}

	projectpb.RegisterProjectServiceServer(g.Server.GRPCServer, g.ProjectSrv)
	//projectpb.RegisterIdentityServiceServer(g.Server.GRPCServer, nil)
	//projectpb.RegisterEchoServiceServer(g.Server.GRPCServer, nil)

	// Expose HTTP endpoints.
	g.RegisterRoute("/project/healthz", g.healthzHandler(), false, true, "GET", "POST")
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
	cfg.Server.Registerer = reg

	app := &Gaip{
		Cfg:        cfg,
		Registerer: reg,
	}

	app.Cfg.Server.Router = mux.NewRouter()

	if err := app.setupServices(); err != nil {
		return nil, err
	}

	return app, nil
}

func (g *Gaip) setupServices() error {
	var err error
	if g.Server, err = service.NewServer(g.Cfg.Server); err != nil {
		return err
	}

	if err = g.initTodo(); err != nil {
		return err
	}

	if err = g.initRouteguide(); err != nil {
		return err
	}

	if err = g.initBookstore(); err != nil {
		return err
	}

	if err = g.initLibrary(); err != nil {
		return err
	}

	if err = g.initProject(); err != nil {
		return err
	}

	if err = g.initVault(); err != nil {
		return err
	}

	return nil
}

// Bootstrap bootstrap gRPC server and HTTP gateway
func (g *Gaip) Bootstrap() error {
	ctx := context.Background()

	// before starting servers, register /healthz handler.
	g.Server.Router.Path("/healthz").Handler(g.healthzHandler())

	//t.ServiceMap, err = t.ModuleManager.InitModuleServices(t.Cfg.Target...)

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
