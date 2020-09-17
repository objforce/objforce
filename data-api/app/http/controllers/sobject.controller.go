package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/data-api/app/domain/services"
	"github.com/objforce/objforce/data-api/app/dtos"
	"go.uber.org/zap"
	"net/http"
)

type SObjectController struct {
	dataService services.DataService
	log *zap.SugaredLogger 
}

func (c *SObjectController) Create(ctx *gin.Context) {
	c.log.Info("STARTING CustomFieldController.Create()")

	var objects []*dtos.SObject

	if err := ctx.ShouldBindJSON(&objects); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	dtoRsp, err := c.dataService.Create(ctx.Request.Context(), objects)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, dtoRsp)
}


func (c *SObjectController) Update(ctx *gin.Context) {
	c.log.Info("STARTING CustomFieldController.Update()")

	dtoReq := &dtos.UpdateSObjectRequest{}

	if err := ctx.ShouldBindJSON(dtoReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dtoRsp, err := c.dataService.Update(ctx.Request.Context(), dtoReq)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dtoRsp)
}

func (c *SObjectController) Upsert(ctx *gin.Context) {
	c.log.Info("STARTING CustomFieldController.Upsert()")

	dtoReq := &dtos.UpsertSObjectRequest{}

	if err := ctx.ShouldBindJSON(dtoReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dtoRsp, err := c.dataService.Upsert(ctx.Request.Context(), dtoReq)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, dtoRsp)
}

func (c *SObjectController) Delete(ctx *gin.Context) {
	c.log.Info("STARTING CustomFieldController.Delete()")

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