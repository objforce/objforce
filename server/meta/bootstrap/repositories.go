package bootstrap

import (
	"github.com/objforce/objforce/service/meta/app/domain/repositories/gorm"
	"go.uber.org/fx"
)

var Repositories = fx.Provide(
	gorm.NewCustomObjectRepository,
	gorm.NewCustomFieldRepository,
	gorm.NewClobRepository,
)

var RepositoryOpts = fx.Options(
	Repositories,
)
