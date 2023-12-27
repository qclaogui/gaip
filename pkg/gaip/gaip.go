// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package gaip

import (
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/felixge/fgprof"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/qclaogui/gaip/pkg/protocol/grpc/interceptors"
	"github.com/qclaogui/gaip/pkg/service"
	"github.com/qclaogui/gaip/pkg/vault"
	lg "github.com/qclaogui/gaip/tools/log"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v3"
)

type Gaip struct {
	Cfg        Config
	Server     *service.Server
	Vault      *vault.Vault
	Registerer prometheus.Registerer
}

func (g *Gaip) RegisterFgprof() {
	g.RegisterRoute("/debug/fgprof", fgprof.Handler(), false, "GET")
}

// Bootstrap makes a new Gaip.
func Bootstrap(cfg Config, reg prometheus.Registerer) (*Gaip, error) {
	if cfg.PrintConfig {
		if err := yaml.NewEncoder(os.Stdout).Encode(&cfg); err != nil {
			fmt.Println("Error encoding config:", err)
		}
		os.Exit(0)
	}

	setUpGoRuntimeMetrics(cfg, reg)

	// Inject the registerer in the Server config too.
	cfg.ServerCfg.Registerer = reg

	g := &Gaip{
		Cfg:        cfg,
		Registerer: reg,
	}

	g.Cfg.ServerCfg.Router = mux.NewRouter()

	// TODO(qc) config gRPC and REST
	//g.Cfg.ServerCfg.HTTPMiddleware = nil

	//g.Cfg.ServerCfg.GRPCMiddleware=nil
	//g.Cfg.ServerCfg.GRPCStreamMiddleware=nil

	g.Cfg.ServerCfg.GRPCOptions = interceptors.RegisterGRPCServerOption()

	if err := g.initServices(); err != nil {
		return nil, err
	}

	// Register reflection service on gRPC server.
	// Enable reflection to allow clients to query the server's services
	reflection.Register(g.Server.GRPCServer)

	return g, nil
}

func (g *Gaip) initServices() error {
	var err error
	if g.Server, err = service.NewServer(g.Cfg.ServerCfg); err != nil {
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

	if err = g.initRouteGuide(); err != nil {
		return err
	}

	if err = g.initTodo(); err != nil {
		return err
	}

	if err = g.initVault(); err != nil {
		return err
	}

	return nil
}

// Run gRPC server and HTTP gateway
func (g *Gaip) Run() error {

	// before starting servers, register /healthz handler.
	//g.Server.Router.Path("/healthz").HandlerFunc(g.healthzHandler())
	g.RegisterRoute("/healthz", g.healthzHandler(), false, http.MethodGet)

	g.RegisterFgprof()

	//g.API.RegisterServiceMapHandler(http.HandlerFunc(g.servicesHandler))

	// Initialize tracing and handle the tracer provider shutdown
	stopTracing := interceptors.InitTracing()
	defer stopTracing()

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
		_ = level.Info(lg.Logger).Log("msg", "[healthz] received request")
		w.WriteHeader(http.StatusOK)
	}
}

// RegisterRoute registers a single route enforcing HTTP methods. A single
// route is expected to be specific about which HTTP methods are supported.
func (g *Gaip) RegisterRoute(path string, handler http.Handler, auth bool, method string, methods ...string) {
	methods = append([]string{method}, methods...)
	_ = level.Debug(lg.Logger).Log("msg", "gaip: registering route", "methods", strings.Join(methods, ","), "path", path, "auth", auth)
	g.newRoute(path, handler, false, auth, methods...)
}

func (g *Gaip) RegisterRoutesWithPrefix(path string, handler http.Handler, auth bool, method string, methods ...string) {
	methods = append([]string{method}, methods...)
	_ = level.Debug(lg.Logger).Log("msg", "gaip: registering route", "methods", strings.Join(methods, ","), "path", path, "auth", auth)
	g.newRoute(path, handler, true, auth, methods...)
}

func (g *Gaip) newRoute(path string, handler http.Handler, isPrefix, auth bool, methods ...string) {
	var route *mux.Route
	//if auth {
	//	//handler = g.AuthMiddleware.Wrap(handler)
	//}

	_ = auth

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
