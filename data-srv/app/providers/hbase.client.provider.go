package providers

import (
	"github.com/micro/go-micro/v2/config"
	"github.com/tsuna/gohbase"
)

func NewHBaseClientProvider(config config.Config) gohbase.Client {
	zkquorum := config.Get("hbase", "zkquorum").String("localhost")

	client := gohbase.NewClient(zkquorum)
	return client
}