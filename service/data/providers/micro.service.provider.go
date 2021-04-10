package providers

import (
	"context"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	"github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	xopentracing "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/go-plugins/wrapper/validator/v2"

	// "github.com/micro/go-plugins/wrapper/validator/v2"
	"github.com/opentracing/opentracing-go"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	xxxmicro_opentracing "github.com/xxxmicro/base/opentracing"
	"go.uber.org/fx"
	"net/http"
	"time"
)

func NewMicroService() (micro.Service, *cli.Context) {

	// use grpc server
	// server := grpc.NewServer(server.WrapHandler(validator.NewHandlerWrapper()))
	QPS := 5000

	service := micro.NewService(
		micro.Name("com.xapis.srv.data"),
		micro.Version("v1"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
		micro.WrapHandler(validator.NewHandlerWrapper()),
		micro.WrapHandler(xopentracing.NewHandlerWrapper(xxxmicro_opentracing.GlobalTracerWrapper())),
		micro.WrapHandler(prometheus.NewHandlerWrapper(prometheus.ServiceName("data-srv"), prometheus.ServiceVersion("v1"))),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		micro.WrapSubscriber(xopentracing.NewSubscriberWrapper(xxxmicro_opentracing.GlobalTracerWrapper())),
		// micro.Server(server),
		micro.Flags(
			&cli.StringFlag{
				Name:  "apollo_namespace",
				Usage: "apollo_namespace",
			},
			&cli.StringFlag{
				Name:  "apollo_address",
				Usage: "apollo_address",
			},
			&cli.StringFlag{
				Name:  "apollo_app_id",
				Usage: "apollo_app_id",
			},
			&cli.StringFlag{
				Name:  "apollo_cluster",
				Usage: "apollo_cluster",
			},
			&cli.StringFlag{
				Name:  "prometheus_addr",
				Usage: "prometheus_addr",
				EnvVars: []string{"PROMETHEUS_ADDR"},
				Value: ":16627",
			},
		),
	)

	var cc *cli.Context
	service.Init(
		micro.Action(func(c *cli.Context) error {
			logger.Log(logger.DebugLevel, "service.Init")
			cc = c

			if len(c.String("prometheus_addr")) > 0 {
				prometheusBoot(c.String("prometheus_addr"))
			}

			return nil
		}),
	)

	return service, cc
}

func StartMicroService(lifecycle fx.Lifecycle, service micro.Service, broker broker.Broker, tracer opentracing.Tracer) {
	o := service.Options()
	micro.Broker(broker)(&o)

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Log(logger.DebugLevel, "lifecycle.OnStart")
			return service.Run()
		},
	})
}

func prometheusBoot(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			logger.Fatal("ListenAndServe: ", err)
		}
	}()
}
