package handlers

import (
	"context"
	"github.com/objforce/objforce/data-srv/app/domain/services"
	"github.com/objforce/objforce/data-srv/app/dtos"
	data "github.com/objforce/objforce/idl/data/gen-go"
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

func (h *SObjectHandler) Create(c context.Context, req *data.CreateSObjectRequest, rsp *data.CreateSObjectResponse) error {
	// TODO validator判断 req.Objects 为空

	dtoList := make([]*dtos.SObject, len(req.Objects))
	for i, pb := range req.Objects {
		dto := &dtos.SObject{}
		mapper.Map(pb, dto)
		dtoList[i] = dto
	}

	dtoResults, err := h.dataService.Create(c, dtoList)
	if err != nil {
		return err
	}

	pbResults := make([]*data.SaveResult, len(dtoResults))
	for i, dtoResult := range dtoResults {
		var err string
		if dtoResult.Error != nil {
			err = dtoResult.Error.Error()
		}
		pbResult := &data.SaveResult{
			Error: err,
			Id: dtoResult.Id,
			Success: dtoResult.Success,
		}
		pbResults[i] = pbResult
	}

	rsp.Results = pbResults

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

func (h *SObjectHandler) Retrieve(c context.Context, req *data.RetrieveSObjectRequest, rsp *data.RetrieveSObjectResponse) error {
	dtoResults, err := h.dataService.Retrieve(c, req.OrgId, req.Type, req.Ids, req.Fields)
	if err != nil {
		return nil
	}

	pbResults := make([]*data.SObject, len(dtoResults))
	for i, dtoResult := range dtoResults {
		pbResults[i] = &data.SObject{
			Type: dtoResult.Type,
			FieldsToNull: dtoResult.FieldsToNull,
			Id: dtoResult.Id,
			Fields: dtoResult.Fields,
		}
	}

	rsp.Results = pbResults
	return nil
}

func (h *SObjectHandler) Delete(c context.Context, req *data.DeleteSObjectRequest, rsp *data.DeleteSObjectResponse) error {
	dtoResults, err := h.dataService.Delete(c, req.OrgId, req.Type, req.Ids)
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
			Error: error,
			Id: dtoResult.Id,
			Success: dtoResult.Success,
		}
	}

	rsp.Results = pbResults

	return nil
}



