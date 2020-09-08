package routes

import(
	"time"
	"github.com/gin-gonic/gin"
	"github.com/cnjack/throttle"
	"go.uber.org/dig"
	"github.com/objforce/objforce/index-server/config"
	"github.com/objforce/objforce/index-server/app/http/controllers"
	"github.com/objforce/objforce/index-server/app/http/middlewares"
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

		documents := api.Group("/documents")
		{
			documents.POST(
				"", 
				throttle.Policy(&throttle.Quota{
					Limit:  100,
					Within: time.Minute,
				}),
				c.DocumentController.Create)
		}
	}
}

type RouterContext struct {
	dig.In

	Router *gin.Engine
	AppConfig *config.AppConfig

	LogMiddleware *middlewares.LogMiddleware

	APIController      *controllers.APIController
	DocumentController *controllers.DocumentController
}