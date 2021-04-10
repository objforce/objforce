package gorm

import (
	"context"

	"github.com/duolacloud/microbase/datasource/gorm/opentracing"
	"github.com/duolacloud/microbase/domain/repository"
	_gorm "github.com/jinzhu/gorm"
	"github.com/objforce/objforce/service/meta/domain/entities"
	"github.com/objforce/objforce/service/meta/domain/repositories"
	"go.uber.org/dig"
)

type RelationshipRepositoryTarget struct {
	dig.In
	DataSourceProvider repository.DataSourceProvider `name:"gorm"`
}

type relationshipRepository struct {
	dataSourceProvider repository.DataSourceProvider
}

func NewRelationshipRepository(target RelationshipRepositoryTarget) repositories.RelationshipRepository {
	return &relationshipRepository{
		dataSourceProvider: target.DataSourceProvider,
	}
}

func (r *relationshipRepository) DB(c context.Context) (*_gorm.DB, error) {
	db, err := r.dataSourceProvider.ProvideDB(c)
	if err != nil {
		return nil, err
	}
	return db.(*_gorm.DB), nil
}

func (r *relationshipRepository) Create(c context.Context, object *entities.MTRelationship) error {
	db, err := r.DB(c)
	if err != nil {
		return err
	}

	db = opentracing.SetSpanToGorm(c, db)

	if err := db.Create(object).Error; err != nil {
		return err
	}

	return nil
}

func (r *relationshipRepository) Upsert(c context.Context, object *entities.MTRelationship) error {
	db, err := r.DB(c)
	if err != nil {
		return err
	}
	db = opentracing.SetSpanToGorm(c, db)

	return db.Save(object).Error
}

func (r *relationshipRepository) Update(c context.Context, object *entities.MTRelationship) error {
	db, err := r.DB(c)
	if err != nil {
		return err
	}
	db = opentracing.SetSpanToGorm(c, db)

	// TODO 要记录变更历史 change
	rr := db.Save(object)

	return rr.Error
}

func (r *relationshipRepository) Get(c context.Context, objId string) (*entities.MTRelationship, error) {
	db, err := r.DB(c)
	db = opentracing.SetSpanToGorm(c, db)

	entity := &entities.MTRelationship{ObjId: objId}

	if err = db.Model(entity).Take(entity).Error; err != nil {
		return nil, err
	}

	return entity, err
}

func (r *relationshipRepository) Delete(c context.Context, objId string) error {
	db, err := r.DB(c)
	if err != nil {
		return err
	}
	return db.Delete(entities.MTRelationship{ObjId: objId}).Error
}
