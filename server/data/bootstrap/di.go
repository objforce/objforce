package bootstrap

import (
	"go.uber.org/fx"

	framework "github.com/duolacloud/microbase/providers"
	"github.com/urfave/cli/v2"
)

func Setup(c *cli.Context) *fx.App {
	return fx.New(
		fx.Provide(
			func() *cli.Context {
				return c
			},
		),
		framework.FrameworkOpts,
		framework.CloudNativeOpts,
		RepositoryOpts,
		ServiceOpts,
		RpcOpts,
		HandlerOpts,
		framework.MakeMicroServiceOpts(c))
}
