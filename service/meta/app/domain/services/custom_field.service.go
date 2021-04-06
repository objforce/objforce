package services

import (
	"context"

	"github.com/goinggo/mapstructure"
	"github.com/objforce/objforce/service/meta/app/domain/entities"
	"github.com/objforce/objforce/service/meta/app/domain/repositories"
	"github.com/objforce/objforce/service/meta/app/models"
)

type CustomFieldService interface {
	Create(c context.Context, customField *models.CustomField) (*models.CustomField, error)
	Update(c context.Context, customField *models.CustomField) (*models.CustomField, error)
	Get(c context.Context, id string) (*models.CustomField, error)
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

func (s *customFieldService) Create(c context.Context, model *models.CustomField) (*models.CustomField, error) {
	entity := &entities.MTField{}
	mapstructure.Decode(model, &entity)

	err := s.customFieldRepository.Create(c, entity)
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(entity, model)

	return model, nil
}

func (s *customFieldService) Update(c context.Context, model *models.CustomField) (*models.CustomField, error) {
	entity := &entities.MTField{FieldId: model.FieldId}

	mapstructure.Decode(model, &entity)

	err := s.customFieldRepository.Update(c, entity, entity)
	if err != nil {
		return nil, err
	}

	err = s.customFieldRepository.Get(c, entity)
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(entity, &model)

	return model, nil
}

func (s *customFieldService) Get(c context.Context, id string) (*models.CustomField, error) {
	entity := &entities.MTField{FieldId: id}
	s.customFieldRepository.Get(c, entity)

	model := &models.CustomField{}
	mapstructure.Decode(entity, model)
	return model, nil
}

func (s *customFieldService) Delete(c context.Context, id string) error {
	return s.customFieldRepository.Delete(c, &entities.MTField{FieldId: id})
}
