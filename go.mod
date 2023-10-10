module github.com/qclaogui/golang-api-server

go 1.20

require (
	cloud.google.com/go/longrunning v0.5.1
	dagger.io/dagger v0.8.7
	entgo.io/ent v0.12.4
	github.com/blang/semver v3.5.1+incompatible
	github.com/containerd/containerd v1.7.6
	github.com/dave/jennifer v1.7.0
	github.com/go-kit/log v0.2.1
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang/protobuf v1.5.3
	github.com/google/uuid v1.3.1
	github.com/googleapis/gax-go/v2 v2.12.0
	github.com/grafana/dskit v0.0.0-20231006094724-ad2fd7e7931e
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.18.0
	github.com/hashicorp/vault/api v1.10.0
	github.com/jackc/pgx/v5 v5.4.3
	github.com/mattn/go-sqlite3 v1.14.16
	github.com/onsi/ginkgo v1.16.5
	github.com/onsi/gomega v1.27.10
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.17.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.45.0
	go.opentelemetry.io/otel v1.19.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.19.0
	go.opentelemetry.io/otel/sdk v1.19.0
	google.golang.org/api v0.146.0
	google.golang.org/genproto/googleapis/api v0.0.0-20231002182017-d307bd883b97
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
	gopkg.in/DATA-DOG/go-sqlmock.v1 v1.3.0
	gopkg.in/yaml.v3 v3.0.1
)

require (
	ariga.io/atlas v0.14.2 // indirect
	cloud.google.com/go v0.110.8 // indirect
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	github.com/99designs/gqlgen v0.17.39 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/Khan/genqlient v0.6.0 // indirect
	github.com/Microsoft/hcsshim v0.11.1 // indirect
	github.com/adrg/xdg v0.4.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/alecthomas/units v0.0.0-20211218093645-b94a6e3cc137 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cenkalti/backoff/v3 v3.2.2 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/coreos/go-systemd/v22 v22.5.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/facette/natsort v0.0.0-20181210072756-2cd4dd1e2dcb // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/go-jose/go-jose/v3 v3.0.0 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.2.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/go-redis/redis/v8 v8.11.5 // indirect
	github.com/gogo/googleapis v1.4.1 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/gogo/status v1.1.1 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/pprof v0.0.0-20210720184732-4bb14d4b1be1 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.1 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/grafana/gomemcache v0.0.0-20230920151301-bbcdec66d3d1 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.4 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/parseutil v0.1.7 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/hashicorp/go-sockaddr v1.0.5 // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.7 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.18.1 // indirect
	github.com/iancoleman/strcase v0.3.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jpillora/backoff v1.0.0 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/miekg/dns v1.1.56 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mwitkow/go-conntrack v0.0.0-20190716064945-2f068394615f // indirect
	github.com/nxadm/tail v1.4.8 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.1.0-rc5 // indirect
	github.com/opentracing-contrib/go-grpc v0.0.0-20210225150812-73cb765af46e // indirect
	github.com/opentracing-contrib/go-stdlib v1.0.0 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/prometheus/client_model v0.5.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/exporter-toolkit v0.10.1-0.20230714054209-2f4150c63f97 // indirect
	github.com/prometheus/procfs v0.12.0 // indirect
	github.com/rogpeppe/go-internal v1.11.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	github.com/sercand/kuberesolver/v5 v5.1.1 // indirect
	github.com/sirupsen/logrus v1.9.3 // indirect
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/sosodev/duration v1.1.0 // indirect
	github.com/uber/jaeger-client-go v2.30.0+incompatible // indirect
	github.com/uber/jaeger-lib v2.4.1+incompatible // indirect
	github.com/vektah/gqlparser/v2 v2.5.10 // indirect
	github.com/zclconf/go-cty v1.14.1 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel/metric v1.19.0 // indirect
	go.opentelemetry.io/otel/trace v1.19.0 // indirect
	go.uber.org/atomic v1.11.0 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/mod v0.13.0 // indirect
	golang.org/x/net v0.16.0 // indirect
	golang.org/x/oauth2 v0.13.0 // indirect
	golang.org/x/sync v0.4.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.14.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20231002182017-d307bd883b97 // indirect
	gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
