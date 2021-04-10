package providers

import (
	"time"

	"github.com/micro/go-micro/v2/config"
	"github.com/tsuna/gohbase"
)

func NewHBaseClientProvider(config config.Config) gohbase.Client {
	zkquorum := config.Get("hbase", "zkquorum").String("localhost")
	regionLookupTimeout := config.Get("hbase", "region_lookup_timeout").Duration(time.Second * 10)
	regionReadTimeout := config.Get("hbase", "region_read_timeout").Duration(time.Second * 10)
	zookeeperTimeout := config.Get("hbase", "zookeeper_timeout").Duration(time.Second * 10)

	client := gohbase.NewClient(
		zkquorum,
		gohbase.RegionLookupTimeout(regionLookupTimeout),
		gohbase.RegionReadTimeout(regionReadTimeout),
		gohbase.ZookeeperTimeout(zookeeperTimeout),
	)
	return client
}
