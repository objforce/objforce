module github.com/objforce/objforce

go 1.14

require (
	github.com/aviddiviner/gin-limit v0.0.0-20170918012823-43b5f79762c1
	github.com/cnjack/throttle v0.0.0-20160727064406-525175b56e18
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/codegangsta/inject v0.0.0-20150114235600-33e0aa1cb7c0 // indirect
	github.com/danielkov/gin-helmet v0.0.0-20171108135313-1387e224435e
	github.com/deckarep/golang-set v1.7.1
	github.com/duolacloud/microbase v0.0.0-20210328154445-22250c28bd60
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-contrib/gzip v0.0.2
	github.com/gin-contrib/static v0.0.0-20191128031702-f81c604d8ac2
	github.com/gin-gonic/gin v1.6.3
	github.com/ginkgoch/godash v1.2.0
	github.com/go-martini/martini v0.0.0-20170121215854-22fa46961aab // indirect
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/golang/protobuf v1.4.3
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/now v1.1.1 // indirect
	github.com/mattheath/base62 v0.0.0-20150408093626-b80cdc656a7a
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/logger/zap/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/breaker/hystrix/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/validator/v2 v2.9.1
	github.com/mitchellh/mapstructure v1.4.1
	github.com/mwitkow/go-proto-validators v0.3.2
	github.com/objforce/objflake v0.0.0-20210403191318-ddf85f176be7 // indirect
	github.com/objforce/objforce/index-srv v0.0.0-20200913195437-10309c804e6f
	github.com/olivere/elastic v6.2.35+incompatible // indirect
	github.com/olivere/elastic/v6 v6.2.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/prometheus/client_golang v1.5.1
	github.com/satori/go.uuid v1.2.0
	github.com/sony/sonyflake v1.0.0
	github.com/thinkerou/favicon v0.1.0
	github.com/tsuna/gohbase v0.0.0-20200831170559-79db14850535
	github.com/urfave/cli/v2 v2.3.0
	github.com/xxxmicro/base v0.1.28
	github.com/xxxmicro/go-plugins-broker-rocketmq/v2 v2.0.0-20200805004454-ef23e9db26f3
	go.uber.org/dig v1.10.0
	go.uber.org/fx v1.13.1
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20201110031124-69a78807bb2b
	google.golang.org/grpc v1.33.2
	google.golang.org/protobuf v1.25.0
	modernc.org/mathutil v1.1.1 // indirect
	modernc.org/strutil v1.1.0 // indirect
)

replace (
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)