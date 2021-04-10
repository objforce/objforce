package handlers

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/goinggo/mapstructure"
	proto "github.com/objforce/objforce/proto/data"
	"github.com/objforce/objforce/service/data/domain/services"
	"github.com/objforce/objforce/service/data/models"
	"google.golang.org/protobuf/types/known/emptypb"
)

type SObjectHandler struct {
	dataService services.DataService
}

func NewSObjectHandler(dataService services.DataService) *SObjectHandler {
	return &SObjectHandler{
		dataService,
	}
}

func (h *SObjectHandler) Create(c context.Context, req *proto.CreateSObjectRequest, rsp *proto.SObject) error {
	object := &models.SObject{}
	mapstructure.Decode(req.Object, &object)

	newObject, err := h.dataService.Create(c, object)
	if err != nil {
		return err
	}

	mapstructure.Decode(newObject, rsp)

	return nil
}

func (h *SObjectHandler) Update(c context.Context, req *proto.UpdateSObjectRequest, rsp *proto.SObject) error {
	if req.Object == nil {
		return errors.New("parameters not valid, object is empty")
	}
	dto := req.Object

	var modelFields map[string]interface{}
	err := json.Unmarshal([]byte(dto.Fields), &modelFields)
	if err != nil {
		return err
	}

	updatedModel, err := h.dataService.Update(c, &models.SObject{
		Id:           dto.Id,
		OrgId:        dto.OrgId,
		Type:         dto.Type,
		FieldsToNull: dto.FieldsToNull,
		Fields:       modelFields,
	})
	if err != nil {
		return err
	}

	dtoFields, err := json.Marshal(updatedModel.Fields)
	if err != nil {
		return err
	}

	*rsp = proto.SObject{
		OrgId:        updatedModel.OrgId,
		Type:         updatedModel.Type,
		FieldsToNull: updatedModel.FieldsToNull,
		Id:           updatedModel.Id,
		Fields:       string(dtoFields),
	}

	return nil
}

func (h *SObjectHandler) Upsert(c context.Context, req *proto.UpsertSObjectRequest, rsp *proto.UpsertSObjectResponse) error {
	if req.Object == nil {
		return errors.New("parameters not valid, object is empty")
	}

	result, err := h.dataService.Upsert(c, sObjectFromDto(req.Object))

	if err != nil {
		return err
	}

	*rsp = proto.UpsertSObjectResponse{
		Created: result.Created,
		// Errors
		Id:      result.Id,
		Success: result.Success,
	}

	return nil
}

func (h *SObjectHandler) Get(c context.Context, req *proto.GetSObjectRequest, rsp *proto.SObject) error {
	model, err := h.dataService.Retrieve(c, req.OrgId, req.Type, req.Id, req.Fields)
	if err != nil {
		return err
	}

	if model == nil {
		return nil
	}

	dtoFields, err := json.Marshal(model.Fields)
	if err != nil {
		return err
	}

	*rsp = proto.SObject{
		OrgId:        model.OrgId,
		Type:         model.Type,
		FieldsToNull: model.FieldsToNull,
		Id:           model.Id,
		Fields:       string(dtoFields),
	}

	return nil
}

func (h *SObjectHandler) Delete(c context.Context, req *proto.DeleteSObjectRequest, _ *emptypb.Empty) error {
	err := h.dataService.Delete(c, req.OrgId, req.Type, req.Id)
	if err != nil {
		return err
	}

	return nil
}
