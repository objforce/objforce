package providers

import (
	"github.com/micro/go-micro/v2"
	"github.com/objforce/objforce/cmd/index-srv/app/handlers"
	"github.com/objforce/objforce/idl/index/gen-go"
)

func RegisterHandlers(service micro.Service,
	indexHandler *handlers.IndexHandler,
	documentHandler *handlers.DocumentHandler,
) {
	index.RegisterDocumentServiceHandler(service.Server(), documentHandler)
}