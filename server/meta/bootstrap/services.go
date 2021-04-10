package bootstrap

import (
	"github.com/objforce/objforce/service/meta/domain/services"
	"go.uber.org/fx"
)

var Services = fx.Provide(
	services.NewCustomFieldService,
	services.NewCustomObjectService,
)

var ServiceOpts = fx.Options(
	Services,
)
