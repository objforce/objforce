package repositories

import (
	"context"

	"github.com/objforce/objforce/service/data/app/domain/entities"
)

type DataRepository interface {
	Get(c context.Context, orgId string, objId string, guid string, fields []string) (*entities.MTData, error)
	Create(c context.Context, m *entities.MTData) error
	Update(c context.Context, m *entities.MTData) error
	Upsert(c context.Context, m *entities.MTData) *entities.UpsertResult
	Delete(c context.Context, orgId string, objId string, guid string) error

	MultiGet(c context.Context, orgId string, objId string, ids []string, fields []string) []*entities.MTData
	MultiCreate(c context.Context, items []*entities.MTData) []*entities.SaveResult
	MultiUpdate(c context.Context, items []*entities.MTData) []*entities.SaveResult
	MultiUpsert(c context.Context, items []*entities.MTData) []*entities.UpsertResult
	MultiDelete(c context.Context, orgId string, objId string, ids []string) []*entities.DeleteResult
}
