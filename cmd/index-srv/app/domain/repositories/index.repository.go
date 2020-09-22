package repositories

import (
	"context"
	"github.com/objforce/objforce/cmd/index-srv/app/domain/models"
	"github.com/olivere/elastic/v6"
)

type indexRepository struct {
	client *elastic.Client
}

type IndexRepository interface {
	FindOne(c context.Context, index, typ string) (*models.Index, error)
	Delete(c context.Context, index string) error
	Create(c context.Context, index *models.Index) error
}

func NewIndexRepository(client *elastic.Client) IndexRepository {
	return &indexRepository{
		client,
	}
}

func (r *indexRepository) FindOne(c context.Context, index, typ string) (*models.Index, error) {
	mapping, err := r.client.GetMapping().Index(index).Type(typ).Do(c)
	if err != nil {
		return nil, err
	}

	m := &models.Index{
		Name: index,
		Mapping: mapping,
	}
	return m, nil
}

func (r *indexRepository) Delete(c context.Context, index string) error {
	_, err := r.client.DeleteIndex(index).Do(c)
	return err
}

func (r *indexRepository) Create(c context.Context, index *models.Index) error {
	_, err := r.client.CreateIndex(index.Name).BodyJson(index).Do(c)
	return err
}