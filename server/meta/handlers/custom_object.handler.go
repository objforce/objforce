package handlers

import (
	"context"
	"errors"

	"github.com/goinggo/mapstructure"
	"github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/service/meta/domain/services"
	"github.com/objforce/objforce/service/meta/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CustomObjectHandler struct {
	customObjectService services.CustomObjectService
}

func (h *CustomObjectHandler) Create(c context.Context, req *meta.CreateCustomObjectRequest, rsp *meta.CustomObject) error {
	model := &models.CustomObject{}
	mapstructure.Decode(req, model)

	newModel, err := h.customObjectService.Create(c, model)
	if err != nil {
		return err
	}

	mapstructure.Decode(newModel, rsp)

	return nil
}

func (h *CustomObjectHandler) Update(c context.Context, req *meta.CustomObject, rsp *meta.CustomObject) error {
	model := &models.CustomObject{}
	mapstructure.Decode(req, model)

	updatedModel, err := h.customObjectService.Update(c, model)
	if err != nil {
		return err
	}

	mapstructure.Decode(updatedModel, rsp)

	return nil
}

func (h *CustomObjectHandler) Retrieve(c context.Context, req *meta.GetCustomObjectRequest, rsp *meta.CustomObject) error {
	model, err := h.customObjectService.Retrieve(c, req.ObjId)
	if err != nil {
		return err
	}

	mapstructure.Decode(model, rsp)

	return nil
}

func (h *CustomObjectHandler) Delete(c context.Context, req *meta.DeleteCustomObjectRequest, _ *emptypb.Empty) error {
	if len(req.ObjId) == 0 {
		return errors.New("params error")
	}

	err := h.customObjectService.Delete(c, req.ObjId)
	if err != nil {
		return err
	}

	return nil
}

func (h *CustomObjectHandler) FindByObjName(c context.Context, req *meta.FindByObjNameRequest, rsp *meta.CustomObject) error {
	model, err := h.customObjectService.FindByObjName(c, req.OrgId, req.ObjName)
	if err != nil {
		return err
	}

	mapstructure.Decode(model, rsp)
	return nil
}

func NewCustomObjectHandler(customObjectService services.CustomObjectService) *CustomObjectHandler {
	return &CustomObjectHandler{
		customObjectService,
	}
}
