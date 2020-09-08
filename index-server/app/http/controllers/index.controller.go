package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/index-server/app/domain/services"
	"github.com/objforce/objforce/index-server/app/dtos"
	"go.uber.org/zap"
	"net/http"
)

type IndexController struct {
	indexService services.IndexService
	log          *zap.SugaredLogger
}

func (c *IndexController) Create(ctx *gin.Context) {
	c.log.Info("STARTING IndexController.Create()")

	index := &dtos.Index{}

	if err := ctx.ShouldBindJSON(&index); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	index, err := c.indexService.Create(ctx.Request.Context(), index)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, index)
	c.log.Infow("FINISHED IndexController.Create()",
		"ObjId", index.ObjId,
	)
}

func (c *IndexController) Delete(ctx *gin.Context) {
	c.log.Info("STARTING IndexController.Delete()")

	id := ctx.Param("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "need id"})
	}

	err := c.indexService.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0})
}

func NewIndexController(indexService services.IndexService, log *zap.SugaredLogger) *IndexController {
	return &IndexController{
		indexService,
		log,
	}
}