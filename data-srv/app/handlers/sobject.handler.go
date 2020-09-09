package handlers

import (
	"context"
	"github.com/objforce/objforce/data-srv/app/domain/services"
	"github.com/objforce/objforce/data-srv/app/dtos"
	data "github.com/objforce/objforce/data-srv/proto/data/gen-go"
	"github.com/xxxmicro/base/mapper"
)

type SObjectHandler struct {
	dataService services.DataService
}

func NewSObjectHandler(dataService services.DataService) *SObjectHandler {
	return &SObjectHandler{
		dataService,
	}
}

func (h *SObjectHandler) Create(c context.Context, req *data.SObject, rsp *data.SObject) error {
	dto := &dtos.SObject{}
	mapper.Map(req, dto)
	dto, err := h.dataService.Create(c, dto)
	if err != nil {
		return err
	}

	mapper.Map(rsp, dto)

	return nil
}

func (h *SObjectHandler) Update(c context.Context, req *data.SObject, rsp *data.SObject) error {
	dto := &dtos.SObject{}
	mapper.Map(req, dto)

	// req.Fields["a"] = any.Any{}

	dto, err := h.dataService.Update(c, dto)
	if err != nil {
		return err
	}

	mapper.Map(dto, rsp)

	return nil
}

func (h *SObjectHandler) FindOne(c context.Context, req *data.FindSObjectRequest, rsp *data.SObject) error {
	dto, err := h.dataService.FindOne(c, req.Id)
	if err != nil {
		return nil
	}

	mapper.Map(dto, rsp)

	return nil
}

func (h *SObjectHandler) Delete(c context.Context, req *data.DeleteSObjectRequest, rsp *data.SObject) error {
	err := h.dataService.Delete(c, req.Id)

	return err
}



