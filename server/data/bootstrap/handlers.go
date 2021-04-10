package bootstrap

import (
	"github.com/micro/go-micro/v2"
	"github.com/objforce/objforce/proto/data"
	"github.com/objforce/objforce/server/data/handlers"
	"go.uber.org/fx"
)

func RegisterHandlers(service micro.Service,
	dataHandler *handlers.SObjectHandler,
) {
	data.RegisterSObjectServiceHandler(service.Server(), dataHandler)
}

var Handlers = fx.Provide(
	handlers.NewSObjectHandler,
)

var HandlerOpts = fx.Options(
	Handlers,
	fx.Invoke(RegisterHandlers),
)
