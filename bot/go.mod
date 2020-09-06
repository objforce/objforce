module github.com/objforce/bot

go 1.14

require (
	github.com/jinzhu/gorm v1.9.16
	github.com/micro/cli/v2 v2.1.2
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/objforce/meta-server v0.0.0-00010101000000-000000000000
	github.com/prometheus/client_golang v1.1.0
	github.com/xxxmicro/base v0.1.27
	go.uber.org/fx v1.13.1
)

replace github.com/objforce/meta-server => ../meta-server
