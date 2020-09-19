package handlers

import (
	"context"
	"github.com/micro/go-micro/v2/logger"
	"github.com/objforce/objforce/cmd/meta-srv/app/domain/services"
	"github.com/objforce/objforce/cmd/meta-srv/app/dtos"
	meta "github.com/objforce/objforce/idl/meta/gen-go"
	"github.com/xxxmicro/base/mapper"
)

type CustomFieldHandler struct {
	customFieldService services.CustomFieldService
}

func (h *CustomFieldHandler) Create(c context.Context, req *meta.CustomField, rsp *meta.CustomField) error {
	logger.Info("STARTING CustomFieldHandler.Create()")

	dto := &dtos.CustomField{}
	mapper.Map(req, dto)

	dto1, err := h.customFieldService.Create(c, dto)
	if err != nil {
		return err
	}

	mapper.Map(dto1, rsp)

	return nil
}

func (h *CustomFieldHandler) Delete(c context.Context, req *meta.DeleteCustomFieldRequest, rsp *meta.CustomField) error {
	logger.Info("STARTING CustomFieldHandler.Delete()")

	err := h.customFieldService.Delete(c, req.FieldId)
	if err != nil {
		return err
	}

	return nil
}

func (h *CustomFieldHandler) Update(c context.Context, req *meta.CustomField, rsp *meta.CustomField) error {
	dto := &dtos.CustomField{}
	mapper.Map(req, dto)

	dto, err := h.customFieldService.Update(c, dto)
	if err != nil {
		return err
	}

	mapper.Map(dto, rsp)

	return nil
}

func (h *CustomFieldHandler) Retrieve(c context.Context, req *meta.RetrieveCustomFieldRequest, rsp *meta.CustomField) error {
	dto, err := h.customFieldService.FindOne(c, req.FieldId)
	if err != nil {
		return err
	}

	mapper.Map(dto, rsp)
	return nil
}

func NewCustomFieldHandler(customFieldService services.CustomFieldService) *CustomFieldHandler {
	return &CustomFieldHandler{
		customFieldService,
	}
}