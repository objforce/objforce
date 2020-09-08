package providers

import (
	"errors"
	"github.com/micro/go-micro/v2/config"
	"github.com/olivere/elastic/v6"
)

func NewElasticClientProvider(config config.Config) (*elastic.Client, error) {
	urls := config.Get("elastic", "urls").StringSlice([]string{"http://localhost:9200"})
	if urls == nil {
		return nil, errors.New("elastic need urls")
	}

	client, err := elastic.NewClient(
		elastic.SetURL(urls...),
		elastic.SetSniff(false),
	)
	return client, err
}
