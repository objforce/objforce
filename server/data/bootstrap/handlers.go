package bootstrap

import (
	"github.com/objforce/objforce/service/data/app/handlers"
	"github.com/objforce/objforce/service/data/app/providers"
	"go.uber.org/fx"
)

var Handlers = fx.Provide(
	handlers.NewSObjectHandler,
)

var HandlerOpts = fx.Options(
	Handlers,
	fx.Invoke(providers.RegisterHandlers),
)
