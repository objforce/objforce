package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/index-server/app/domain/models"
	"github.com/objforce/objforce/index-server/app/domain/services"
	"go.uber.org/zap"
	"net/http"
)

type DocumentController struct {
	documentService services.IndexService
	log             *zap.SugaredLogger
}

func (c *DocumentController) Create(ctx *gin.Context) {
	c.log.Info("STARTING DocumentController.Create()")

	doc := &models.Document{}
	if err := ctx.ShouldBindJSON(&doc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.documentService.Create(ctx.Request.Context(), doc)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, doc)
	c.log.Infow("FINISHED DocumentController.Create()",
		"id", doc.Id,
	)
}

func (c *DocumentController) Delete(ctx *gin.Context) {
	c.log.Info("STARTING DocumentController.Delete()")

	index := ctx.Param("index")
	if len(index) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "need index"})
	}

	typ := ctx.Param("type")
	if len(typ) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "need type"})
	}

	id := ctx.Param("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "need id"})
	}

	err := c.documentService.Delete(ctx.Request.Context(), index, typ, id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0})
}

func NewDocumentController(documentService services.IndexService, log *zap.SugaredLogger) *DocumentController {
	return &DocumentController{
		documentService,
		log,
	}
}