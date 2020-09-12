package providers

import (
	"github.com/micro/go-micro/v2"
	"github.com/objforce/objforce/index-srv/app/handlers"
	"github.com/objforce/objforce/index-srv/proto/index/gen-go"
)

func RegisterHandlers(service micro.Service,
	indexHandler *handlers.IndexHandler,
	documentHandler *handlers.DocumentHandler,
) {
	index.RegisterDocumentServiceHandler(service.Server(), documentHandler)
}