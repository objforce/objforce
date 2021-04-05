package middlewares

import (
	limit "github.com/aviddiviner/gin-limit"
	helmet "github.com/danielkov/gin-helmet"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/api/meta/config"
	"github.com/opentracing/opentracing-go"
	"github.com/xxxmicro/base/http/gin/middlewares"
)

func GlobalMiddlewares(server *gin.Engine, config *config.AppConfig, logMiddleware *LogMiddleware, tracer opentracing.Tracer) {
	promMonitor := NewPrometheusMonitor("data_api", "com.xapis.api.meta")

	server.Use(
		// logMiddleware.Handler(),
		cors.Default(),
		limit.MaxAllowed(config.Connection),
		gzip.Gzip(gzip.DefaultCompression),
		helmet.NoSniff(),
		helmet.DNSPrefetchControl(),
		helmet.FrameGuard(),
		helmet.SetHSTS(true),
		helmet.IENoOpen(),
		helmet.XSSFilter(),
		helmet.NoCache(),
		static.Serve("/", static.LocalFile("./public", false)),
		middlewares.Middleware(tracer),
		TenantHandler(),
		promMonitor.PromMiddleware(),
	)
}
