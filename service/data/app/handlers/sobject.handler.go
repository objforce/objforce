package handlers

import (
	"context"

	data "github.com/objforce/objforce/proto/data"
	"github.com/objforce/objforce/service/data/app/domain/services"
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

func (h *SObjectHandler) Create(c context.Context, req *data.CreateSObjectRequest, rsp *data.SObject) error {
	object := &dtos.SObject{}
	mapstructure.Decode(req.Object, &object)

	newObject, err := h.dataService.Create(c, object)
	if err != nil {
		return err
	}

	mapstructure.Decode(newObject, rsp)

	return nil
}

func (h *SObjectHandler) Update(c context.Context, req *data.UpdateSObjectRequest, rsp *data.UpdateSObjectResponse) error {
	// TODO validator判断 req.Objects 为空

	dtoList := make([]*dtos.SObject, len(req.Objects))
	mapper.Map(req.Objects, dtoList)
	dtoResults, err := h.dataService.Update(c, dtoList)
	if err != nil {
		return err
	}

	pbResults := make([]*data.SaveResult, len(dtoResults))
	mapper.Map(dtoResults, pbResults)

	rsp.Results = pbResults

	return nil
}

func (h *SObjectHandler) Upsert(c context.Context, req *data.UpsertSObjectRequest, rsp *data.UpsertSObjectResponse) error {
	// TODO validator判断 req.Objects 为空

	dtoList := make([]*dtos.SObject, len(req.Objects))
	mapper.Map(req.Objects, dtoList)
	dtoResults, err := h.dataService.Upsert(c, dtoList)
	if err != nil {
		return err
	}

	pbResults := make([]*data.UpsertSObjectResult, len(dtoResults))
	mapper.Map(dtoResults, pbResults)

	rsp.Results = pbResults

	return nil
}

func (h *SObjectHandler) Get(c context.Context, req *data.GetSObjectRequest, rsp *data.SObject) error {
	object, err := h.dataService.Get(c, req.OrgId, req.Type, req.Ids, req.Fields)
	if err != nil {
		return nil
	}

	mapstructure.Decode(object, *rsp)

	return nil
}

func (h *SObjectHandler) Delete(c context.Context, req *data.DeleteSObjectRequest, rsp *data.DeleteSObjectResponse) error {
	dtoResults, err := h.dataService.Delete(c, req.OrgId, req.Type, req.Id)
	if err != nil {
		return err
	}

	pbResults := make([]*data.DeleteResult, len(dtoResults))
	for i, dtoResult := range dtoResults {
		var error string
		if dtoResult.Error != nil {
			error = dtoResult.Error.Error()
		}

		pbResults[i] = &data.DeleteResult{
			Error:   error,
			Id:      dtoResult.Id,
			Success: dtoResult.Success,
		}
	}

	rsp.Results = pbResults

	return nil
}
