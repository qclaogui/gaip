// Copyright © Weifeng Wang <qclaogui@gmail.com>
//
// Licensed under the Apache License 2.0.

package service

import (
	"flag"
	"math"
	"time"

	gokitlog "github.com/go-kit/log"
	"github.com/gorilla/mux"
	"github.com/grafana/dskit/log"
	"github.com/grafana/dskit/middleware"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
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
	HTTPListenNetwork             string                 `yaml:"http_listen_network"`
	HTTPListenAddress             string                 `yaml:"http_listen_address"`
	HTTPListenPort                int                    `yaml:"http_listen_port"`
	HTTPConnLimit                 int                    `yaml:"http_listen_conn_limit"`
	HTTPServerReadTimeout         time.Duration          `yaml:"http_server_read_timeout"`
	HTTPServerReadHeaderTimeout   time.Duration          `yaml:"http_server_read_header_timeout"`
	HTTPServerWriteTimeout        time.Duration          `yaml:"http_server_write_timeout"`
	HTTPServerIdleTimeout         time.Duration          `yaml:"http_server_idle_timeout"`
	HTTPMiddleware                []middleware.Interface `yaml:"-"`
	DoNotAddDefaultHTTPMiddleware bool                   `yaml:"-"`

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
	GRPCServerStatsTrackingEnabled        bool                `yaml:"grpc_server_stats_tracking_enabled"`
	GRPCServerRecvBufferPoolsEnabled      bool                `yaml:"grpc_server_recv_buffer_pools_enabled"`
	GRPCOptions                           []grpc.ServerOption `yaml:"-"`
	ReportGRPCCodesInInstrumentationLabel bool                `yaml:"report_grpc_codes_in_instrumentation_label_enabled"`

	GRPCMiddleware       []grpc.UnaryServerInterceptor  `yaml:"-"`
	GRPCStreamMiddleware []grpc.StreamServerInterceptor `yaml:"-"`

	Router                  *mux.Router `yaml:"-"`
	RegisterInstrumentation bool        `yaml:"register_instrumentation"`
	RegisterOpenAPI         bool        `yaml:"register_open_api"`

	HTTPLogClosedConnectionsWithoutResponse bool `yaml:"http_log_closed_connections_without_response_enabled"`

	LogFormat                    string          `yaml:"log_format"`
	LogLevel                     log.Level       `yaml:"log_level"`
	Log                          gokitlog.Logger `yaml:"-"`
	ExcludeRequestInLog          bool            `yaml:"-"`
	DisableRequestSuccessLog     bool            `yaml:"-"`
	LogSourceIPs                 bool            `yaml:"log_source_ips_enabled"`
	LogSourceIPsFull             bool            `yaml:"log_source_ips_full"`
	LogSourceIPsHeader           string          `yaml:"log_source_ips_header"`
	LogSourceIPsRegex            string          `yaml:"log_source_ips_regex"`
	LogRequestHeaders            bool            `yaml:"log_request_headers"`
	LogRequestAtInfoLevel        bool            `yaml:"log_request_at_info_level_enabled"`
	LogRequestExcludeHeadersList string          `yaml:"log_request_exclude_headers_list"`

	ServerGracefulShutdownTimeout time.Duration `yaml:"graceful_shutdown_timeout"`

	// If not set, default signal handler is used.
	SignalHandler SignalHandler `yaml:"-"`

	// If not set, default Prometheus Registerer is used.
	Registerer prometheus.Registerer
	// If not set, default Prometheus Gatherer is used.
	Gatherer prometheus.Gatherer

	Throughput Throughput `yaml:"-"`
}
type Throughput struct {
	LatencyCutoff time.Duration `yaml:"throughput_latency_cutoff"`
	Unit          string        `yaml:"throughput_unit"`
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
	fs.BoolVar(&cfg.GRPCServerStatsTrackingEnabled, "server.grpc.stats-tracking-enabled", true, "If true, the request_message_bytes, response_message_bytes, and inflight_requests metrics will be tracked. Enabling this option prevents the use of memory pools for parsing gRPC request bodies and may lead to more memory allocations.")
	fs.BoolVar(&cfg.GRPCServerRecvBufferPoolsEnabled, "server.grpc.recv-buffer-pools-enabled", false, "If true, gGPC's buffer pools will be used to handle incoming requests. Enabling this feature can reduce memory allocation, but also requires disabling GRPC server stats tracking by setting `server.grpc.stats-tracking-enabled=false`. This is an experimental gRPC feature, so it might be removed in a future version of the gRPC library.")
	fs.BoolVar(&cfg.ReportGRPCCodesInInstrumentationLabel, "server.report-grpc-codes-in-instrumentation-label-enabled", false, "If set to true, gRPC statuses will be reported in instrumentation labels with their string representations. Otherwise, they will be reported as \"error\".")

	fs.BoolVar(&cfg.RegisterInstrumentation, "server.register-instrumentation", true, "Register the intrumentation handlers (/metrics etc).")
	fs.BoolVar(&cfg.RegisterOpenAPI, "server.register-open-api", true, "Register and handlers OpenAPI UI (/openapi).")

	fs.BoolVar(&cfg.HTTPLogClosedConnectionsWithoutResponse, "server.http-log-closed-connections-without-response-enabled", false, "Log closed connections that did not receive any response, most likely because client didn't send any request within timeout.")

	fs.StringVar(&cfg.LogFormat, "log.format", log.LogfmtFormat, "Output log messages in the given format. Valid formats: [logfmt, json]")
	cfg.LogLevel.RegisterFlags(fs)
	fs.BoolVar(&cfg.LogSourceIPs, "server.log-source-ips-enabled", true, "Optionally log the source IPs.")
	fs.BoolVar(&cfg.LogSourceIPsFull, "server.log-source-ips-full", true, "Log all source IPs instead of only the originating one. Only used if server.log-source-ips-enabled is true")
	fs.StringVar(&cfg.LogSourceIPsHeader, "server.log-source-ips-header", "", "Header field storing the source IPs. Only used if server.log-source-ips-enabled is true. If not set the default Forwarded, X-Real-IP and X-Forwarded-For headers are used")
	fs.StringVar(&cfg.LogSourceIPsRegex, "server.log-source-ips-regex", "", "Regex for matching the source IPs. Only used if server.log-source-ips-enabled is true. If not set the default Forwarded, X-Real-IP and X-Forwarded-For headers are used")
	fs.BoolVar(&cfg.LogRequestHeaders, "server.log-request-headers", false, "Optionally log request headers.")
	fs.BoolVar(&cfg.LogRequestAtInfoLevel, "server.log-request-at-info-level-enabled", false, "Optionally log requests at info level instead of debug level. Applies to request headers as well if server.log-request-headers is enabled.")
	fs.StringVar(&cfg.LogRequestExcludeHeadersList, "server.log-request-headers-exclude-list", "", "Comma separated list of headers to exclude from loggin. Only used if server.log-request-headers is true.")
	fs.DurationVar(&cfg.Throughput.LatencyCutoff, "server.throughput.latency-cutoff", 0, "Requests taking over the cutoff are be observed to measure throughput. Server-Timing header is used with specified unit as the indicator, for example 'Server-Timing: unit;val=8.2'. If set to 0, the throughput is not calculated.")
	fs.StringVar(&cfg.Throughput.Unit, "server.throughput.unit", "samples_processed", "Unit of the server throughput metric, for example 'processed_bytes' or 'samples_processed'. Observed values are gathered from the 'Server-Timing' header with the 'val' key. If set, it is appended to the request_server_throughput metric name.")
	fs.DurationVar(&cfg.ServerGracefulShutdownTimeout, "server.graceful-shutdown-timeout", 30*time.Second, "Timeout for graceful shutdowns")
}

func (cfg *Config) registererOrDefault() prometheus.Registerer {
	// If user doesn't supply a Registerer/gatherer, use Prometheus' by default.
	if cfg.Registerer != nil {
		return cfg.Registerer
	}
	return prometheus.DefaultRegisterer
}
