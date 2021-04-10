package bootstrap

import (
	"github.com/objforce/objforce/service/data/domain/repositories/hbase"
	"github.com/objforce/objforce/service/data/providers"
	"go.uber.org/fx"
)

var Repositories = fx.Provide(
	providers.NewHBaseClientProvider,
	hbase.NewDataRepository,
	// gorm.NewClobRepository,
)

var RepositoryOpts = fx.Options(
	Repositories,
)
