package providers

import (
	"github.com/micro/go-micro/v2/client"
	meta "github.com/objforce/objforce/index-srv/proto/meta/gen-go"
)

func NewMetaService(client client.Client) meta.CustomObjectService {
	return meta.NewCustomObjectService("com.xapis.srv.meta", client)
}