package providers

import (
	"github.com/micro/go-micro/v2"
	data "github.com/objforce/objforce/proto/data"
	"github.com/objforce/objforce/service/data/app/handlers"
)

func RegisterHandlers(service micro.Service,
	dataHandler *handlers.SObjectHandler,
) {
	data.RegisterSObjectServiceHandler(service.Server(), dataHandler)
}
