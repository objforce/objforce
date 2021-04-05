package routes

import(
	"time"
	"github.com/gin-gonic/gin"
	"github.com/cnjack/throttle"
	"go.uber.org/dig"
	"github.com/objforce/objforce/cmd/meta-api/config"
	"github.com/objforce/objforce/cmd/meta-api/app/http/controllers"
	"github.com/objforce/objforce/cmd/meta-api/app/http/middlewares"
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

		customFields := api.Group("/customFields")
		{
			customFields.POST(
				"", 
				throttle.Policy(&throttle.Quota{
					Limit:  100,
					Within: time.Minute,
				}),
				c.CustomFieldController.Create)
		}
	}
}

type RouterContext struct {
	dig.In

	Router *gin.Engine
	AppConfig *config.AppConfig

	LogMiddleware *middlewares.LogMiddleware

	APIController *controllers.APIController
	CustomFieldController *controllers.CustomFieldController
}