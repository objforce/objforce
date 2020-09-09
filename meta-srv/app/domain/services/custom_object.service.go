package services

import (
	"context"
	"github.com/micro/go-micro/v2/client"
	"github.com/objforce/objforce/meta-srv/app/domain/models"
	"github.com/objforce/objforce/meta-srv/app/domain/repositories"
	"github.com/objforce/objforce/meta-srv/app/dtos"
	"github.com/objforce/objforce/meta-srv/config"
	meta "github.com/objforce/objforce/meta-srv/proto/meta/gen-go"
	"github.com/xxxmicro/base/mapper"
)

type CustomObjectService interface {
	Create(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	FindOne(c context.Context, id string) (*dtos.CustomObject, error)
	Delete(c context.Context, id string) error
}

type customObjectService struct {
	client client.Client
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
	entity := &models.CustomField{}
	mapper.Map(dto, entity)

	err := s.customObjectRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *customObjectService) Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	entity := &models.CustomField{ObjId: dto.Id}

	mapper.Map(dto, entity)

	err := s.customObjectRepository.Update(c, entity, entity)
	if err != nil {
		return nil, err
	}

	err = s.customObjectRepository.FindOne(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *customObjectService) FindOne(c context.Context, id string) (*dtos.CustomObject, error) {
	entity := &models.CustomObject{ObjId: id}
	s.customObjectRepository.FindOne(c, entity)

	dto := &dtos.CustomObject{}
	mapper.Map(entity, dto)
	return dto, nil
}

func (s *customObjectService) Delete(c context.Context, id string) error {
	return s.customObjectRepository.Delete(c, &models.CustomObject{ObjId: id})
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

func (s *CustomObjectServiceEventWrapper) FindOne(c context.Context, id string) (*dtos.CustomObject, error) {
	return s.ref.FindOne(c, id)
}
func (s *CustomObjectServiceEventWrapper) Delete(c context.Context, id string) error {
	dto, err := s.ref.FindOne(c, id)
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
