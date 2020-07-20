package bootstrap

import(
	"go.uber.org/fx"
	"github.com/objforce/meta-server/config"
	"github.com/objforce/meta-server/app/http/controllers"
	"github.com/objforce/meta-server/app/http/middlewares"
	"github.com/objforce/meta-server/app/services"
	"github.com/objforce/meta-server/app/providers"
	"github.com/objforce/meta-server/app/repositories"
	"github.com/objforce/meta-server/routes"
)


func App() *fx.App {
	return fx.New(
		// Configurations (./config)
		fx.Provide(config.NewAppConfig),
		fx.Provide(config.NewDatabaseConfig),

		// Providers (./app/providers)
		fx.Provide(providers.NewConfigProvider),
		fx.Provide(providers.NewDatabaseProvider),
		fx.Provide(providers.NewLoggerProvider),
		fx.Provide(providers.NewServerProvider),

		// Repositories (./app/repositories)
		fx.Provide(repositories.NewCustomFieldRepository),

		// Services (./app/services)
		fx.Provide(services.NewCustomFieldService),

		// Middlewares (./app/middlewares)
		fx.Provide(middlewares.NewLogMiddleware),

		// Controllers (./app/controllers)
		fx.Provide(controllers.NewAPIController),
		fx.Provide(controllers.NewCustomFieldController),


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
		fx.Invoke(routes.APIRoutes),
	)
}