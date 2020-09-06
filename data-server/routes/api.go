package routes

import(
	"time"
	"github.com/gin-gonic/gin"
	"github.com/cnjack/throttle"
	"go.uber.org/dig"
	"github.com/objforce/objforce/data-server/config"
	"github.com/objforce/objforce/data-server/app/http/controllers"
	"github.com/objforce/objforce/data-server/app/http/middlewares"
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

		sobjects := api.Group("/sobjects")
		{
			sobjects.POST(
				"", 
				throttle.Policy(&throttle.Quota{
					Limit:  100,
					Within: time.Minute,
				}),
				c.SObjectController.Create)
		}
	}
}

type RouterContext struct {
	dig.In

	Router *gin.Engine
	AppConfig *config.AppConfig

	LogMiddleware *middlewares.LogMiddleware

	APIController *controllers.APIController
	SObjectController *controllers.SObjectController
}