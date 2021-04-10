package repositories

import (
	"context"

	"github.com/objforce/objforce/service/meta/domain/entities"
)

type RelationshipRepository interface {
	Create(c context.Context, object *entities.MTRelationship) error
	Upsert(c context.Context, object *entities.MTRelationship) error
	Update(c context.Context, object *entities.MTRelationship) error
	Get(c context.Context, objId string) (*entities.MTRelationship, error)
	Delete(c context.Context, objId string) error
}
