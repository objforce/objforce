package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/cmd/index-api/app/domain/models"
	"github.com/objforce/objforce/cmd/index-api/app/domain/services"
	"github.com/ginkgoch/godash"
	"go.uber.org/zap"
	"net/http"
)

type DocumentController struct {
	documentService services.DocumentService
	log             *zap.SugaredLogger
}

func (c *DocumentController) Upsert(ctx *gin.Context) {
	c.log.Info("STARTING DocumentController.Upsert()")

	doc := &models.Document{}
	if err := ctx.ShouldBindJSON(&doc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := c.documentService.Upsert(ctx.Request.Context(), doc)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, doc)
	c.log.Infow("FINISHED DocumentController.Upsert()",
		"id", doc.Id,
	)
}

func (c *DocumentController) Bulk(ctx *gin.Context) {
	c.log.Info("STARTING DocumentController.Bulk()")

	docs := make([]*models.Document, 0)
	if err := ctx.ShouldBindJSON(&docs); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dashSlice := godash.NewDashSlice(docs)
	ids := godash.Map(dashSlice, func(in interface{}) interface{} {
		return in.(models.Document).Id
	})
	ctx.JSON(http.StatusOK, docs)
	c.log.Infow("FINISHED DocumentController.Bulk()", "ids", ids)
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

func NewDocumentController(documentService services.DocumentService, log *zap.SugaredLogger) *DocumentController {
	return &DocumentController{
		documentService,
		log,
	}
}