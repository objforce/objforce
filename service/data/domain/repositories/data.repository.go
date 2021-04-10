package repositories

import (
	"context"

	"github.com/objforce/objforce/service/data/domain/entities"
)

type DataRepository interface {
	Get(c context.Context, orgId, objId, id string, fields []string) (*entities.MTData, error)
	Create(c context.Context, m *entities.MTData) error
	Update(c context.Context, m *entities.MTData) error
	Upsert(c context.Context, m *entities.MTData) *entities.UpsertResult
	Delete(c context.Context, orgId, objId, id string) error

	MultiGet(c context.Context, orgId, objId string, ids []string, fields []string) []*entities.MTData
	MultiCreate(c context.Context, items []*entities.MTData) []*entities.SaveResult
	MultiUpdate(c context.Context, items []*entities.MTData) []*entities.SaveResult
	MultiUpsert(c context.Context, items []*entities.MTData) []*entities.UpsertResult
	MultiDelete(c context.Context, orgId, objId string, ids []string) []*entities.DeleteResult
}
