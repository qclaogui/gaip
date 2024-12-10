// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package service

import (
	"context"
	"errors"
	"fmt"
	"io/fs"
	"mime"
	"net"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	lg "github.com/grafana/dskit/log"
	"github.com/grafana/dskit/middleware"
	"github.com/grafana/dskit/signals"
	_ "github.com/grafana/pyroscope-go/godeltaprof/http/pprof" // anonymous import to get godelatprof handlers registered
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/qclaogui/gaip/third_party"
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/experimental"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/mem"
)

// Listen on the named network
const (
	// DefaultNetwork  the host resolves to multiple IP addresses,
	// Dial will try each IP address in order until one succeeds
	DefaultNetwork = "tcp"
)

// SignalHandler used by Server.
type SignalHandler interface {
	// Loop Starts the signals handler. This method is blocking, and returns only after signal is received,
	// or "Stop" is called.
	Loop()

	// Stop blocked "Loop" method.
	Stop()
}

type Server struct {
	cfg     Config
	handler SignalHandler

	httpListener net.Listener
	HTTPServer   *http.Server

	grpcListener net.Listener
	GRPCServer   *grpc.Server

	Router     *mux.Router
	Log        log.Logger
	Gatherer   prometheus.Gatherer
	Registerer prometheus.Registerer
}

// NewServer makes a new Server. It will panic if the metrics cannot be registered.
func NewServer(cfg Config) (*Server, error) {
	metrics := NewServerMetrics(cfg)
	return newServer(cfg, metrics)
}

func newServer(cfg Config, metrics *Metrics) (*Server, error) {
	// If user doesn't supply a logging implementation, by default instantiate go-kit.
	logger := cfg.Log
	if logger == nil {
		logger = lg.NewGoKitWithLevel(cfg.LogLevel, cfg.LogFormat)
	}

	gatherer := cfg.Gatherer
	if gatherer == nil {
		gatherer = prometheus.DefaultGatherer
	}

	// Setup Router
	var router *mux.Router
	if cfg.Router != nil {
		router = cfg.Router
	} else {
		router = mux.NewRouter()
	}

	if cfg.RegisterInstrumentation {
		RegisterInstrumentationWithGatherer(router, gatherer)
	}

	if cfg.RegisterOpenAPI {
		RegisterOpenAPI(router)
	}

	// Setup HTTP Server
	httpListener, httpServer, err := newEndpointREST(cfg, router, metrics, logger)
	if err != nil {
		return nil, err
	}

	// Setup gRPC Server
	grpcListener, grpcServer, err := newEndpointGRPC(cfg, router, metrics, logger)
	if err != nil {
		return nil, err
	}

	_ = level.Info(logger).Log("msg", "server listening on addresses", "http", httpListener.Addr(), "grpc", grpcListener.Addr())

	handler := cfg.SignalHandler
	if handler == nil {
		handler = signals.NewHandler(logger)
	}

	return &Server{
		cfg:     cfg,
		handler: handler,

		httpListener: httpListener,
		HTTPServer:   httpServer,

		grpcListener: grpcListener,
		GRPCServer:   grpcServer,

		Router:     router,
		Log:        logger,
		Registerer: cfg.registererOrDefault(),
		Gatherer:   gatherer,
	}, nil
}

// newEndpointREST HTTP REST
func newEndpointREST(cfg Config, router *mux.Router, metrics *Metrics, logger log.Logger) (net.Listener, *http.Server, error) {
	network := cfg.HTTPListenNetwork
	if network == "" {
		network = DefaultNetwork
	}

	// Setup listeners first, so we can fail early if the port is in use.
	httpListener, err := net.Listen(network, net.JoinHostPort(cfg.HTTPListenAddress, strconv.Itoa(cfg.HTTPListenPort)))
	if err != nil {
		return nil, nil, err
	}

	httpListener = middleware.CountingListener(httpListener, metrics.TCPConnections.WithLabelValues("http"))

	if cfg.HTTPLogClosedConnectionsWithoutResponse {
		httpListener = middleware.NewZeroResponseListener(httpListener, level.Warn(logger))
	}

	metrics.TCPConnectionsLimit.WithLabelValues("http").Set(float64(cfg.HTTPConnLimit))
	if cfg.HTTPConnLimit > 0 {
		httpListener = netutil.LimitListener(httpListener, cfg.HTTPConnLimit)
	}

	httpMiddleware, err := BuildHTTPMiddleware(cfg, router, metrics, logger)
	if err != nil {
		return nil, nil, fmt.Errorf("error building http middleware: %w", err)
	}

	httpServer := &http.Server{
		ReadTimeout:       cfg.HTTPServerReadTimeout,
		ReadHeaderTimeout: cfg.HTTPServerReadHeaderTimeout,
		WriteTimeout:      cfg.HTTPServerWriteTimeout,
		IdleTimeout:       cfg.HTTPServerIdleTimeout,
		Handler:           middleware.Merge(httpMiddleware...).Wrap(router),
	}

	return httpListener, httpServer, nil
}

// newEndpointGRPC grpc
func newEndpointGRPC(cfg Config, router *mux.Router, metrics *Metrics, logger log.Logger) (net.Listener, *grpc.Server, error) {
	network := cfg.GRPCListenNetwork
	if network == "" {
		network = DefaultNetwork
	}

	grpcListener, err := net.Listen(network, net.JoinHostPort(cfg.GRPCListenAddress, strconv.Itoa(cfg.GRPCListenPort)))
	if err != nil {
		return nil, nil, err
	}
	grpcListener = middleware.CountingListener(grpcListener, metrics.TCPConnections.WithLabelValues("grpc"))

	metrics.TCPConnectionsLimit.WithLabelValues("grpc").Set(float64(cfg.GRPCConnLimit))
	if cfg.GRPCConnLimit > 0 {
		grpcListener = netutil.LimitListener(grpcListener, cfg.GRPCConnLimit)
	}

	grpcServerLog := middleware.GRPCServerLog{
		Log:                      logger,
		WithRequest:              !cfg.ExcludeRequestInLog,
		DisableRequestSuccessLog: cfg.DisableRequestSuccessLog,
	}

	var reportGRPCStatusesOptions []middleware.InstrumentationOption
	if cfg.ReportGRPCCodesInInstrumentationLabel {
		reportGRPCStatusesOptions = []middleware.InstrumentationOption{middleware.ReportGRPCStatusOption}
	}

	grpcMiddleware := []grpc.UnaryServerInterceptor{
		grpcServerLog.UnaryServerInterceptor,
		otgrpc.OpenTracingServerInterceptor(opentracing.GlobalTracer()),
		middleware.HTTPGRPCTracingInterceptor(router), // This must appear after the OpenTracingServerInterceptor.
		middleware.UnaryServerInstrumentInterceptor(metrics.RequestDuration, reportGRPCStatusesOptions...),
	}
	grpcMiddleware = append(grpcMiddleware, cfg.GRPCMiddleware...)

	grpcStreamMiddleware := []grpc.StreamServerInterceptor{
		grpcServerLog.StreamServerInterceptor,
		otgrpc.OpenTracingStreamServerInterceptor(opentracing.GlobalTracer()),
		middleware.StreamServerInstrumentInterceptor(metrics.RequestDuration, reportGRPCStatusesOptions...),
	}
	grpcStreamMiddleware = append(grpcStreamMiddleware, cfg.GRPCStreamMiddleware...)

	grpcKeepAliveOptions := keepalive.ServerParameters{
		MaxConnectionIdle:     cfg.GRPCServerMaxConnectionIdle,
		MaxConnectionAge:      cfg.GRPCServerMaxConnectionAge,
		MaxConnectionAgeGrace: cfg.GRPCServerMaxConnectionAgeGrace,
		Time:                  cfg.GRPCServerTime,
		Timeout:               cfg.GRPCServerTimeout,
	}

	grpcKeepAliveEnforcementPolicy := keepalive.EnforcementPolicy{
		MinTime:             cfg.GRPCServerMinTimeBetweenPings,
		PermitWithoutStream: cfg.GRPCServerPingWithoutStreamAllowed,
	}

	grpcOptions := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(grpcMiddleware...),
		grpc.ChainStreamInterceptor(grpcStreamMiddleware...),
		grpc.KeepaliveParams(grpcKeepAliveOptions),
		grpc.KeepaliveEnforcementPolicy(grpcKeepAliveEnforcementPolicy),
		grpc.MaxRecvMsgSize(cfg.GRPCServerMaxRecvMsgSize),
		grpc.MaxSendMsgSize(cfg.GRPCServerMaxSendMsgSize),
		grpc.MaxConcurrentStreams(uint32(cfg.GRPCServerMaxConcurrentStreams)),
		grpc.NumStreamWorkers(uint32(cfg.GRPCServerNumWorkers)),
	}

	if cfg.GRPCServerStatsTrackingEnabled {
		grpcOptions = append(grpcOptions,
			grpc.StatsHandler(middleware.NewStatsHandler(
				metrics.ReceivedMessageSize,
				metrics.SentMessageSize,
				metrics.InflightRequests,
			)),
		)
	}

	if cfg.GRPCServerRecvBufferPoolsEnabled {
		if cfg.GRPCServerStatsTrackingEnabled {
			return nil, nil, fmt.Errorf("grpc_server_stats_tracking_enabled must be set to false if grpc_server_recv_buffer_pools_enabled is true")
		}
		grpcOptions = append(grpcOptions, experimental.BufferPool(mem.DefaultBufferPool()))
	}

	grpcOptions = append(grpcOptions, cfg.GRPCOptions...)
	grpcServer := grpc.NewServer(grpcOptions...)

	return grpcListener, grpcServer, nil
}

// RegisterInstrumentationWithGatherer on the given router.
func RegisterInstrumentationWithGatherer(router *mux.Router, gatherer prometheus.Gatherer) {
	router.Handle("/metrics", promhttp.HandlerFor(gatherer, promhttp.HandlerOpts{
		EnableOpenMetrics: true,
	}))
	router.PathPrefix("/debug/pprof").Handler(http.DefaultServeMux)
}

// RegisterOpenAPI on the given router.
func RegisterOpenAPI(router *mux.Router) {
	router.Handle("/openapi", getOpenAPIHandler())
}

// getOpenAPIHandler serves an OpenAPI UI.
// Adapted from https://github.com/philips/grpc-gateway-example/blob/a269bcb5931ca92be0ceae6130ac27ae89582ecc/cmd/serve.go#L63
func getOpenAPIHandler() http.Handler {
	_ = mime.AddExtensionType(".svg", "image/svg+xml")
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "gen/openapiv2")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}
	return http.FileServer(http.FS(subFS))
}

// HTTPListenAddr exposes `net.Addr` that `Server` is listening to for Router connections.
func (s *Server) HTTPListenAddr() net.Addr {
	return s.httpListener.Addr()
}

// GRPCListenAddr exposes `net.Addr` that `Server` is listening to for GRPCServer connections.
func (s *Server) GRPCListenAddr() net.Addr {
	return s.grpcListener.Addr()
}

func BuildHTTPMiddleware(cfg Config, router *mux.Router, metrics *Metrics, logger log.Logger) ([]middleware.Interface, error) {
	sourceIPs, err := middleware.NewSourceIPs(cfg.LogSourceIPsHeader, cfg.LogSourceIPsRegex, cfg.LogSourceIPsFull)
	if err != nil {
		return nil, fmt.Errorf("error setting up source IP extraction: %v", err)
	}

	logSourceIPs := sourceIPs
	if !cfg.LogSourceIPs {
		// We always include the source IPs for traces,
		// but only want to log them in the middleware if that is enabled.
		logSourceIPs = nil
	}

	defaultLogMiddleware := middleware.NewLogMiddleware(logger, cfg.LogRequestHeaders, cfg.LogRequestAtInfoLevel, logSourceIPs, strings.Split(cfg.LogRequestExcludeHeadersList, ","))
	defaultLogMiddleware.DisableRequestSuccessLog = cfg.DisableRequestSuccessLog

	defaultHTTPMiddleware := []middleware.Interface{
		middleware.RouteInjector{
			RouteMatcher: router,
		},
		middleware.Tracer{
			SourceIPs: sourceIPs,
		},
		defaultLogMiddleware,
		middleware.Instrument{
			Duration:         metrics.RequestDuration,
			RequestBodySize:  metrics.ReceivedMessageSize,
			ResponseBodySize: metrics.SentMessageSize,
			InflightRequests: metrics.InflightRequests,
		},
	}
	var httpMiddleware []middleware.Interface
	if cfg.DoNotAddDefaultHTTPMiddleware {
		httpMiddleware = cfg.HTTPMiddleware
	} else {
		httpMiddleware = append(defaultHTTPMiddleware, cfg.HTTPMiddleware...)
	}

	return httpMiddleware, nil
}

// Run Run the server; blocks until SIGTERM (if signal handling is enabled), an error is received, or Stop() is called.
func (s *Server) Run() error {
	errChan := make(chan error, 1)

	// Wait for a signal
	go func() {
		s.handler.Loop()
		select {
		case errChan <- nil:
		default:
		}
	}()

	// Setup gRPC server
	go func() {
		err := s.GRPCServer.Serve(s.grpcListener)
		handleGRPCError(err, errChan)
	}()

	// Setup HTTP server
	go func() {
		err := s.HTTPServer.Serve(s.httpListener)
		handleHTTPError(err, errChan)
	}()

	return <-errChan
}

// handleHTTPError consolidates HTTP Server error handling by sending
// any error to errChan except for http.ErrServerClosed which is ignored.
func handleHTTPError(err error, errChan chan error) {
	if errors.Is(err, http.ErrServerClosed) {
		err = nil
	}

	select {
	case errChan <- err:
	default:
	}
}

// handleGRPCError consolidates GRPC Server error handling by sending
// any error to errChan except for grpc.ErrServerStopped which is ignored.
func handleGRPCError(err error, errChan chan error) {
	if errors.Is(err, grpc.ErrServerStopped) {
		err = nil
	}

	select {
	case errChan <- err:
	default:
	}
}

// Stop unblocks Run().
func (s *Server) Stop() {
	s.handler.Stop()
}

// Shutdown the server, gracefully.  Should be defer after NewServer().
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ServerGracefulShutdownTimeout)
	defer cancel() // releases resources if httpServer.Shutdown completes before timeout elapses

	_ = s.HTTPServer.Shutdown(ctx)
	s.GRPCServer.GracefulStop()
}
