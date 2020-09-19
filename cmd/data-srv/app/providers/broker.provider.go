package providers

import (
	"github.com/micro/go-micro/v2/broker"
	"github.com/micro/go-micro/v2/config"
	"github.com/xxxmicro/go-plugins-broker-rocketmq/v2"
	xconfig "github.com/objforce/objforce/data-srv/config"
)

func NewBrokerProvider(config config.Config) (broker.Broker, error) {
	var cfg xconfig.RocketmqConfig
	if err := config.Get("rocketmq").Scan(&cfg); err != nil {
		return nil, err
	}

	opts := make([]broker.Option, 0)
	opts = append(opts, broker.Addrs(cfg.Addrs...))

	if len(cfg.AccessKey) > 0 {
		opts = append(opts, rocketmq.WithCredentials(
			rocketmq.Credentials{
				AccessKey: cfg.AccessKey,
				SecretKey: cfg.SecretKey,
			},
		))
	}	
	
	opts = append(opts, rocketmq.WithRetry(cfg.Retry))

	b := rocketmq.NewBroker(opts...)

	b.Init()
	if err := b.Connect(); err != nil {
		return nil, err
	}

	return b, nil
}