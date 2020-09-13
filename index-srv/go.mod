module github.com/objforce/objforce/index-srv

go 1.14

require (
	github.com/golang/protobuf v1.4.0
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/logger/zap/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/validator/v2 v2.9.1
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/objforce/objforce/meta-srv v0.0.0-20200911204911-c7416e930fde
	github.com/olivere/elastic v6.2.35+incompatible // indirect
	github.com/olivere/elastic/v6 v6.2.1
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.5.1
	github.com/xxxmicro/base v0.1.28
	go.uber.org/fx v1.13.0
)
