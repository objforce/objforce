package providers

import (
	"github.com/micro/go-micro/v2"
	"github.com/objforce/objforce/proto/index"
	"github.com/objforce/objforce/service/index/app/handlers"
)

func RegisterHandlers(service micro.Service,
	indexHandler *handlers.IndexHandler,
	documentHandler *handlers.DocumentHandler,
) {
	index.RegisterDocumentServiceHandler(service.Server(), documentHandler)
}
