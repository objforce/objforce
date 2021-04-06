package gorm

import (
	"context"

	"github.com/duolacloud/microbase/datasource/gorm/opentracing"
	"github.com/duolacloud/microbase/domain/repository"
	_gorm "github.com/jinzhu/gorm"
	"github.com/objforce/objforce/service/meta/app/domain/entities"
	"github.com/objforce/objforce/service/meta/app/domain/repositories"
	"go.uber.org/dig"
)

type CustomerObjectRepositoryTarget struct {
	dig.In
	DataSourceProvider repository.DataSourceProvider `name:"gorm"`
}

type customObjectRepository struct {
	dataSourceProvider repository.DataSourceProvider
}

func NewCustomObjectRepository(target CustomerObjectRepositoryTarget) repositories.CustomObjectRepository {
	return &customObjectRepository{
		dataSourceProvider: target.DataSourceProvider,
	}
}

func (r *customObjectRepository) DB(c context.Context) (*_gorm.DB, error) {
	db, err := r.dataSourceProvider.ProvideDB(c)
	if err != nil {
		return nil, err
	}
	return db.(*_gorm.DB), nil
}

func (r *customObjectRepository) Create(c context.Context, object *entities.MTObject) error {
	db, err := r.DB(c)
	if err != nil {
		return err
	}

	db = opentracing.SetSpanToGorm(c, db)
	/*return db.Transaction(func(tx *_gorm.DB) error {

	})*/

	// Create 方法内部默认就用了事务

	fieldNum := 0
	for _, customField := range object.Fields {
		customField.FieldNum = fieldNum
		fieldNum++
	}

	if err := db.Create(object).Error; err != nil {
		return err
	}

	return nil
}

func (r *customObjectRepository) Upsert(c context.Context, object *entities.MTObject) error {
	db, err := r.DB(c)
	if err != nil {
		return err
	}
	db = opentracing.SetSpanToGorm(c, db)

	return db.Save(object).Error
}

func (r *customObjectRepository) Update(c context.Context, object *entities.MTObject) error {
	db, err := r.DB(c)
	if err != nil {
		return err
	}
	db = opentracing.SetSpanToGorm(c, db)

	// TODO 要记录变更历史 change
	rr := db.Save(object)

	return rr.Error
}

func (r *customObjectRepository) Get(c context.Context, objId string) (*entities.MTObject, error) {
	db, err := r.DB(c)
	db = opentracing.SetSpanToGorm(c, db)

	entity := &entities.MTObject{ObjId: objId}

	if err = db.Model(entity).Take(entity).Error; err != nil {
		return nil, err
	}

	if err := db.Model(entity).Association("Fields").Find(&entity.Fields).Error; err != nil {
		return nil, err
	}

	return entity, err
}

func (r *customObjectRepository) Delete(c context.Context, objId string) error {
	db, err := r.DB(c)
	if err != nil {
		return err
	}
	return db.Delete(entities.MTObject{ObjId: objId}).Error
}

func (r *customObjectRepository) FindCustomObjectByOrgAndType(c context.Context, orgId, objType string) (*entities.MTObject, error) {
	db, err := r.DB(c)
	db = opentracing.SetSpanToGorm(c, db)

	entity := &entities.MTObject{}
	if err = db.Where("org_id = ? AND obj_name = ?", orgId, objType).Take(entity).Error; err != nil {
		return nil, err
	}
	if err := db.Model(entity).Association("Fields").Find(&entity.Fields).Error; err != nil {
		return nil, err
	}

	return entity, nil
}
