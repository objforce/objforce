package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/api/data/domain/services"
	"github.com/objforce/objforce/service/data/models"
	"go.uber.org/zap"
)

type SObjectController struct {
	dataService services.DataService
	log         *zap.SugaredLogger
}

func (c *SObjectController) Create(ctx *gin.Context) {
	var object *models.SObject

	if err := ctx.ShouldBindJSON(object); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newObject, err := c.dataService.Create(ctx.Request.Context(), object)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newObject)
}

func (c *SObjectController) Update(ctx *gin.Context) {
	object := &models.SObject{}

	if err := ctx.ShouldBindJSON(object); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedObject, err := c.dataService.Update(ctx.Request.Context(), object)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, updatedObject)
}

func (c *SObjectController) Upsert(ctx *gin.Context) {
	object := &models.SObject{}

	if err := ctx.ShouldBindJSON(object); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	upsertResult, err := c.dataService.Upsert(ctx.Request.Context(), object)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dtoRsp)
}

func (c *SObjectController) Delete(ctx *gin.Context) {
	c.log.Info("STARTING SObjectController.Delete()")

	dtoReq := &dtos.DeleteSObjectRequest{}
	if err := ctx.ShouldBindJSON(dtoReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rsp, err := c.dataService.Delete(ctx.Request.Context(), dtoReq)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, rsp)
}

func NewSObjectController(dataService services.DataService, log *zap.SugaredLogger) *SObjectController {
	return &SObjectController{
		dataService,
		log,
	}
}
