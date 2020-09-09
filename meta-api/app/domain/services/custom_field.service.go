package services

import(
	"context"
	"github.com/objforce/objforce/meta-api/app/dtos"
	meta "github.com/objforce/objforce/meta-api/proto/meta/gen-go"
	"github.com/xxxmicro/base/mapper"
)

type CustomFieldService interface {
	Create(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error)
	Update(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error)
	FindOne(c context.Context, id string) (*dtos.CustomField, error)
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

func (s *customFieldService) Create(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error) {
	pb := &meta.CustomField{}
	mapper.Map(dto, pb)

	rsp, err := s.customFieldServiceClient.Create(c, pb)
	if err != nil {
		return nil, err
	}

	mapper.Map(rsp, dto)

	return dto, nil
}

func (s *customFieldService) Update(c context.Context, dto *dtos.CustomField) (*dtos.CustomField, error) {
	pb := &meta.CustomField{}
	mapper.Map(dto, pb)

	rsp, err := s.customFieldServiceClient.Update(c, pb)
	if err != nil {
		return nil, err
	}

	mapper.Map(rsp, dto)

	return dto, nil
}

func (s *customFieldService) FindOne(c context.Context, id string) (*dtos.CustomField, error) {
	rsp, err := s.customFieldServiceClient.FindOne(c, &meta.FindCustomFieldRequest{FieldId: id})
	if err != nil {
		return nil, err
	}

	dto := &dtos.CustomField{}
	mapper.Map(rsp, dto)
	return dto, nil
}

func (s *customFieldService) Delete(c context.Context, id string) error {
	_, err :=  s.customFieldServiceClient.Delete(c, &meta.DeleteCustomFieldRequest{FieldId: id})
	return err
}