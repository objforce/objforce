package services

import(
	"context"
	"github.com/objforce/meta-server/app/domain/models"
	"github.com/objforce/meta-server/app/domain/repositories"
	"github.com/objforce/meta-server/app/dtos"
	"github.com/objforce/meta-server/app/mapper"
)

type customFieldService struct {
	customFieldRepository repositories.CustomFieldRepository
}

func (s *customFieldService) Create(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error) {
	entity := &models.CustomField{}
	mapper.CUSTOMER_FIELD_MAPPER.ConvertToEntity(dto, entity)

	err := s.customFieldRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.CUSTOMER_FIELD_MAPPER.ConvertToDto(entity, dto)

	return dto, nil
}

func (s *customFieldService) Update(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error) {
	entity := &models.CustomField{}
	mapper.CUSTOMER_FIELD_MAPPER.ConvertToEntity(dto, entity)

	err := s.customFieldRepository.Update(c, entity, entity)
	if err != nil {
		return nil, err
	}

	err = s.customFieldRepository.FindOne(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.CUSTOMER_FIELD_MAPPER.ConvertToDto(entity, dto)

	return dto, nil
}

func (s *customFieldService) FindOne(c context.Context, id string) (*dtos.CustomField, error) {
	entity := &models.CustomField{Id: id}
	s.customFieldRepository.FindOne(c, entity)

	dto := &dtos.CustomField{}
	mapper.CUSTOMER_FIELD_MAPPER.ConvertToDto(entity, dto)
	return dto, nil
}

func (s *customFieldService) Delete(c context.Context, id string) error {
	return s.customFieldRepository.Delete(c, &models.CustomField{Id: id})
}

type CustomFieldService interface {
	Create(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error)
	Update(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error)
	FindOne(c context.Context, id string) (*dtos.CustomField, error)
	Delete(c context.Context, id string) error
}

func NewCustomFieldService(customFieldRepository repositories.CustomFieldRepository) CustomFieldService {
	return &customFieldService{
		customFieldRepository,
	}
}