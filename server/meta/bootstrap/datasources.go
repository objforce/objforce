package bootstrap

import (
	"github.com/duolacloud/microbase/datasource/gorm"
	"github.com/duolacloud/microbase/domain/repository"
	"github.com/micro/go-micro/v2/config"
	"github.com/objforce/objforce/service/meta/app/providers"
	"go.uber.org/fx"
)

var DataSources = fx.Provide(
	fx.Annotated{
		Name: "gorm",
		Target: func(config config.Config) (repository.DataSourceProvider, error) {
			tenancy, err := gorm.NewGormTenancy(config, providers.NewOLTPEntityMap())
			if err != nil {
				return nil, err
			}
			return repository.NewMultitenancyProvider(tenancy), nil
		},
	},
)

var DataSourceOpts = fx.Options(
	DataSources,
)
