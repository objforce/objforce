package bootstrap

import (
	"github.com/objforce/objforce/service/data/domain/services"
	"go.uber.org/fx"
)

var Services = fx.Provide(
	services.NewDataService,
)

var ServiceOpts = fx.Options(
	Services,
)
