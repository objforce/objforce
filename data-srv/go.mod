module github.com/objforce/objforce/data-srv

go 1.14

require (
	github.com/aviddiviner/gin-limit v0.0.0-20170918012823-43b5f79762c1 // indirect
	github.com/cnjack/throttle v0.0.0-20160727064406-525175b56e18 // indirect
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e // indirect
	github.com/deckarep/golang-set v1.7.1
	github.com/gin-contrib/cors v1.3.1 // indirect
	github.com/gin-contrib/gzip v0.0.2 // indirect
	github.com/gin-contrib/static v0.0.0-20191128031702-f81c604d8ac2 // indirect
	github.com/gin-contrib/zap v0.0.1 // indirect
	github.com/gin-gonic/gin v1.6.3 // indirect
	github.com/golang/protobuf v1.4.0
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/logger/zap/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/validator/v2 v2.9.1
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/opentracing/opentracing-go v1.1.0
	github.com/prometheus/client_golang v1.5.1
	github.com/satori/go.uuid v1.2.0
	github.com/thinkerou/favicon v0.1.0 // indirect
	github.com/tsuna/gohbase v0.0.0-20200831170559-79db14850535
	github.com/wantedly/gorm-zap v0.0.0-20171015071652-372d3517a876 // indirect
	github.com/xxxmicro/base v0.1.28
	github.com/xxxmicro/go-plugins-broker-rocketmq/v2 v2.0.0-20200805004454-ef23e9db26f3
	go.uber.org/dig v1.10.0 // indirect
	go.uber.org/fx v1.13.0
	go.uber.org/zap v1.15.0 // indirect
	gorm.io/gorm v0.2.23 // indirect
	modernc.org/mathutil v1.1.1 // indirect
	modernc.org/strutil v1.1.0 // indirect
)

replace (
	github.com/objforce/objforce/idl => ../idl
)
