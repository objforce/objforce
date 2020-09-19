package services

import (
	"context"
	"github.com/objforce/objforce/cmd/index-api/app/domain/models"
	"github.com/objforce/objforce/cmd/index-api/app/domain/repositories"
)

type DocumentService interface {
	Upsert(c context.Context, m *models.Document) error
	FindOne(c context.Context, index, typ, id string) (*models.Document, error)
	Delete(c context.Context, index, typ, id string) error
}

type documentService struct {
	documentRepository repositories.DocumentRepository
}

func NewDocumentService(indexRepository repositories.DocumentRepository) DocumentService {
	return &documentService{
		indexRepository,
	}
}

func (s *documentService) Upsert(c context.Context, m *models.Document) error {
	return s.documentRepository.Upsert(c, m)
}

func (s *documentService) FindOne(c context.Context, index, typ, id string) (*models.Document, error) {
	document, err := s.documentRepository.FindOne(c, index, typ, id)

	return document, err
}

func (s *documentService) Delete(c context.Context, index, typ, id string) error {
	return s.documentRepository.Delete(c, index, typ, id)
}