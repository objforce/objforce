package providers

import (
	"github.com/micro/go-micro/v2/client"
	"github.com/objforce/objforce/proto/data"
)

func NewSObjectClient(client client.Client) data.SObjectService {
	return data.NewSObjectService("com.xapis.srv.data", client)
}
