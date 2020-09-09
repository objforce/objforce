package providers

import (
	"github.com/micro/go-micro/v2"
	"github.com/objforce/objforce/meta-srv/app/handlers"
	"github.com/objforce/objforce/meta-srv/proto/meta/gen-go"
)

func RegisterHandlers(service micro.Service,
	customObjectHandler *handlers.CustomObjectHandler,
	customFieldHandler *handlers.CustomFieldHandler,
) {
	meta.RegisterCustomObjectServiceHandler(service.Server(), customObjectHandler)
	meta.RegisterCustomFieldServiceHandler(service.Server(), customFieldHandler)
}