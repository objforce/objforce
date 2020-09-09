package handlers

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/v2/logger"
	"github.com/objforce/objforce/meta-srv/app/domain/services"
	"github.com/objforce/objforce/meta-srv/app/dtos"
	meta "github.com/objforce/objforce/idl/meta/gen-go"
	"net/http"
)

type CustomObjectHandler struct {
	customFieldService services.CustomFieldService
}

func (c *CustomObjectHandler) Create(c context.Context, ) {
	logger.Info("STARTING CustomFieldController.Create()")

	customField := &dtos.CustomField{}

	if err := ctx.ShouldBindJSON(&customField); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customField, err := c.customFieldService.Create(ctx.Request.Context(), customField)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, customField)
	c.log.Infow("FINISHED CustomFieldController.Create()",
		"FieldId", customField.FieldId,
		"ObjId", customField.ObjId,
	)
}

func (c *CustomFieldController) Delete(ctx *gin.Context) {
	c.log.Info("STARTING CustomFieldController.Delete()")

	id := ctx.Param("id")
	if len(id) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"msg": "need id"})
	}

	err := c.customFieldService.Delete(ctx.Request.Context(), id)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"code": 0})
}

func NewCustomFieldHandler(customFieldService services.CustomFieldService, log *zap.SugaredLogger) *CustomFieldController {
	return &CustomFieldController{
		customFieldService,
		log,
	}
}