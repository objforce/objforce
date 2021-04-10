package bootstrap

import (
	"github.com/objforce/objforce/service/data/providers"
	"go.uber.org/fx"
)

var RpcOpts = fx.Options(
	fx.Provide(providers.NewMetaClient),
)
