package providers


import (
	// ginzap "github.com/gin-contrib/zap"
	// "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
)

func NewGinProvider() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin := gin.New()
	// gin.Use(ginzap.Ginzap(log.Desugar(), time.RFC3339, true))
	// gin.Use(ginzap.RecoveryWithZap(log.Desugar(), true))

	return gin
}
