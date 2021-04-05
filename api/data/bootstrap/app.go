package bootstrap

import (
	"github.com/objforce/objforce/api/data/app/domain/services"
	"github.com/objforce/objforce/api/data/app/http/controllers"
	"github.com/objforce/objforce/api/data/app/http/middlewares"
	"github.com/objforce/objforce/api/data/app/providers"
	"github.com/objforce/objforce/api/data/config"
	"github.com/objforce/objforce/api/data/routes"
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

		fx.Provide(providers.NewMicroClientProvider),
		fx.Provide(providers.NewSObjectClient),

		// Services (./app/services)
		fx.Provide(services.NewDataService),

		// Middlewares (./app/middlewares)
		fx.Provide(middlewares.NewLogMiddleware),

		// Controllers (./app/controllers)
		fx.Provide(controllers.NewAPIController),
		fx.Provide(controllers.NewSObjectController),

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
