package bootstrap

import(
	"github.com/objforce/objforce/index-srv/app/events"
	"github.com/objforce/objforce/index-srv/app/handlers"
	"go.uber.org/fx"
	xconfig "github.com/xxxmicro/base/config"
	xsource "github.com/xxxmicro/base/config/source"
	gorm "github.com/xxxmicro/base/database/gorm"
	"github.com/xxxmicro/base/opentracing/jaeger"
	"github.com/objforce/objforce/index-srv/config"
	"github.com/objforce/objforce/index-srv/app/providers"
	"github.com/objforce/objforce/index-srv/app/domain/services"
	"github.com/objforce/objforce/index-srv/app/domain/repositories"
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
		fx.Provide(providers.NewMicroClientProvider),
		fx.Provide(providers.NewElasticClientProvider),
		fx.Provide(providers.NewMetaService),

		// Repositories (./app/repositories)
		fx.Provide(repositories.NewIndexRepository),
		fx.Provide(repositories.NewDocumentRepository),

		// Services (./app/services)
		fx.Provide(services.NewDocumentService),
		fx.Provide(services.NewIndexService),

		// Events (./app/events)
		fx.Invoke(events.NewSObjectSubscriber),

		// Handlers (./app/handlers)
		fx.Provide(handlers.NewIndexHandler),
		fx.Provide(handlers.NewDocumentHandler),
		fx.Invoke(providers.RegisterHandlers),

		fx.Invoke(providers.StartMicroService),
	)
}