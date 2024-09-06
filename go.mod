module github.com/qclaogui/gaip

go 1.22.5

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.34.2-20240717164558-a6c49f84cc0f.2
	cloud.google.com/go/longrunning v0.5.12
	dagger.io/dagger v0.12.4
	entgo.io/ent v0.14.0
	github.com/blang/semver v3.5.1+incompatible
	github.com/bufbuild/protovalidate-go v0.6.3
	github.com/containerd/platforms v0.2.1
	github.com/dave/jennifer v1.7.0
	github.com/felixge/fgprof v0.9.4
	github.com/ghodss/yaml v1.0.0
	github.com/go-kit/log v0.2.1
	github.com/go-sql-driver/mysql v1.8.1
	github.com/golang/protobuf v1.5.4
	github.com/google/generative-ai-go v0.17.0
	github.com/google/go-cmp v0.6.0
	github.com/google/uuid v1.6.0
	github.com/googleapis/gax-go/v2 v2.13.0
	github.com/gorilla/mux v1.8.1
	github.com/grafana/dskit v0.0.0-20240807162847-859371e7a400
	github.com/grafana/pyroscope-go/godeltaprof v0.1.7
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.1.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.21.0
	github.com/hashicorp/vault/api v1.14.0
	github.com/iancoleman/strcase v0.3.0
	github.com/jackc/pgx/v5 v5.6.0
	github.com/mattn/go-sqlite3 v1.14.19
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.34.1
	github.com/opentracing-contrib/go-grpc v0.0.0-20240724223109-9dec25a38fa8
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.20.3
	github.com/spf13/cobra v1.8.1
	github.com/stretchr/testify v1.9.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.53.0
	go.opentelemetry.io/otel v1.28.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.28.0
	go.opentelemetry.io/otel/sdk v1.28.0
	golang.org/x/net v0.28.0
	google.golang.org/api v0.191.0
	google.golang.org/genproto v0.0.0-20240808171019-573a1156607a
	google.golang.org/genproto/googleapis/api v0.0.0-20240808171019-573a1156607a
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240808171019-573a1156607a
	google.golang.org/grpc v1.65.0
	google.golang.org/protobuf v1.34.2
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	ariga.io/atlas v0.26.0 // indirect
	cloud.google.com/go v0.115.0 // indirect
	cloud.google.com/go/ai v0.8.2 // indirect
	cloud.google.com/go/auth v0.8.0 // indirect
	cloud.google.com/go/auth/oauth2adapt v0.2.4 // indirect
	cloud.google.com/go/compute/metadata v0.5.0 // indirect
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/99designs/gqlgen v0.17.49 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Khan/genqlient v0.7.0 // indirect
	github.com/adrg/xdg v0.5.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/units v0.0.0-20240626203959-61d1e3462e30 // indirect
	github.com/antlr4-go/antlr/v4 v4.13.1 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cenkalti/backoff/v4 v4.3.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/containerd/log v0.1.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/go-jose/go-jose/v4 v4.0.4 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.21.0 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gogo/status v1.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/cel-go v0.21.0 // indirect
	github.com/google/pprof v0.0.0-20240727154555-813a5fbdbec8 // indirect
	github.com/google/s2a-go v0.1.8 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.2 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.7 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.8 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.6 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.21.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0 // indirect
	github.com/opentracing-contrib/go-stdlib v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/prometheus/client_model v0.6.1 // indirect
	github.com/prometheus/common v0.55.0 // indirect
	github.com/prometheus/procfs v0.15.1 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/sosodev/duration v1.3.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stoewer/go-strcase v1.3.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/vektah/gqlparser/v2 v2.5.16 // indirect
	github.com/zclconf/go-cty v1.15.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.53.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.0.0-20240809064039-69e235829efd // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.4.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.28.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.28.0 // indirect
	go.opentelemetry.io/otel/log v0.4.0 // indirect
	go.opentelemetry.io/otel/metric v1.28.0 // indirect
	go.opentelemetry.io/otel/sdk/log v0.4.0 // indirect
	go.opentelemetry.io/otel/trace v1.28.0 // indirect
	go.opentelemetry.io/proto/otlp v1.3.1 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.26.0 // indirect
	golang.org/x/exp v0.0.0-20240808152545-0cdaa3abc0fa // indirect
	golang.org/x/mod v0.20.0 // indirect
	golang.org/x/oauth2 v0.22.0 // indirect
	golang.org/x/sync v0.8.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	golang.org/x/time v0.6.0 // indirect
	golang.org/x/tools v0.24.0 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
