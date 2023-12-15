// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package service

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"math"
	"mime"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/gorilla/mux"
	lg "github.com/grafana/dskit/log"
	"github.com/grafana/dskit/middleware"
	"github.com/grafana/dskit/signals"
	otgrpc "github.com/opentracing-contrib/go-grpc"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/qclaogui/gaip/third_party"
	"golang.org/x/net/netutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

// Listen on the named network
const (
	// DefaultNetwork  the host resolves to multiple IP addresses,
	// Dial will try each IP address in order until one succeeds
	DefaultNetwork = "tcp"
	// NetworkTCPV4 for IPV4 only
	NetworkTCPV4 = "tcp4"
)

// Config for a Server
type Config struct {
	MetricsNamespace string `yaml:"-"`

	// Set to > 1 to add native histograms to requestDuration.
	// See documentation for NativeHistogramBucketFactor in
	// https://pkg.go.dev/github.com/prometheus/client_golang/prometheus#HistogramOpts
	// for details. A generally useful value is 1.1.
	MetricsNativeHistogramFactor float64 `yaml:"-"`

	// HTTP configs
	HTTPListenNetwork            string                 `yaml:"http_listen_network"`
	HTTPListenAddress            string                 `yaml:"http_listen_address"`
	HTTPListenPort               int                    `yaml:"http_listen_port"`
	HTTPConnLimit                int                    `yaml:"http_listen_conn_limit"`
	HTTPServerReadTimeout        time.Duration          `yaml:"http_server_read_timeout"`
	HTTPServerReadHeaderTimeout  time.Duration          `yaml:"http_server_read_header_timeout"`
	HTTPServerWriteTimeout       time.Duration          `yaml:"http_server_write_timeout"`
	HTTPServerIdleTimeout        time.Duration          `yaml:"http_server_idle_timeout"`
	HTTPMiddleware               []middleware.Interface `yaml:"-"`
	DisableDefaultHTTPMiddleware bool                   `yaml:"-"`

	// gRPC configs
	GRPCListenNetwork                     string              `yaml:"grpc_listen_network"`
	GRPCListenAddress                     string              `yaml:"grpc_listen_address"`
	GRPCListenPort                        int                 `yaml:"grpc_listen_port"`
	GRPCConnLimit                         int                 `yaml:"grpc_listen_conn_limit"`
	GRPCServerMaxRecvMsgSize              int                 `yaml:"grpc_server_max_recv_msg_size"`
	GRPCServerMaxSendMsgSize              int                 `yaml:"grpc_server_max_send_msg_size"`
	GRPCServerMaxConcurrentStreams        uint                `yaml:"grpc_server_max_concurrent_streams"`
	GRPCServerMaxConnectionIdle           time.Duration       `yaml:"grpc_server_max_connection_idle"`
	GRPCServerMaxConnectionAge            time.Duration       `yaml:"grpc_server_max_connection_age"`
	GRPCServerMaxConnectionAgeGrace       time.Duration       `yaml:"grpc_server_max_connection_age_grace"`
	GRPCServerTime                        time.Duration       `yaml:"grpc_server_keepalive_time"`
	GRPCServerTimeout                     time.Duration       `yaml:"grpc_server_keepalive_timeout"`
	GRPCServerMinTimeBetweenPings         time.Duration       `yaml:"grpc_server_min_time_between_pings"`
	GRPCServerPingWithoutStreamAllowed    bool                `yaml:"grpc_server_ping_without_stream_allowed"`
	GRPCServerNumWorkers                  int                 `yaml:"grpc_server_num_workers"`
	GRPCOptions                           []grpc.ServerOption `yaml:"-"`
	ReportGRPCCodesInInstrumentationLabel bool                `yaml:"report_grpc_codes_in_instrumentation_label_enabled"`

	GRPCMiddleware       []grpc.UnaryServerInterceptor  `yaml:"-"`
	GRPCStreamMiddleware []grpc.StreamServerInterceptor `yaml:"-"`

	Router                  *mux.Router `yaml:"-"`
	RegisterInstrumentation bool        `yaml:"register_instrumentation"`
	RegisterOpenAPI         bool        `yaml:"register_open_api"`

	HTTPLogClosedConnectionsWithoutResponse bool `yaml:"http_log_closed_connections_without_response_enabled"`

	LogFormat                    string     `yaml:"log_format"`
	LogLevel                     lg.Level   `yaml:"log_level"`
	Log                          log.Logger `yaml:"-"`
	ExcludeRequestInLog          bool       `yaml:"-"`
	DisableRequestSuccessLog     bool       `yaml:"-"`
	LogSourceIPs                 bool       `yaml:"log_source_ips_enabled"`
	LogSourceIPsHeader           string     `yaml:"log_source_ips_header"`
	LogSourceIPsRegex            string     `yaml:"log_source_ips_regex"`
	LogRequestHeaders            bool       `yaml:"log_request_headers"`
	LogRequestAtInfoLevel        bool       `yaml:"log_request_at_info_level_enabled"`
	LogRequestExcludeHeadersList string     `yaml:"log_request_exclude_headers_list"`

	ServerGracefulShutdownTimeout time.Duration `yaml:"graceful_shutdown_timeout"`

	// If not set, default Prometheus Registerer is used.
	Registerer prometheus.Registerer
	// If not set, default Prometheus Gatherer is used.
	Gatherer prometheus.Gatherer
}

var infinity = time.Duration(math.MaxInt64)

// RegisterFlags adds the flags required to config this to the given FlagSet
func (cfg *Config) RegisterFlags(fs *flag.FlagSet) {
	fs.StringVar(&cfg.HTTPListenNetwork, "server.http-listen-network", DefaultNetwork, "Router server listen network, default tcp")
	fs.StringVar(&cfg.HTTPListenAddress, "server.http-listen-address", "", "Router server listen address.")
	fs.IntVar(&cfg.HTTPListenPort, "server.http-listen-port", 80, "Router server listen port.")
	fs.IntVar(&cfg.HTTPConnLimit, "server.http-conn-limit", 0, "Maximum number of simultaneous http connections, <=0 to disable")
	fs.DurationVar(&cfg.HTTPServerReadTimeout, "server.http-read-timeout", 30*time.Second, "Read timeout for entire Router request, including headers and body.")
	fs.DurationVar(&cfg.HTTPServerReadHeaderTimeout, "server.http-read-header-timeout", 0, "Read timeout for Router request headers. If set to 0, value of -server.http-read-timeout is used.")
	fs.DurationVar(&cfg.HTTPServerWriteTimeout, "server.http-write-timeout", 30*time.Second, "Write timeout for Router server")
	fs.DurationVar(&cfg.HTTPServerIdleTimeout, "server.http-idle-timeout", 120*time.Second, "Idle timeout for Router server")

	fs.StringVar(&cfg.GRPCListenNetwork, "server.grpc-listen-network", DefaultNetwork, "gRPC server listen network")
	fs.StringVar(&cfg.GRPCListenAddress, "server.grpc-listen-address", "", "gRPC server listen address.")
	fs.IntVar(&cfg.GRPCListenPort, "server.grpc-listen-port", 9095, "gRPC server listen port.")
	fs.IntVar(&cfg.GRPCConnLimit, "server.grpc-conn-limit", 0, "Maximum number of simultaneous grpc connections, <=0 to disable")
	fs.IntVar(&cfg.GRPCServerMaxRecvMsgSize, "server.grpc-max-recv-msg-size-bytes", 4*1024*1024, "Limit on the size of a gRPC message this server can receive (bytes).")
	fs.IntVar(&cfg.GRPCServerMaxSendMsgSize, "server.grpc-max-send-msg-size-bytes", 4*1024*1024, "Limit on the size of a gRPC message this server can send (bytes).")
	fs.UintVar(&cfg.GRPCServerMaxConcurrentStreams, "server.grpc-max-concurrent-streams", 100, "Limit on the number of concurrent streams for gRPC calls per client connection (0 = unlimited)")
	fs.DurationVar(&cfg.GRPCServerMaxConnectionIdle, "server.grpc.keepalive.max-connection-idle", infinity, "The duration after which an idle connection should be closed. Default: infinity")
	fs.DurationVar(&cfg.GRPCServerMaxConnectionAge, "server.grpc.keepalive.max-connection-age", infinity, "The duration for the maximum amount of time a connection may exist before it will be closed. Default: infinity")
	fs.DurationVar(&cfg.GRPCServerMaxConnectionAgeGrace, "server.grpc.keepalive.max-connection-age-grace", infinity, "An additive period after max-connection-age after which the connection will be forcibly closed. Default: infinity")
	fs.DurationVar(&cfg.GRPCServerTime, "server.grpc.keepalive.time", time.Hour*2, "Duration after which a keepalive probe is sent in case of no activity over the connection., Default: 2h")
	fs.DurationVar(&cfg.GRPCServerTimeout, "server.grpc.keepalive.timeout", time.Second*20, "After having pinged for keepalive check, the duration after which an idle connection should be closed, Default: 20s")
	fs.DurationVar(&cfg.GRPCServerMinTimeBetweenPings, "server.grpc.keepalive.min-time-between-pings", 5*time.Minute, "Minimum amount of time a client should wait before sending a keepalive ping. If client sends keepalive ping more often, server will send GOAWAY and close the connection.")
	fs.BoolVar(&cfg.GRPCServerPingWithoutStreamAllowed, "server.grpc.keepalive.ping-without-stream-allowed", false, "If true, server allows keepalive pings even when there are no active streams(RPCs). If false, and client sends ping when there are no active streams, server will send GOAWAY and close the connection.")
	fs.IntVar(&cfg.GRPCServerNumWorkers, "server.grpc.num-workers", 0, "If non-zero, configures the amount of GRPC server workers used to serve the requests.")
	fs.BoolVar(&cfg.ReportGRPCCodesInInstrumentationLabel, "server.report-grpc-codes-in-instrumentation-label-enabled", false, "If set to true, gRPC statuses will be reported in instrumentation labels with their string representations. Otherwise, they will be reported as \"error\".")

	fs.BoolVar(&cfg.RegisterInstrumentation, "server.register-instrumentation", true, "Register the intrumentation handlers (/metrics etc).")
	fs.BoolVar(&cfg.RegisterOpenAPI, "server.register-open-api", true, "Register and handlers OpenAPI UI (/openapi).")

	fs.BoolVar(&cfg.HTTPLogClosedConnectionsWithoutResponse, "server.http-log-closed-connections-without-response-enabled", false, "Log closed connections that did not receive any response, most likely because client didn't send any request within timeout.")

	fs.StringVar(&cfg.LogFormat, "log.format", lg.LogfmtFormat, "Output log messages in the given format. Valid formats: [logfmt, json]")
	cfg.LogLevel.RegisterFlags(fs)
	fs.BoolVar(&cfg.LogSourceIPs, "server.log-source-ips-enabled", false, "Optionally log the source IPs.")
	fs.StringVar(&cfg.LogSourceIPsHeader, "server.log-source-ips-header", "", "Header field storing the source IPs. Only used if server.log-source-ips-enabled is true. If not set the default Forwarded, X-Real-IP and X-Forwarded-For headers are used")
	fs.StringVar(&cfg.LogSourceIPsRegex, "server.log-source-ips-regex", "", "Regex for matching the source IPs. Only used if server.log-source-ips-enabled is true. If not set the default Forwarded, X-Real-IP and X-Forwarded-For headers are used")
	fs.BoolVar(&cfg.LogRequestHeaders, "server.log-request-headers", false, "Optionally log request headers.")
	fs.BoolVar(&cfg.LogRequestAtInfoLevel, "server.log-request-at-info-level-enabled", false, "Optionally log requests at info level instead of debug level. Applies to request headers as well if server.log-request-headers is enabled.")
	fs.StringVar(&cfg.LogRequestExcludeHeadersList, "server.log-request-headers-exclude-list", "", "Comma separated list of headers to exclude from loggin. Only used if server.log-request-headers is true.")

	fs.DurationVar(&cfg.ServerGracefulShutdownTimeout, "server.graceful-shutdown-timeout", 30*time.Second, "Timeout for graceful shutdowns")
}

func (cfg *Config) registererOrDefault() prometheus.Registerer {
	// If user doesn't supply a Registerer/gatherer, use Prometheus' by default.
	if cfg.Registerer != nil {
		return cfg.Registerer
	}
	return prometheus.DefaultRegisterer
}

type Server struct {
	cfg        Config
	Gatherer   prometheus.Gatherer
	Registerer prometheus.Registerer

	httpListener net.Listener
	HTTPServer   *http.Server

	grpcListener net.Listener
	GRPCServer   *grpc.Server

	Router *mux.Router
	Log    log.Logger
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

	network := cfg.HTTPListenNetwork
	if network == "" {
		network = DefaultNetwork
	}

	// Setup listeners first, so we can fail early if the port is in use.
	httpListener, err := net.Listen(network, net.JoinHostPort(cfg.HTTPListenAddress, strconv.Itoa(cfg.HTTPListenPort)))
	if err != nil {
		return nil, err
	}

	httpListener = middleware.CountingListener(httpListener, metrics.TCPConnections.WithLabelValues("http"))

	if cfg.HTTPLogClosedConnectionsWithoutResponse {
		httpListener = middleware.NewZeroResponseListener(httpListener, level.Warn(logger))
	}

	metrics.TCPConnectionsLimit.WithLabelValues("http").Set(float64(cfg.HTTPConnLimit))
	if cfg.HTTPConnLimit > 0 {
		httpListener = netutil.LimitListener(httpListener, cfg.HTTPConnLimit)
	}

	network = cfg.GRPCListenNetwork
	if network == "" {
		network = DefaultNetwork
	}

	grpcListener, err := net.Listen(network, net.JoinHostPort(cfg.GRPCListenAddress, strconv.Itoa(cfg.GRPCListenPort)))
	if err != nil {
		return nil, err
	}

	grpcListener = middleware.CountingListener(grpcListener, metrics.TCPConnections.WithLabelValues("grpc"))

	metrics.TCPConnectionsLimit.WithLabelValues("grpc").Set(float64(cfg.GRPCConnLimit))
	if cfg.GRPCConnLimit > 0 {
		grpcListener = netutil.LimitListener(grpcListener, cfg.GRPCConnLimit)
	}

	_ = level.Info(logger).Log("msg", "server listening on addresses", "http", httpListener.Addr(), "grpc", grpcListener.Addr())

	// Setup HTTP server
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

	sourceIPs, err := middleware.NewSourceIPs(cfg.LogSourceIPsHeader, cfg.LogSourceIPsRegex)
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
		middleware.Tracer{
			RouteMatcher: router,
			SourceIPs:    sourceIPs,
		},
		defaultLogMiddleware,
		middleware.Instrument{
			RouteMatcher:     router,
			Duration:         metrics.RequestDuration,
			RequestBodySize:  metrics.ReceivedMessageSize,
			ResponseBodySize: metrics.SentMessageSize,
			InflightRequests: metrics.InflightRequests,
		},
	}

	var httpMiddleware []middleware.Interface
	if cfg.DisableDefaultHTTPMiddleware {
		httpMiddleware = cfg.HTTPMiddleware
	} else {
		httpMiddleware = append(defaultHTTPMiddleware, cfg.HTTPMiddleware...)
	}

	httpServer := &http.Server{
		ReadTimeout:       cfg.HTTPServerReadTimeout,
		ReadHeaderTimeout: cfg.HTTPServerReadHeaderTimeout,
		WriteTimeout:      cfg.HTTPServerWriteTimeout,
		IdleTimeout:       cfg.HTTPServerIdleTimeout,
		Handler:           middleware.Merge(httpMiddleware...).Wrap(router),
	}

	// Setup gRPC server
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

	grpcOptions = append(grpcOptions,
		grpc.StatsHandler(middleware.NewStatsHandler(
			metrics.ReceivedMessageSize,
			metrics.SentMessageSize,
			metrics.InflightRequests,
		)),
	)

	grpcOptions = append(grpcOptions, cfg.GRPCOptions...)

	grpcServer := grpc.NewServer(grpcOptions...)

	return &Server{
		cfg:          cfg,
		httpListener: httpListener,
		HTTPServer:   httpServer,
		grpcListener: grpcListener,
		GRPCServer:   grpcServer,
		Router:       router,
		Log:          logger,
		Registerer:   cfg.registererOrDefault(),
		Gatherer:     gatherer,
	}, nil
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

// Run Run the server; blocks until SIGTERM (if signal handling is enabled), an error is received, or Stop() is called.
func (s *Server) Run() error {
	errChan := make(chan error, 1)

	// Wait for a signal
	go func() {
		signals.SignalHandlerLoop(s.Log)

		select {
		case errChan <- nil:
		default:
		}
	}()

	// Setup HTTP server
	go func() {
		handleHTTPError(s.HTTPServer.Serve(s.httpListener), errChan)
	}()

	// Setup gRPC server
	go func() {
		handleGRPCError(s.GRPCServer.Serve(s.grpcListener), errChan)
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

// Shutdown the server, gracefully.  Should be defer after NewServer().
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), s.cfg.ServerGracefulShutdownTimeout)
	defer cancel() // releases resources if httpServer.Shutdown completes before timeout elapses

	_ = s.HTTPServer.Shutdown(ctx)
	s.GRPCServer.GracefulStop()
}
