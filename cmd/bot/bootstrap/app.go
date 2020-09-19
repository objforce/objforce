package bootstrap

import(
	"github.com/objforce/objforce/cmd/bot/app/commands"
	"github.com/objforce/objforce/cmd/bot/app/providers"
	"github.com/xxxmicro/base/database/gorm"
	"go.uber.org/fx"
	xconfig "github.com/xxxmicro/base/config"
	xsource "github.com/xxxmicro/base/config/source"
	"github.com/xxxmicro/base/opentracing/jaeger"
)


func App() *fx.App {
	return fx.New(
		fx.Provide(providers.NewMicroService),

		// Providers (./app/providers)
		fx.Provide(xsource.NewSourceProvider),
		fx.Provide(xconfig.NewConfigProvider),
		fx.Provide(jaeger.NewTracerProvider),
		fx.Provide(gorm.NewDbProvider),

		// Commands (./app/commands)
		fx.Provide(commands.NewMigrateCommand),

		fx.Invoke(providers.RegisterCommands),
		fx.Invoke(providers.StartMicroService),
	)
}