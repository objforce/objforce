package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


type ActionHandler func (ctx *gin.Context)

type APIController struct{
	actionHandlers map[string]ActionHandler
}

func (c *APIController) Ping(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}

func (c *APIController) Call(ctx *gin.Context) {
	action := ctx.Param("action")
	if len(action) == 0 {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	params := map[string]interface{}{}
	if err := ctx.BindJSON(params); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	handler, ok := c.actionHandlers[action]
	if !ok {
		ctx.AbortWithStatusJSON(http.StatusOK, gin.H{"code": -1, "msg": fmt.Sprintf("handler %s not exists", action)})
		return
	}

	handler(ctx)
}

func NewAPIController() *APIController {
	c := APIController{}

	actionHandlers := make(map[string]ActionHandler)
	actionHandlers["metadata.describe"] = c.CreateMetadata
	actionHandlers["metadata.describe"] = c.DescribeMetadata
	actionHandlers["metadata.list"] = c.ListMetadata
	actionHandlers["metadata.read"] = c.ReadMetadata
	actionHandlers["metadata.rename"] = c.RenameMetadata
	actionHandlers["metadata.update"] = c.UpdateMetadata
	actionHandlers["metadata.upsert"] = c.UpsertMetadata
	actionHandlers["metadata.delete"] = c.DeleteMetadata

	c.actionHandlers = actionHandlers

	return &c
}