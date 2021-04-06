package bootstrap

import (
	"github.com/objforce/objforce/service/data/app/domain/repositories"
	"github.com/objforce/objforce/service/data/app/providers"
	"go.uber.org/fx"
)

var Repositories = fx.Provide(
	providers.NewHBaseClientProvider,
	repositories.NewDataRepository,
	repositories.NewClobRepository,
)

var RepositoryOpts = fx.Options(
	Repositories,
)
