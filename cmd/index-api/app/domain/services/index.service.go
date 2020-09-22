package services

import (
	"context"
	"github.com/objforce/objforce/cmd/index-api/app/domain/models"
)

type IndexService interface {
	Create(c context.Context, m *models.Index) error
	Retrieve(c context.Context, objId string) (*models.Index, error)
	Delete(c context.Context, objId string) error
}

type indexService struct {
}

func NewIndexService() IndexService {
	return &indexService{
	}
}

func (s *indexService) Create(c context.Context, m *models.Index) error {
	return nil
}

func (s *indexService) Retrieve(c context.Context, objId string) (*models.Index, error) {
	return nil, nil
}

func (s *indexService) Delete(c context.Context, objId string) error {
	return nil
}