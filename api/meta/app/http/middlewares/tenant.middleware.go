package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
)

func TenantHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		orgId := ctx.GetHeader("ORG_ID")

		ctx.Request.WithContext(context.WithValue(ctx.Request.Context(), "orgId", orgId))

		ctx.Next()
	}
}