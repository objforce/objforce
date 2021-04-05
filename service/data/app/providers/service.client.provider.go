package providers

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/objforce/objforce/proto/meta"
)

func NewMetaClient(client client.Client) meta.CustomObjectService {
	return meta.NewCustomObjectService("com.xapis.srv.meta", client)
}
