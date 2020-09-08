package repositories

import (
	"context"
	"encoding/json"
	"github.com/objforce/objforce/index-server/app/domain/models"
	"github.com/olivere/elastic/v6"
)

type indexRepository struct {
	client *elastic.Client
}

type IndexRepository interface {
	FindOne(c context.Context, index, typ, id string) (*models.Document, error)
	Delete(c context.Context, index, typ, id string) error
	Upsert(c context.Context, doc *models.Document) error
	Bulk(c context.Context, docs []*models.Document) error
}

func NewIndexRepository(client *elastic.Client) IndexRepository {
	return &indexRepository{
		client,
	}
}

func (r *indexRepository) FindOne(c context.Context, index, typ, id string) (*models.Document, error) {
	res, err := r.client.Get().Index(index).Type(typ).Id(id).Do(c)
	if err != nil {
		return nil, err
	}

	if !res.Found {
		return nil/* ErrNotFound */, nil
	}

	m := &models.Document{
		Index: index,
		Type: typ,
		Id: id,
	}

	err = json.Unmarshal(*res.Source, &m.Fields)
	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *indexRepository) Delete(c context.Context, index, typ, id string) error {
	_, err := r.client.Delete().Index(index).Type(typ).Id(id).Do(c)
	return err
}

func (r *indexRepository) Bulk(c context.Context, docs []*models.Document) error {
	bulkReq := r.client.Bulk()
	for _, doc := range docs {
		req := elastic.NewBulkIndexRequest().Index(doc.Index).Type(doc.Type).Id(doc.Id).Doc(doc.Fields)
		bulkReq = bulkReq.Add(req)
	}

	_/*bulkResponse*/, err := bulkReq.Do(c)

	return err
}

func (r *indexRepository) Upsert(c context.Context, doc *models.Document) error {
	_, err := r.client.Index().
		Index(doc.Index).
		Type(doc.Type).
		Id(doc.Id).
		BodyJson(doc.Fields).
		Do(c)
	return err
}