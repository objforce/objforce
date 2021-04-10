package bootstrap

import (
	"go.uber.org/fx"

	framework "github.com/duolacloud/microbase/providers"
	"github.com/objforce/objflake"
	"github.com/urfave/cli/v2"
)

func Setup(c *cli.Context) *fx.App {
	return fx.New(
		fx.Provide(
			func() *cli.Context {
				return c
			},
		),
		fx.Provide(objflake.NewIDGenerator),
		framework.FrameworkOpts,
		framework.CloudNativeOpts,
		DataSourceOpts,
		RepositoryOpts,
		ServiceOpts,
		HandlerOpts,
		framework.MakeMicroServiceOpts(c))
}
