package services

import (
	"context"
	"github.com/objforce/objforce/cmd/data-srv/app/domain/models"
	"github.com/objforce/objforce/cmd/data-srv/app/domain/repositories"
	"github.com/objforce/objforce/cmd/data-srv/app/dtos"
	meta "github.com/objforce/objforce/idl/meta/gen-go"
	"github.com/xxxmicro/base/mapper"
	mapset "github.com/deckarep/golang-set"
)

type DataService interface {
	Create(c context.Context, items []*dtos.SObject) ([]*dtos.SaveResult, error)
	Update(c context.Context, items []*dtos.SObject) ([]*dtos.SaveResult, error)
	Upsert(c context.Context, items []*dtos.SObject) ([]*dtos.UpsertResult, error)
	Retrieve(c context.Context, orgId, objType string, ids []string, fields []string) ([]*dtos.SObject, error)
	Delete(c context.Context, orgId, objType string, ids []string) ([]*dtos.DeleteResult, error)
}

type dataService struct {
	dataRepository repositories.DataRepository
	customObjectService meta.CustomObjectService
}

func NewDataService(dataRepository repositories.DataRepository, customObjectService meta.CustomObjectService) DataService {
	return &dataService{
		dataRepository,
		customObjectService,
	}
}

func (s *dataService) Create(c context.Context, dtoList []*dtos.SObject) ([]*dtos.SaveResult, error) {
	dtoResults := make([]*dtos.SaveResult, len(dtoList))

	for i, dto := range dtoList {
		metaObj, err := s.customObjectService.FindCustomObjectByOrgAndType(c, &meta.OrgAndObjTypeRequest{OrgId: dto.OrgId, ObjType: dto.Type})
		if err != nil {
			dtoResults[i] = &dtos.SaveResult{
				Error: err,
				Success: false,
			}
			continue
		}

		model := &models.MTData{}
		mapper.Map(dto, model)
		model.ObjId = metaObj.ObjId
		success := true
		err = s.dataRepository.Create(c, model)
		if err != nil {
			success = false
		}

		dtoResults[i] = &dtos.SaveResult{
			Error: err,
			Success: success,
			Id: model.GUID,
		}
	}

	return dtoResults, nil
}

func (s *dataService) Update(c context.Context, dtoList []*dtos.SObject) ([]*dtos.SaveResult, error) {
	dtoResults := make([]*dtos.SaveResult, len(dtoList))

	for i, dto := range dtoList {
		metaObj, err := s.customObjectService.FindCustomObjectByOrgAndType(c, &meta.OrgAndObjTypeRequest{OrgId: dto.OrgId, ObjType: dto.Type})
		if err != nil {
			dtoResults[i] = &dtos.SaveResult{
				Error: err,
				Success: false,
			}
			continue
		}

		model := &models.MTData{}
		mapper.Map(dto, model)
		model.ObjId = metaObj.ObjId

		success := true
		err = s.dataRepository.Update(c, model)
		if err != nil {
			success = false
		}

		dtoResults[i] = &dtos.SaveResult{
			Error: err,
			Success: success,
			Id: model.GUID,
		}
	}

	return dtoResults, nil
}

func (s *dataService) Retrieve(c context.Context, orgId, objType string, ids []string, fields []string) ([]*dtos.SObject, error) {
	metaObj, err := s.customObjectService.FindCustomObjectByOrgAndType(c, &meta.OrgAndObjTypeRequest{OrgId: orgId, ObjType: objType})
	if err != nil {
		return nil, err
	}

	objId := metaObj.ObjId

	families := map[string][]string{
		"basic": {
			"created_at", "created_by", "updated_at", "updated_by",
		},
		"ext": fields,
	}

	modelList := s.dataRepository.MultiGet(c, orgId, objId, ids, families)

	dtoList := make([]*dtos.SObject, len(modelList))
	mapper.Map(modelList, dtoList)

	return dtoList, nil
}

func (s *dataService) Upsert(c context.Context, dtoList []*dtos.SObject) ([]*dtos.UpsertResult, error) {
	modelList := make([]*models.MTData, len(dtoList))
	mapper.Map(dtoList, modelList)
	for i, dto := range dtoList {
		// 通过Org和type得到 对象元数据, 进而得到 objId
		metaObj, err := s.customObjectService.FindCustomObjectByOrgAndType(c, &meta.OrgAndObjTypeRequest{OrgId: dto.OrgId, ObjType: dto.Type})
		if err != nil {
			return nil, err
		}
		objId := metaObj.ObjId

		model := modelList[i]
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
	}

	results := s.dataRepository.MultiUpsert(c, modelList)

	dtoResults := make([]*dtos.UpsertResult, len(results))
	mapper.Map(results, dtoResults)

	return dtoResults, nil
}

func (s *dataService) Delete(c context.Context, orgId, objType string, ids []string) ([]*dtos.DeleteResult, error) {
	metaObj, err := s.customObjectService.FindCustomObjectByOrgAndType(c, &meta.OrgAndObjTypeRequest{OrgId: orgId, ObjType: objType})
	if err != nil {
		return nil, err
	}
	objId := metaObj.ObjId

	modelResults := s.dataRepository.MultiDelete(c, orgId, objId, ids)

	dtoResults := make([]*dtos.DeleteResult, len(modelResults))
	mapper.Map(modelResults, dtoResults)

	return dtoResults, nil
}