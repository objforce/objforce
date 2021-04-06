package config

type RocketmqConfig struct {
	Addrs     []string `json:"addrs"`
	AccessKey string   `json:"access_key"`
	SecretKey string   `json:"secret_key"`
	Retry     int      `json:"retry"`
}
