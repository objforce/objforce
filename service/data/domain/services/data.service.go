package services

import (
	"context"

	mapset "github.com/deckarep/golang-set"
	"github.com/objforce/objflake"
	"github.com/objforce/objforce/service/data/domain/entities"
	"github.com/objforce/objforce/service/data/domain/repositories"
	"github.com/objforce/objforce/service/data/models"
	meta_client "github.com/objforce/objforce/service/meta/client"
	meta_models "github.com/objforce/objforce/service/meta/models"
)

type DataService interface {
	Create(c context.Context, object *models.SObject) (*models.SObject, error)
	Retrieve(c context.Context, orgId, objType string, id string, fields []string) (*models.SObject, error)
	Update(c context.Context, object *models.SObject) (*models.SObject, error)
	Delete(c context.Context, orgId, objType string, id string) error
	Upsert(c context.Context, object *models.SObject) (*models.UpsertResult, error)
}

type dataService struct {
	idGenerator        *objflake.IDGenerator
	dataRepository     repositories.DataRepository
	customObjectClient meta_client.CustomObjectClient
}

func NewDataService(idGenerator *objflake.IDGenerator, dataRepository repositories.DataRepository, customObjectClient meta_client.CustomObjectClient) DataService {
	return &dataService{
		idGenerator,
		dataRepository,
		customObjectClient,
	}
}

func (s *dataService) Create(c context.Context, model *models.SObject) (*models.SObject, error) {
	customObject, err := s.customObjectClient.FindByObjName(c, model.OrgId, model.Type)
	if err != nil {
		return nil, err
	}

	// 过滤无效列
	fieldSet := mapset.NewSet()
	for _, field := range customObject.Fields {
		fieldSet.Add(field.FieldName)
	}

	var fieldValues map[string]interface{}
	for fieldName, fieldValue := range model.Fields {
		if fieldSet.Contains(fieldName) {
			fieldValues[fieldName] = fieldValue
		}
	}

	if err = s.dataRepository.Create(c, &entities.MTData{
		GUID:   model.Id,
		OrgId:  model.OrgId,
		ObjId:  customObject.ObjId,
		Fields: fieldValues,
	}); err != nil {
		return nil, err
	}

	return s.Retrieve(c, model.OrgId, customObject.ObjName, model.Id, nil)
}

func buildMTData(customObject *meta_models.CustomObject, model *models.SObject) (*entities.MTData, error) {
	// 过滤无效列
	fieldSet := mapset.NewSet()
	for _, field := range customObject.Fields {
		fieldSet.Add(field.FieldName)
	}

	var fieldValues map[string]interface{}
	for fieldName, fieldValue := range model.Fields {
		if fieldSet.Contains(fieldName) {
			fieldValues[fieldName] = fieldValue
		}
	}

	return &entities.MTData{
		GUID:   model.Id,
		OrgId:  model.OrgId,
		ObjId:  customObject.ObjId,
		Fields: fieldValues,
	}, nil
}

func (s *dataService) Update(c context.Context, model *models.SObject) (*models.SObject, error) {
	customObject, err := s.customObjectClient.FindByObjName(c, model.OrgId, model.Type)
	if err != nil {
		return nil, err
	}

	objId, err := objflake.New(customObject.ObjId)
	if err != nil {
		return nil, err
	}

	id, err := s.idGenerator.NextID(objId.KeyPrefix, objId.PodIdentifier)
	if err != nil {
		return nil, err
	}

	entity, err := buildMTData(customObject, model)
	if err != nil {
		return nil, err
	}

	if err = s.dataRepository.Update(c, entity); err != nil {
		return nil, err
	}

	return s.Retrieve(c, model.OrgId, customObject.ObjName, id, nil)
}

func (s *dataService) Retrieve(c context.Context, orgId, objName, id string, fields []string) (*models.SObject, error) {
	customObject, err := s.customObjectClient.FindByObjName(c, orgId, objName)
	if err != nil {
		return nil, err
	}

	entity, err := s.dataRepository.Retrieve(c, orgId, customObject.ObjId, id, fields)
	if err != nil {
		return nil, err
	}

	return &models.SObject{
		Id:     entity.GUID,
		OrgId:  entity.OrgId,
		Type:   customObject.ObjName,
		Fields: entity.Fields,
	}, nil
}

func (s *dataService) Upsert(c context.Context, model *models.SObject) (*models.UpsertResult, error) {
	customObject, err := s.customObjectClient.FindByObjName(c, model.OrgId, model.Type)
	if err != nil {
		return nil, err
	}

	// 过滤无效列
	fieldSet := mapset.NewSet()
	for _, field := range customObject.Fields {
		fieldSet.Add(field.FieldName)
	}

	for fieldName, _ := range model.Fields {
		if !fieldSet.Contains(fieldName) {
			delete(model.Fields, fieldName)
		}
	}

	entity, err := buildMTData(customObject, model)
	if err != nil {
		return nil, err
	}

	result := s.dataRepository.Upsert(c, entity)

	return &models.UpsertResult{
		Created: result.Created,
		Errors:  result.Errors,
		Id:      result.Id,
		Success: result.Success,
	}, nil
}

func (s *dataService) Delete(c context.Context, orgId, objName, id string) error {
	customObject, err := s.customObjectClient.FindByObjName(c, orgId, objName)
	if err != nil {
		return err
	}

	err = s.dataRepository.Delete(c, orgId, customObject.ObjId, id)
	if err != nil {
		return err
	}

	return nil
}
