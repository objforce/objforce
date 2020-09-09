package handlers

import (
	"context"
	"errors"
	"github.com/micro/go-micro/v2/logger"
	"github.com/objforce/objforce/meta-srv/app/domain/services"
	"github.com/objforce/objforce/meta-srv/app/dtos"
	meta "github.com/objforce/objforce/meta-srv/proto/meta/gen-go"
	"github.com/xxxmicro/base/mapper"
)

type CustomObjectHandler struct {
	customObjectService services.CustomObjectService
}

func (h *CustomObjectHandler) Create(c context.Context, req *meta.CustomObject, rsp *meta.CustomObject) error {
	logger.Info("STARTING CustomObjectHandler.Create()")

	dto := &dtos.CustomObject{}
	mapper.Map(req, dto)

	dto1, err := h.customObjectService.Create(c, dto)
	if err != nil {
		return err
	}

	mapper.Map(dto1, rsp)

	return nil
}

func (h *CustomObjectHandler) Delete(c context.Context, req *meta.DeleteCustomObjectRequest, rsp *meta.CustomObject) error {
	logger.Info("STARTING CustomFieldController.Delete()")

	if len(req.ObjId) == 0 {
		return errors.New("params error")
	}

	err := h.customObjectService.Delete(c, req.ObjId)
	if err != nil {
		return err
	}

	return nil
}

func (h *CustomObjectHandler) Update(c context.Context, req *meta.CustomObject, rsp *meta.CustomObject) error {
	dto := &dtos.CustomObject{}
	mapper.Map(req, dto)

	dto1, err := h.customObjectService.Update(c, dto)
	if err != nil {
		return err
	}

	mapper.Map(dto1, rsp)

	return nil
}

func (h *CustomObjectHandler) FindOne(c context.Context, req *meta.FindCustomObjectRequest, rsp *meta.CustomObject) error {
	dto, err := h.customObjectService.FindOne(c, req.ObjId)
	if err != nil {
		return err
	}

	mapper.Map(dto, rsp)

	return nil
}


func NewCustomObjectHandler(customObjectService services.CustomObjectService) *CustomObjectHandler {
	return &CustomObjectHandler{
		customObjectService,
	}
}