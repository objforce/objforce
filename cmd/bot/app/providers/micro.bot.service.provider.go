package providers

import (
	"context"
	"github.com/micro/cli/v2"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/fx"
	"net/http"
	"time"
)

func NewMicroService() (micro.Service, *cli.Context) {
	service := micro.NewService(
		micro.Name("com.xapis.bot.meta"),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
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
				Name: "prometheus_addr",
				Usage: "prometheus_addr",
				EnvVars: []string{"PROMETHEUS_ADDR"},
				Value: ":16627",
			},
		),
	)

	var cc *cli.Context
	service.Init(
		micro.Action(func(c *cli.Context) error {
			serverName := c.String("server_name")
			logger.Debugf("service.Init: %s", serverName)
			cc = c

			// prometheusBoot
			if len(c.String("prometheus_addr")) > 0 {
				prometheusBoot(c.String("prometheus_addr"))
			}
			return nil
		}),
	)

	return service, cc
}

func StartMicroService(lifecycle fx.Lifecycle, service micro.Service) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Logf(logger.DebugLevel, "lifecycle.OnStart")

			return service.Run()
		},
	})
}


func prometheusBoot(addr string) {
	logger.Debugf("prometheus start Listen: %s", addr)
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(addr, nil)
		if err != nil {
			logger.Fatal("ListenAndServe: ", err)
		}
	}()
}