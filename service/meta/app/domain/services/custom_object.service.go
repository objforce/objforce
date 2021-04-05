package services

import (
	"context"

	"github.com/micro/go-micro/v2/client"
	"github.com/objforce/objforce/proto/meta"
	"github.com/objforce/objforce/service/meta/app/domain/models"
	"github.com/objforce/objforce/service/meta/app/domain/repositories"
	"github.com/objforce/objforce/service/meta/app/dtos"
	"github.com/objforce/objforce/service/meta/config"
	"github.com/xxxmicro/base/mapper"
)

type CustomObjectService interface {
	Create(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	Upsert(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	Retrieve(c context.Context, id string) (*dtos.CustomObject, error)
	Delete(c context.Context, id string) error
	FindCustomObjectByOrgAndType(c context.Context, orgId, objType string) (*dtos.CustomObject, error)
}

type customObjectService struct {
	client                 client.Client
	customObjectRepository repositories.CustomObjectRepository
}

func NewCustomObjectService(client client.Client, customObjectRepository repositories.CustomObjectRepository) CustomObjectService {
	s := &customObjectService{
		client,
		customObjectRepository,
	}

	s1 := CustomObjectServiceEventWrapper{}
	s1.Wrap(s)
	return s
}

func (s *customObjectService) Create(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	entity := &models.MTObject{}
	mapstructure.Decode(dto, entity)

	err := s.customObjectRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *customObjectService) Upsert(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	entity := &models.MTObject{ObjId: dto.ObjId}

	mapper.Map(dto, entity)

	if err := s.customObjectRepository.Upsert(c, entity); err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *customObjectService) Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	entity := &models.MTObject{ObjId: dto.ObjId}

	mapper.Map(dto, entity)

	if err := s.customObjectRepository.Update(c, entity); err != nil {
		return nil, err
	}

	entity, err := s.customObjectRepository.Retrieve(c, dto.ObjId)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *customObjectService) Retrieve(c context.Context, ObjId string) (*dtos.CustomObject, error) {
	model, err := s.customObjectRepository.Retrieve(c, ObjId)
	if err != nil {
		return nil, err
	}

	dto := &dtos.CustomObject{}
	mapper.Map(model, dto)
	return dto, nil
}

func (s *customObjectService) FindCustomObjectByOrgAndType(c context.Context, orgId string, objType string) (*dtos.CustomObject, error) {
	model, err := s.customObjectRepository.FindCustomObjectByOrgAndType(c, orgId, objType)
	if err != nil {
		return nil, err
	}

	dto := &dtos.CustomObject{}
	mapper.Map(model, dto)
	return dto, nil
}

func (s *customObjectService) Delete(c context.Context, objId string) error {
	return s.customObjectRepository.Delete(c, objId)
}

type CustomObjectServiceEventWrapper struct {
	ref CustomObjectService
}

func (s *CustomObjectServiceEventWrapper) Wrap(ref CustomObjectService) {
	s.ref = ref
}

func (s *CustomObjectServiceEventWrapper) Create(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	dto, err := s.ref.Create(c, dto)
	if err != nil {
		return nil, err
	}

	pb := &meta.CustomObject{}
	mapper.Map(dto, pb)
	err = client.Publish(c, client.NewMessage(config.CustomObjectCreatedTopic, pb))
	if err != nil {
		return nil, err
	}

	return dto, nil
}

func (s *CustomObjectServiceEventWrapper) Upsert(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	dto, err := s.ref.Upsert(c, dto)
	if err != nil {
		return nil, err
	}

	pb := &meta.CustomObject{}
	mapper.Map(dto, pb)
	err = client.Publish(c, client.NewMessage(config.CustomObjectUpsertedTopic, pb))
	if err != nil {
		return dto, err
	}

	return dto, err
}

func (s *CustomObjectServiceEventWrapper) Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	dto, err := s.ref.Update(c, dto)
	if err != nil {
		return nil, err
	}

	pb := &meta.CustomObject{}
	mapper.Map(dto, pb)
	err = client.Publish(c, client.NewMessage(config.CustomObjectUpdatedTopic, pb))
	if err != nil {
		return dto, err
	}
	return dto, err
}

func (s *CustomObjectServiceEventWrapper) Retrieve(c context.Context, id string) (*dtos.CustomObject, error) {
	return s.ref.Retrieve(c, id)
}
func (s *CustomObjectServiceEventWrapper) Delete(c context.Context, id string) error {
	dto, err := s.ref.Retrieve(c, id)
	if err != nil {
		return err
	}

	err = s.Delete(c, id)
	if err != nil {
		return err
	}

	pb := &meta.CustomObject{}
	mapper.Map(dto, pb)
	err = client.Publish(c, client.NewMessage(config.CustomObjectDeletedTopic, pb))
	if err != nil {
		return err
	}

	return nil
}
