package services

import (
	"context"

	"github.com/objforce/objforce/service/index/app/domain/models"
	"github.com/objforce/objforce/service/index/app/domain/repositories"
	"github.com/objforce/objforce/service/index/app/dtos"
	"github.com/xxxmicro/base/mapper"
)

type DocumentService interface {
	Update(c context.Context, dto *dtos.Document) error
	Upsert(c context.Context, dto *dtos.Document) error
	Retrieve(c context.Context, objId, id string) (*dtos.Document, error)
	Delete(c context.Context, objId, id string) error
}

type documentService struct {
	documentRepository repositories.DocumentRepository
}

func NewDocumentService(indexRepository repositories.DocumentRepository) DocumentService {
	return &documentService{
		indexRepository,
	}
}

func (s *documentService) Upsert(c context.Context, dto *dtos.Document) error {
	m := &models.Document{}
	mapper.Map(dto, m)
	return s.documentRepository.Upsert(c, m)
}

func (s *documentService) Update(c context.Context, dto *dtos.Document) error {
	m := &models.Document{}
	mapper.Map(dto, m)
	return s.documentRepository.Update(c, m)
}

func (s *documentService) Retrieve(c context.Context, objId, id string) (*dtos.Document, error) {
	model, err := s.documentRepository.Retrieve(c, objId, "doc", id)

	dto := &dtos.Document{}
	mapper.Map(model, dto)
	return dto, err
}

func (s *documentService) Delete(c context.Context, objId, id string) error {
	return s.documentRepository.Delete(c, objId, "doc", id)
}
