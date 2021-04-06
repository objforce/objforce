package bootstrap

import (
	"github.com/objforce/objforce/service/meta/app/handlers"
	"github.com/objforce/objforce/service/meta/app/providers"
	"go.uber.org/fx"
)

var Handlers = fx.Provide(
	handlers.NewCustomObjectHandler,
	handlers.NewCustomFieldHandler,
)

var HandlerOpts = fx.Options(
	Handlers,
	fx.Invoke(providers.RegisterHandlers),
)
