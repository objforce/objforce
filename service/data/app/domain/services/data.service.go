package services

import (
	"context"

	mapset "github.com/deckarep/golang-set"
	"github.com/mitchellh/mapstructure"
	"github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/service/data/app/domain/models"
	"github.com/objforce/objforce/service/data/app/domain/repositories"
)

type DataService interface {
	Create(c context.Context, object *models.SObject) ([]*models.SaveResult, error)
	Update(c context.Context, object *models.SObject) ([]*models.SaveResult, error)
	Upsert(c context.Context, object *models.SObject) ([]*models.UpsertResult, error)
	Get(c context.Context, orgId, objType string, id string, fields []string) ([]*models.SObject, error)
	Delete(c context.Context, orgId, objType string, id string) error
}

type dataService struct {
	dataRepository      repositories.DataRepository
	customObjectService meta.CustomObjectService
}

func NewDataService(dataRepository repositories.DataRepository, customObjectService meta.CustomObjectService) DataService {
	return &dataService{
		dataRepository,
		customObjectService,
	}
}

func (s *dataService) Create(c context.Context, object *models.SObject) (*models.SObject, error) {
	objId := calculateObjId(object.OrgId, object.Type)

	metaObj, err := s.customObjectService.Get(c, &meta.GetCustomObjectRequest{ObjId: objId})
	if err != nil {
		return nil, err
	}

	// TODO 存储的约束

	entity := &models.MTData{}
	mapstructure.Decode(dto, entity)
	entity.ObjId = metaObj.ObjId

	success := true
	err = s.dataRepository.Create(c, entity)
	if err != nil {
		success = false
	}

	newObject := &dtos.SObject{}
	mapstructure.Decode(entity, newObject)

	return newObject, nil
}

func (s *dataService) Update(c context.Context, object dtos.SObject) ([]*dtos.SaveResult, error) {
	objId := calculateObjId(dto.OrgId, object.Type)

	metaObj, err := s.customObjectService.Get(c, &meta.GetRequest{ObjId: objId})
	if err != nil {
		return nil, err
	}

	entity := &models.MTData{}
	mapstructure.Decode(object, entity)
	model.ObjId = metaObj.ObjId

	err = s.dataRepository.Update(c, model)
	if err != nil {
		return nil, err
	}

	updatedObject := &dtos.SObject{}
	mapstructure.Decode(entity, updatedObject)

	return updatedObject, nil
}

func (s *dataService) Get(c context.Context, orgId, objType string, id string, fields []string) (*models.SObject, error) {
	objId := calculateObjId(dto.OrgId, object.Type)

	objforceId, err := objflake.New(id)
	if err != nil {
		return nil, err
	}

	objType := objforceId.KeyPrefix

	metaObj, err := s.customObjectService.Get(c, &meta.GetRequest{ObjId: objId})
	if err != nil {
		return nil, err
	}

	entity := s.dataRepository.Get(c, orgId, objId, ids, fields)

	object := &dtos.SObject{}
	mapstructure.Decode(entity, object)

	return object, nil
}

func (s *dataService) Upsert(c context.Context, object *dtos.SObject) (*dtos.UpsertResult, error) {
	objId := calculateObjId(dto.OrgId, object.Type)

	metaObj, err := s.customObjectService.Get(c, &meta.GetRequest{ObjId: objId})
	if err != nil {
		return nil, err
	}

	model.ObjId = objId

	// 过滤无效列
	fieldSet := mapset.NewSet()
	for _, field := range metaObj.Fields {
		fieldSet.Add(field.FieldName)
	}

	for fieldName, _ := range model.Fields {
		if !fieldSet.Contains(fieldName) {
			delete(model.Fields, fieldName)
		}
	}

	entity := &models.MTData{}
	mapstructure.Decode(object, entity)

	result := s.dataRepository.Upsert(c, entity)

	var upsertResult *dtos.UpsertResult
	mapstructure.Map(result, upsertResult)

	return upsertResult, nil
}

func (s *dataService) Delete(c context.Context, orgId, objType string, id string) error {
	objId := calculateObjId(dto.OrgId, object.Type)

	err := s.dataRepository.Delete(c, orgId, objId, id)
	if err != nil {
		return nil, err
	}

	return nil
}
