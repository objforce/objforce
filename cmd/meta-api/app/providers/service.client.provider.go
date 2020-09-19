package providers

import (
	"github.com/micro/go-micro/v2/client"
	meta "github.com/objforce/objforce/meta-api/proto/meta/gen-go"
)

func NewCustomObjectService(client client.Client) meta.CustomObjectService {
	return meta.NewCustomObjectService("com.xapis.srv.meta", client)
}

func NewCustomFieldService(client client.Client) meta.CustomFieldService {
	return meta.NewCustomFieldService("com.xapis.srv.meta", client)
}