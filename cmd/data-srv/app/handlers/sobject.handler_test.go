package handlers_test

import (
	"context"
	"github.com/micro/go-micro/v2"
	config2 "github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/memory"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/micro/go-plugins/wrapper/monitoring/prometheus/v2"
	ratelimit "github.com/micro/go-plugins/wrapper/ratelimiter/uber/v2"
	xopentracing "github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"github.com/micro/go-plugins/wrapper/validator/v2"
	"github.com/objforce/objforce/data-srv/app/domain/repositories"
	"github.com/objforce/objforce/data-srv/app/domain/services"
	"github.com/objforce/objforce/data-srv/app/handlers"
	"github.com/objforce/objforce/data-srv/app/providers"
	"github.com/objforce/objforce/data-srv/config"
	data "github.com/objforce/objforce/idl/data/gen-go"
	xxxmicro_opentracing "github.com/xxxmicro/base/opentracing"
	"testing"
	"time"
)

func TestNewSObjectHandler(t *testing.T) {
	source := memory.NewSource(
		memory.WithJSON([]byte(`
	{
		"hbase": {
			"zkquorum": "localhost"
		}
	}
	`)))

	conf, err := config2.NewConfig()
	if err != nil {
		t.Fatal(err)
	}

	if err := conf.Load(source); err != nil {
		t.Fatal(err)
	}

	gohbaseClient := providers.NewHBaseClientProvider(conf)

	dataRepository := repositories.NewDataRepository(gohbaseClient)

	registry := consul.NewRegistry(registry.Addrs("localhost:8500"))
	service := micro.NewService(
		micro.Name("com.xapis.srv.data"),
		micro.Version("v1"),
		micro.Registry(registry),
		micro.RegisterTTL(time.Minute),
		micro.RegisterInterval(time.Second*30),
		micro.WrapHandler(validator.NewHandlerWrapper()),
		micro.WrapHandler(xopentracing.NewHandlerWrapper(xxxmicro_opentracing.GlobalTracerWrapper())),
		micro.WrapHandler(prometheus.NewHandlerWrapper(prometheus.ServiceName("data-srv"), prometheus.ServiceVersion("v1"))),
		micro.WrapHandler(ratelimit.NewHandlerWrapper(5000)),
		micro.WrapSubscriber(xopentracing.NewSubscriberWrapper(xxxmicro_opentracing.GlobalTracerWrapper())),
	)

	client := providers.NewMicroClientProvider(service)

	customObjectService := providers.NewMetaClient(client)

	dataService := services.NewDataService(dataRepository, customObjectService)

	objectHandler := handlers.NewSObjectHandler(dataService)



	ctx := context.Background()
	ctx = context.WithValue(ctx, config.OrgIdKey{}, "1234567890123456789012345678901234567890")

	req := &data.CreateSObjectRequest{
		Objects: []*data.SObject{
			{
				Type: "customers",

			},
		},
	}

	rsp := &data.CreateSObjectResponse{}

	err = objectHandler.Create(ctx, req, rsp)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(rsp)
}