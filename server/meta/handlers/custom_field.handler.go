package handlers

import (
	"context"

	"github.com/goinggo/mapstructure"
	"github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/service/meta/domain/services"
	"github.com/objforce/objforce/service/meta/models"
)

type CustomFieldHandler struct {
	customFieldService services.CustomFieldService
}

func (h *CustomFieldHandler) Create(c context.Context, req *meta.CustomField, rsp *meta.CustomField) error {
	model := &models.CustomField{}
	mapstructure.Decode(req, model)

	newModel, err := h.customFieldService.Create(c, model)
	if err != nil {
		return err
	}

	mapstructure.Decode(newModel, rsp)

	return nil
}

func (h *CustomFieldHandler) Delete(c context.Context, req *meta.DeleteCustomFieldRequest, rsp *meta.CustomField) error {
	err := h.customFieldService.Delete(c, req.Id)
	if err != nil {
		return err
	}

	return nil
}

func (h *CustomFieldHandler) Update(c context.Context, req *meta.CustomField, rsp *meta.CustomField) error {
	model := &models.CustomField{}
	mapstructure.Decode(req, model)

	updatedModel, err := h.customFieldService.Update(c, model)
	if err != nil {
		return err
	}

	mapstructure.Decode(updatedModel, rsp)

	return nil
}

func (h *CustomFieldHandler) Get(c context.Context, req *meta.GetCustomFieldRequest, rsp *meta.CustomField) error {
	model, err := h.customFieldService.Get(c, req.Id)
	if err != nil {
		return err
	}

	mapstructure.Decode(model, rsp)
	return nil
}

func NewCustomFieldHandler(customFieldService services.CustomFieldService) *CustomFieldHandler {
	return &CustomFieldHandler{
		customFieldService,
	}
}
