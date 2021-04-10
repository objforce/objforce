package handlers

import (
	"encoding/json"

	proto "github.com/objforce/objforce/proto/data"
	"github.com/objforce/objforce/service/data/models"
)

func sObjectFromDto(dto *proto.SObject) *models.SObject {
	var modelFields map[string]interface{}
	_ = json.Unmarshal([]byte(dto.Fields), &modelFields)

	return &models.SObject{
		Id:           dto.Id,
		OrgId:        dto.OrgId,
		Type:         dto.Type,
		FieldsToNull: dto.FieldsToNull,
		Fields:       modelFields,
	}
}
