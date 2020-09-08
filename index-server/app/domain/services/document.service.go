package services

import (
	"context"
	"github.com/objforce/objforce/index-server/app/domain/models"
	"github.com/objforce/objforce/index-server/app/domain/repositories"
)

type IndexService interface {
	Create(c context.Context, m *models.Document) error
	Update(c context.Context, m *models.Document) error
	FindOne(c context.Context, index, typ, id string) (*models.Document, error)
	Delete(c context.Context, index, typ, id string) error
}

type indexService struct {
	indexRepository repositories.IndexRepository
}

func NewIndexService(indexRepository repositories.IndexRepository) IndexService {
	return &indexService{
		indexRepository,
	}
}

func (s *indexService) Create(c context.Context, m *models.Document) error {
	return s.indexRepository.Upsert(c, m)
}

func (s *indexService) Update(c context.Context, m *models.Document) error {
	return s.indexRepository.Upsert(c, m)
}

func (s *indexService) FindOne(c context.Context, index, typ, id string) (*models.Document, error) {
	document, err := s.indexRepository.FindOne(c, index, typ, id)

	return document, err
}

func (s *indexService) Delete(c context.Context, index, typ, id string) error {
	return s.indexRepository.Delete(c, index, typ, id)
}