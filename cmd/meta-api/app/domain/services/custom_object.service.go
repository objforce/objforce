package services

import (
	"context"
	"github.com/objforce/objforce/cmd/meta-api/app/dtos"
	meta "github.com/objforce/objforce/idl/meta/gen-go"
	"github.com/xxxmicro/base/mapper"
)

type CustomObjectService interface {
	Create(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error)
	FindOne(c context.Context, id string) (*dtos.CustomObject, error)
	Delete(c context.Context, id string) error
}

type customObjectService struct {
	customObjectServiceClient meta.CustomObjectService
}

func NewCustomObjectService(customObjectServiceClient meta.CustomObjectService) CustomObjectService {
	return &customObjectService{
		customObjectServiceClient,
	}
}

func (s *customObjectService) Create(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	pb := &meta.CustomObject{}
	mapper.Map(dto, pb)

	rsp, err := s.customObjectServiceClient.Create(c, pb)
	if err != nil {
		return nil, err
	}

	mapper.Map(rsp, dto)

	return dto, nil
}

func (s *customObjectService) Update(c context.Context, dto *dtos.CustomObject) (*dtos.CustomObject, error) {
	pb := &meta.CustomObject{}

	mapper.Map(dto, pb)

	rsp, err := s.customObjectServiceClient.Update(c, pb)
	if err != nil {
		return nil, err
	}

	mapper.Map(rsp, dto)

	return dto, nil
}

func (s *customObjectService) FindOne(c context.Context, id string) (*dtos.CustomObject, error) {
	pb, err := s.customObjectServiceClient.FindOne(c, &meta.FindCustomObjectRequest{ObjId: id})
	if err != nil {
		return nil, err
	}

	dto := &dtos.CustomObject{}
	mapper.Map(pb, dto)

	return dto, nil
}

func (s *customObjectService) Delete(c context.Context, id string) error {
	_, err := s.customObjectServiceClient.Delete(c, &meta.DeleteCustomObjectRequest{ObjId: id})
	return err
}