package services

import (
	"context"

	"github.com/objforce/objforce/api/index/app/domain/models"
)

type DocumentService interface {
	Upsert(c context.Context, m *models.Document) error
	Retrieve(c context.Context, index, typ, id string) (*models.Document, error)
	Delete(c context.Context, index, typ, id string) error
}

type documentService struct {
}

func NewDocumentService() DocumentService {
	return &documentService{}
}

func (s *documentService) Upsert(c context.Context, m *models.Document) error {
	return nil
}

func (s *documentService) Retrieve(c context.Context, index, typ, id string) (*models.Document, error) {
	return nil, nil
}

func (s *documentService) Delete(c context.Context, index, typ, id string) error {
	return nil
}
