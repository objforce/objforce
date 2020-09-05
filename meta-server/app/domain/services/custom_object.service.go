package services

import (
	"context"
	"github.com/objforce/meta-server/app/domain/models"
	"github.com/objforce/meta-server/app/domain/repositories"
	"github.com/objforce/meta-server/app/dtos"
	"github.com/objforce/meta-server/app/mapper"
	xmapper "github.com/devfeel/mapper"
)

type CustomObjectService interface {
	Create(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	FindOne(c context.Context, id string) (*dtos.CustomObject, error)
	Delete(c context.Context, id string) error
}

type customObjectService struct {
	customObjectRepository repositories.CustomObjectRepository
}

func NewCustomObjectService(customObjectRepository repositories.CustomObjectRepository) CustomObjectService {
	return &customObjectService{
		customObjectRepository,
	}
}

func (s *customObjectService) Create(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	entity := &models.CustomField{}
	mapper.CUSTOMER_OBJECT_MAPPER.ConvertToEntity(dto, entity)

	err := s.customObjectRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.CUSTOMER_OBJECT_MAPPER.ConvertToDto(entity, dto)

	return dto, nil
}

func (s *customObjectService) Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	entity := &models.CustomField{Id: dto.Id}

	var change map[string]interface{}
	xmapper.Mapper(dto, change)

	err := s.customObjectRepository.Update(c, entity, change)
	if err != nil {
		return nil, err
	}

	err = s.customObjectRepository.FindOne(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.CUSTOMER_OBJECT_MAPPER.ConvertToDto(entity, dto)

	return dto, nil
}

func (s *customObjectService) FindOne(c context.Context, id string) (*dtos.CustomObject, error) {
	entity := &models.CustomObject{Id: id}
	s.customObjectRepository.FindOne(c, entity)

	dto := &dtos.CustomObject{}
	mapper.CUSTOMER_OBJECT_MAPPER.ConvertToDto(entity, dto)
	return dto, nil
}

func (s *customObjectService) Delete(c context.Context, id string) error {
	return s.customObjectRepository.Delete(c, &models.CustomObject{Id: id})
}