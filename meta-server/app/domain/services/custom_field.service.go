package services

import(
	"context"
	"github.com/objforce/objforce/meta-server/app/domain/models"
	"github.com/objforce/objforce/meta-server/app/domain/repositories"
	"github.com/objforce/objforce/meta-server/app/dtos"
	"github.com/xxxmicro/base/mapper"
)

type CustomFieldService interface {
	Create(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error)
	Update(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error)
	FindOne(c context.Context, id string) (*dtos.CustomField, error)
	Delete(c context.Context, id string) error
}

type customFieldService struct {
	customFieldRepository repositories.CustomFieldRepository
}

func NewCustomFieldService(customFieldRepository repositories.CustomFieldRepository) CustomFieldService {
	return &customFieldService{
		customFieldRepository,
	}
}

func (s *customFieldService) Create(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error) {
	entity := &models.CustomField{}
	mapper.Map(dto, entity)

	err := s.customFieldRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *customFieldService) Update(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error) {
	entity := &models.CustomField{FieldId: dto.FieldId}

	mapper.Map(dto, entity)

	err := s.customFieldRepository.Update(c, entity, entity)
	if err != nil {
		return nil, err
	}

	err = s.customFieldRepository.FindOne(c, entity)
	if err != nil {
		return nil, err
	}

	mapper.Map(entity, dto)

	return dto, nil
}

func (s *customFieldService) FindOne(c context.Context, id string) (*dtos.CustomField, error) {
	entity := &models.CustomField{FieldId: id}
	s.customFieldRepository.FindOne(c, entity)

	dto := &dtos.CustomField{}
	mapper.Map(entity, dto)
	return dto, nil
}

func (s *customFieldService) Delete(c context.Context, id string) error {
	return s.customFieldRepository.Delete(c, &models.CustomField{FieldId: id})
}