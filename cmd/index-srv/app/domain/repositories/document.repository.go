package repositories

import (
	"context"
	"encoding/json"
	"github.com/objforce/objforce/cmd/index-srv/app/domain/models"
	"github.com/olivere/elastic/v6"
)

type documentRepository struct {
	client *elastic.Client
}

type DocumentRepository interface {
	Retrieve(c context.Context, index, typ, id string) (*models.Document, error)
	Delete(c context.Context, index, typ, id string) error
	Update(c context.Context, doc *models.Document) error
	Upsert(c context.Context, doc *models.Document) error
	Bulk(c context.Context, docs []*models.Document) error
}

func NewDocumentRepository(client *elastic.Client) DocumentRepository {
	return &documentRepository{
		client,
	}
}

func (r *documentRepository) Retrieve(c context.Context, index, typ, id string) (*models.Document, error) {
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

func (r *documentRepository) Delete(c context.Context, index, typ, id string) error {
	_, err := r.client.Delete().Index(index).Type(typ).Id(id).Do(c)
	return err
}

func (r *documentRepository) Bulk(c context.Context, docs []*models.Document) error {
	bulkReq := r.client.Bulk()
	for _, doc := range docs {
		req := elastic.NewBulkIndexRequest().Index(doc.Index).Type(doc.Type).Id(doc.Id).Doc(doc.Fields)
		bulkReq = bulkReq.Add(req)
	}

	_/*bulkResponse*/, err := bulkReq.Do(c)

	return err
}

func (r *documentRepository) Upsert(c context.Context, doc *models.Document) error {
	_, err := r.client.Update().
		Index(doc.Index).
		Type(doc.Type).
		Id(doc.Id).
		Upsert(doc).
		Do(c)
	return err
}

func (r *documentRepository) Update(c context.Context, doc *models.Document) error {
	_, err := r.client.Update().
		Index(doc.Index).
		Type(doc.Type).
		Id(doc.Id).
		Doc(doc).
		Do(c)

	return err
}