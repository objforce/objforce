package repositories

import (
	"context"

	"github.com/objforce/objforce/service/meta/domain/entities"
)

type CustomObjectRepository interface {
	Create(c context.Context, object *entities.MTObject) error
	Upsert(c context.Context, object *entities.MTObject) error
	Update(c context.Context, object *entities.MTObject) error
	Get(c context.Context, objId string) (*entities.MTObject, error)
	Delete(c context.Context, objId string) error
	FindByObjName(c context.Context, orgId, objName string) (*entities.MTObject, error)
}
