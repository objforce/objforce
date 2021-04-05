package routes

import (
	"time"

	"github.com/cnjack/throttle"
	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/api/index/app/http/controllers"
	"github.com/objforce/objforce/api/index/app/http/middlewares"
	"github.com/objforce/objforce/api/index/config"
	"go.uber.org/dig"
)

func APIRoutes(c RouterContext) {
	api := c.Router.Group(c.AppConfig.Prefix)
	{
		/*
		 |----------------------------------------
		 | Application Routes
		 |----------------------------------------
		 |
		*/

		api.GET("/ping", c.APIController.Ping)

		documents := api.Group("/documents/upsert")
		{
			documents.POST(
				"",
				throttle.Policy(&throttle.Quota{
					Limit:  100,
					Within: time.Minute,
				}),
				c.DocumentController.Upsert,
			)

			documents.POST(
				"/bulk",
				throttle.Policy(&throttle.Quota{
					Limit:  100,
					Within: time.Minute,
				}),
				c.DocumentController.Bulk,
			)

			documents.DELETE(
				"",
				throttle.Policy(&throttle.Quota{
					Limit:  100,
					Within: time.Minute,
				}),
				c.DocumentController.Delete,
			)
		}
	}
}

type RouterContext struct {
	dig.In

	Router    *gin.Engine
	AppConfig *config.AppConfig

	LogMiddleware *middlewares.LogMiddleware

	APIController      *controllers.APIController
	DocumentController *controllers.DocumentController
}
