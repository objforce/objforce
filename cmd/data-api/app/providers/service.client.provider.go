package providers

import (
	"github.com/micro/go-micro/v2/client"
	data "github.com/objforce/objforce/idl/data/gen-go"
)

func NewSObjectClient(client client.Client) data.SObjectService {
	return data.NewSObjectService("com.xapis.srv.data", client)
}