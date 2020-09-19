package providers

import (
	"github.com/micro/go-micro/v2"
	data "github.com/objforce/objforce/idl/data/gen-go"
	"github.com/objforce/objforce/cmd/data-srv/app/handlers"
)

func RegisterHandlers(service micro.Service,
	dataHandler *handlers.SObjectHandler,
) {
	data.RegisterSObjectServiceHandler(service.Server(), dataHandler)
}