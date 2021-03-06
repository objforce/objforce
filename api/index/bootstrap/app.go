package bootstrap

import (
	"github.com/objforce/objforce/api/index/app/domain/services"
	"github.com/objforce/objforce/api/index/app/http/controllers"
	"github.com/objforce/objforce/api/index/app/http/middlewares"
	"github.com/objforce/objforce/api/index/app/providers"
	"github.com/objforce/objforce/api/index/config"
	"github.com/objforce/objforce/api/index/routes"
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
		fx.Provide(providers.NewGinProvider),
		fx.Provide(providers.NewLoggerProvider),
		fx.Provide(providers.NewElasticClientProvider),

		// Services (./app/services)
		fx.Provide(services.NewDocumentService),
		fx.Provide(services.NewIndexService),

		// Middlewares (./app/middlewares)
		fx.Provide(middlewares.NewLogMiddleware),

		// Controllers (./app/controllers)
		fx.Provide(controllers.NewAPIController),
		fx.Provide(controllers.NewDocumentController),

		/*
			|--------------------------------------------------------------------------
			| Invoke Register Routes
			|--------------------------------------------------------------------------
			|
			| Here we add our api endpoints to the application. These routes are prefixed
			| with the default value 'api'. Moreover we pass a function to the container build
			| up, which can create a new database connection.
			|
		*/
		fx.Invoke(middlewares.GlobalMiddlewares),
		fx.Invoke(routes.APIRoutes),
		fx.Invoke(providers.StartMicroService),
	)
}
