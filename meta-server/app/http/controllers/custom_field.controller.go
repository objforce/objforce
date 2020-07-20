package controllers

import(
	"net/http"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"github.com/objforce/meta-server/app/models"
	"github.com/objforce/meta-server/app/services"
)

type CustomFieldController struct {
	customFieldService services.CustomFieldService
	log *zap.SugaredLogger 
}

func (c *CustomFieldController) Create(ctx *gin.Context) {
	c.log.Info("STARTING CustomFieldController.Create()")

	customField := models.CustomField{}

	if err := ctx.ShouldBindJSON(&customField); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	customField = c.customFieldService.Create(customField)

	ctx.JSON(http.StatusCreated, customField)
	c.log.Infow("FINISHED CustomFieldController.Create()",
		"id", customField.Id,
		"ObjId", customField.ObjId,
	)
}

func NewCustomFieldController(customFieldService services.CustomFieldService, log *zap.SugaredLogger) *CustomFieldController {
	return &CustomFieldController{
		customFieldService,
		log,
	}
}