package bootstrap

import (
	"github.com/micro/go-micro/v2"
	"github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/server/meta/handlers"
	"go.uber.org/fx"
)

func RegisterHandlers(service micro.Service,
	customObjectHandler *handlers.CustomObjectHandler,
	customFieldHandler *handlers.CustomFieldHandler,
) {
	meta.RegisterCustomObjectServiceHandler(service.Server(), customObjectHandler)
	meta.RegisterCustomFieldServiceHandler(service.Server(), customFieldHandler)
}

var Handlers = fx.Provide(
	handlers.NewCustomObjectHandler,
	handlers.NewCustomFieldHandler,
)

var HandlerOpts = fx.Options(
	Handlers,
	fx.Invoke(RegisterHandlers),
)
