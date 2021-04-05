package providers

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/objforce/objforce/proto/meta"
)

func NewCustomObjectService(client client.Client) meta.CustomObjectService {
	return meta.NewCustomObjectService("com.xapis.srv.meta", client)
}

func NewCustomFieldService(client client.Client) meta.CustomFieldService {
	return meta.NewCustomFieldService("com.xapis.srv.meta", client)
}
