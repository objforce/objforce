package services

import (
	"context"

	"github.com/objforce/objforce/api/meta/app/dtos"
	"github.com/objforce/objforce/api/meta/app/models"
	"github.com/objforce/objforce/proto/meta"
)

type CustomFieldService interface {
	Create(c context.Context, customField *models.CustomField) (*dtos.CustomField, error)
	Update(c context.Context, customField *models.CustomField) (*dtos.CustomField, error)
	Get(c context.Context, id string) (*dtos.CustomField, error)
	Delete(c context.Context, id string) error
}

type customFieldService struct {
	customFieldServiceClient meta.CustomFieldService
}

func NewCustomFieldService(customFieldServiceClient meta.CustomFieldService) CustomFieldService {
	return &customFieldService{
		customFieldServiceClient,
	}
}

func (s *customFieldService) Create(c context.Context, model *models.CustomField) (*models.CustomField, error) {
	pb := &meta.CustomField{}
	mapstructure.Decode(model, pb)

	rsp, err := s.customFieldServiceClient.Create(c, pb)
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(rsp, model)

	return model, nil
}

func (s *customFieldService) Update(c context.Context, model *models.CustomField) (*models.CustomField, error) {
	pb := &meta.CustomField{}
	mapstructure.Decode(model, pb)

	rsp, err := s.customFieldServiceClient.Update(c, pb)
	if err != nil {
		return nil, err
	}

	mapstructure.Decode(rsp, model)

	return model, nil
}

func (s *customFieldService) Get(c context.Context, id string) (*dtos.CustomField, error) {
	rsp, err := s.customFieldServiceClient.Get(c, &meta.GetCustomFieldRequest{FieldId: id})
	if err != nil {
		return nil, err
	}

	model := &models.CustomField{}
	mapstructure.Decode(rsp, model)
	return model, nil
}

func (s *customFieldService) Delete(c context.Context, id string) error {
	_, err := s.customFieldServiceClient.Delete(c, &meta.DeleteCustomFieldRequest{FieldId: id})
	return err
}
