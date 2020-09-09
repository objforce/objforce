package providers

import (
	"github.com/micro/go-micro/v2"
	data "github.com/objforce/objforce/data-srv/proto/data/gen-go"
	"github.com/objforce/objforce/data-srv/app/handlers"
)

func RegisterHandlers(service micro.Service,
	dataHandler *handlers.SObjectHandler,
) {
	data.RegisterSObjectServiceHandler(service.Server(), dataHandler)
}