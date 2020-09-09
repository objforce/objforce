package providers

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

func NewMicroClientProvider(service micro.Service) client.Client {
	return service.Client()
}