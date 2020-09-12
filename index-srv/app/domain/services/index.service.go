package services

import (
	"context"
	"github.com/objforce/objforce/index-srv/app/domain/models"
	"github.com/objforce/objforce/index-srv/app/domain/repositories"
)

type IndexService interface {
	Create(c context.Context, m *models.Index) error
	FindOne(c context.Context, objId string) (*models.Index, error)
	Delete(c context.Context, objId string) error
}

type indexService struct {
	indexRepository repositories.IndexRepository
}

func NewIndexService(indexRepository repositories.IndexRepository) IndexService {
	return &indexService{
		indexRepository,
	}
}

func (s *indexService) Create(c context.Context, m *models.Index) error {
	return s.indexRepository.Create(c, m)
}

func (s *indexService) FindOne(c context.Context, objId string) (*models.Index, error) {
	index, err := s.indexRepository.FindOne(c, objId, "doc")
	return index, err
}

func (s *indexService) Delete(c context.Context, objId string) error {

	return s.indexRepository.Delete(c, objId)
}