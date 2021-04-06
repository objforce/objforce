package main

import (
	"context"

	"github.com/duolacloud/microbase/cmd"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"github.com/objforce/objforce/server/meta/bootstrap"
	"github.com/urfave/cli/v2"
)

func main() {
	cmd.Run(func(c *cli.Context) error {
		app := bootstrap.Setup(c)
		return app.Start(context.Background())
	}, nil)
}
