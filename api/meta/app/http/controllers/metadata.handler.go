package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/objforce/objforce/api/meta/app/dtos"
)

func (c *APIController) CreateMetadata(ctx *gin.Context) {
	// 根据 metadata fullName 转换具体类型

	customObject := &dtos.CustomObject{}
	if err := ctx.BindJSON(customObject); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

func (c *APIController) DescribeMetadata(ctx *gin.Context) {

}

func (c *APIController) ListMetadata(ctx *gin.Context) {

}

func (c *APIController) ReadMetadata(ctx *gin.Context) {

}

func (c *APIController) RenameMetadata(ctx *gin.Context) {

}

func (c *APIController) UpdateMetadata(ctx *gin.Context) {

}

func (c *APIController) UpsertMetadata(ctx *gin.Context) {

}

func (c *APIController) DeleteMetadata(ctx *gin.Context) {

}
