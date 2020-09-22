package bootstrap

import (
	"github.com/objforce/objforce/cmd/data-srv/app/domain/repositories"
	"github.com/objforce/objforce/cmd/data-srv/app/domain/services"
	"github.com/objforce/objforce/cmd/data-srv/app/handlers"
	"github.com/objforce/objforce/cmd/data-srv/app/providers"
	"github.com/objforce/objforce/cmd/data-srv/config"
	xconfig "github.com/xxxmicro/base/config"
	xsource "github.com/xxxmicro/base/config/source"
	"github.com/xxxmicro/base/opentracing/jaeger"
	"go.uber.org/fx"
)


func App() *fx.App {
	return fx.New(
		fx.Provide(providers.NewMicroService),

		// Configurations (./config)
		fx.Provide(config.NewAppConfig),

		// Providers (./app/providers)
		fx.Provide(xsource.NewSourceProvider),
		fx.Provide(xconfig.NewConfigProvider),
		fx.Provide(jaeger.NewTracerProvider),
		fx.Provide(providers.NewHBaseClientProvider),
		fx.Provide(providers.NewBrokerProvider),
		fx.Provide(providers.NewMicroClientProvider),

		// apis
		fx.Provide(providers.NewMetaClient),

		// Repositories (./app/repositories)
		fx.Provide(repositories.NewDataRepository),
		fx.Provide(repositories.NewClobRepository),

		// Services (./app/services)
		fx.Provide(services.NewDataService),

		// Handlers (./app/handlers)
		fx.Provide(handlers.NewSObjectHandler),
		fx.Invoke(providers.RegisterHandlers),

		fx.Invoke(providers.InitLogger),
		fx.Invoke(providers.StartMicroService),
	)
}