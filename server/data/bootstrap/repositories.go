package bootstrap

import (
	"github.com/objforce/objforce/service/data/domain/repositories/hbase"
	"github.com/objforce/objforce/service/data/domain/repositories/tablestore"
	"github.com/objforce/objforce/service/data/providers"
	"go.uber.org/fx"
)

var Repositories = fx.Provide(
	fx.Annotated{
		Name:   "tablestore.tenancy",
		Target: providers.NewTableStoreTenancy,
	},
	fx.Annotated{
		Name:   "tablestore.data.repository",
		Target: tablestore.NewDataRepository,
	},
	fx.Annotated{
		Name:   "hbase.data.repository",
		Target: hbase.NewDataRepository,
	},
	providers.NewHBaseClientProvider,
	// gorm.NewClobRepository,
)

var RepositoryOpts = fx.Options(
	Repositories,
)
