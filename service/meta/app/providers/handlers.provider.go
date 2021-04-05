package providers

import (
	"github.com/micro/go-micro/v2"
	"github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/service/meta/app/handlers"
)

func RegisterHandlers(service micro.Service,
	customObjectHandler *handlers.CustomObjectHandler,
	customFieldHandler *handlers.CustomFieldHandler,
) {
	meta.RegisterCustomObjectServiceHandler(service.Server(), customObjectHandler)
	meta.RegisterCustomFieldServiceHandler(service.Server(), customFieldHandler)
}
