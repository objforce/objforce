package bootstrap

import(
	"github.com/objforce/objforce/data-srv/app/handlers"
	"go.uber.org/fx"
	xconfig "github.com/xxxmicro/base/config"
	xsource "github.com/xxxmicro/base/config/source"
	"github.com/xxxmicro/base/database/gorm"
	"github.com/xxxmicro/base/opentracing/jaeger"
	"github.com/objforce/objforce/data-srv/config"
	"github.com/objforce/objforce/data-srv/app/providers"
	"github.com/objforce/objforce/data-srv/app/domain/services"
	"github.com/objforce/objforce/data-srv/app/domain/repositories"
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
		fx.Provide(gorm.NewDbProvider),
		fx.Provide(providers.NewBrokerProvider),
		fx.Provide(providers.NewMicroClientProvider),

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