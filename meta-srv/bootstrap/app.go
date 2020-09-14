package bootstrap

import (
	"github.com/objforce/objforce/meta-srv/app/domain/repositories"
	"github.com/objforce/objforce/meta-srv/app/domain/services"
	"github.com/objforce/objforce/meta-srv/app/handlers"
	"github.com/objforce/objforce/meta-srv/app/providers"
	"github.com/objforce/objforce/meta-srv/config"
	xconfig "github.com/xxxmicro/base/config"
	xsource "github.com/xxxmicro/base/config/source"
	gorm "github.com/xxxmicro/base/database/gorm"
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
		fx.Provide(gorm.NewDbProvider),

		fx.Provide(providers.NewBrokerProvider),
		fx.Provide(providers.NewMicroClientProvider),

		// Repositories (./app/repositories)
		fx.Provide(repositories.NewCustomObjectRepository),
		fx.Provide(repositories.NewCustomFieldRepository),
		fx.Provide(repositories.NewClobRepository),

		// Services (./app/services)
		fx.Provide(services.NewCustomObjectService),
		fx.Provide(services.NewCustomFieldService),

		// Handlers (./app/handlers)
		fx.Provide(handlers.NewCustomObjectHandler),
		fx.Provide(handlers.NewCustomFieldHandler),
		fx.Provide(providers.RegisterHandlers),

		fx.Invoke(providers.StartMicroService),
	)
}