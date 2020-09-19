package middlewares

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/index-server/config"
	"github.com/opentracing/opentracing-go"
	"github.com/xxxmicro/base/http/gin/middlewares"
)

func GlobalMiddlewares(server *gin.Engine, config *config.AppConfig, logMiddleware *LogMiddleware, tracer opentracing.Tracer) {
	promMonitor := NewPrometheusMonitor("objforce", "com.xapis.api.index")

	server.Use(
		logMiddleware.Handler(),
		cors.Default(),
		// limit.MaxAllowed(config.Connection),
		gzip.Gzip(gzip.DefaultCompression),
		// traceMiddleware.Handler(),
		middlewares.Middleware(tracer),
		promMonitor.PromMiddleware(),
	)
}