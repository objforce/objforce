package bootstrap

import(
	"go.uber.org/fx"
	xconfig "github.com/xxxmicro/base/config"
	xsource "github.com/xxxmicro/base/config/source"
	gorm "github.com/xxxmicro/base/database/gorm"
	"github.com/xxxmicro/base/opentracing/jaeger"
	"github.com/objforce/objforce/data-server/config"
	"github.com/objforce/objforce/data-server/app/http/controllers"
	"github.com/objforce/objforce/data-server/app/http/middlewares"
	"github.com/objforce/objforce/data-server/app/providers"
	"github.com/objforce/objforce/data-server/app/domain/services"
	"github.com/objforce/objforce/data-server/app/domain/repositories"
	"github.com/objforce/objforce/data-server/routes"
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

		// Repositories (./app/repositories)
		fx.Provide(repositories.NewDataRepository),
		fx.Provide(repositories.NewClobRepository(),

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
	)
}