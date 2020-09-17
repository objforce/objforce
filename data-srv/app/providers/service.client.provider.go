package providers

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/objforce/objforce/idl/meta/gen-go"
)


func NewMetaClient(client client.Client) meta.CustomObjectService {
	return meta.NewCustomObjectService("com.xapis.srv.meta", client)
}