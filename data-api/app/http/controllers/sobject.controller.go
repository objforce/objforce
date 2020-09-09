package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/data-server/app/domain/services"
	"github.com/objforce/objforce/data-server/app/dtos"
	"go.uber.org/zap"
	"net/http"
)

type SObjectController struct {
	dataService services.DataService
	log *zap.SugaredLogger 
}

func (c *SObjectController) Create(ctx *gin.Context) {
	c.log.Info("STARTING CustomFieldController.Create()")

	customField := &dtos.SObject{}

	if err := ctx.ShouldBindJSON(&customField); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customField, err := c.dataService.Create(ctx.Request.Context(), customField)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, customField)
	c.log.Infow("FINISHED CustomFieldController.Create()",
		"id", customField.Id,
		"ObjId", customField.ObjId,
	)
}

func (c *SObjectController) Delete(ctx *gin.Context) {
	c.log.Info("STARTING CustomFieldController.Delete()")

	id := ctx.Param("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "need id"})
	}

	err := c.dataService.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0})
}

func NewSObjectController(dataService services.DataService, log *zap.SugaredLogger) *SObjectController {
	return &SObjectController{
		dataService,
		log,
	}
}