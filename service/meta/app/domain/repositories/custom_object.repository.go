package repositories

import (
	"context"

	_gorm "github.com/jinzhu/gorm"
	"github.com/objforce/objforce/service/meta/app/domain/models"
	"github.com/xxxmicro/base/database/gorm/opentracing"
)

type CustomObjectRepository interface {
	Create(c context.Context, object *models.MTObject) error
	Upsert(c context.Context, object *models.MTObject) error
	Update(c context.Context, object *models.MTObject) error
	Retrieve(c context.Context, objId string) (*models.MTObject, error)
	Delete(c context.Context, objId string) error
	FindCustomObjectByOrgAndType(c context.Context, orgId, objType string) (*models.MTObject, error)
}

type customObjectRepository struct {
	db *_gorm.DB
}

func NewCustomObjectRepository(db *_gorm.DB) CustomObjectRepository {
	return &customObjectRepository{
		db,
	}
}

func (r *customObjectRepository) Create(c context.Context, object *models.MTObject) error {
	db := opentracing.SetSpanToGorm(c, r.db)
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

func (r *customObjectRepository) Upsert(c context.Context, object *models.MTObject) error {
	db := opentracing.SetSpanToGorm(c, r.db)

	return db.Save(object).Error
}

func (r *customObjectRepository) Update(c context.Context, object *models.MTObject) error {
	db := opentracing.SetSpanToGorm(c, r.db)

	// TODO 要记录变更历史 change
	rr := db.Save(object)

	return rr.Error
}

func (r *customObjectRepository) Retrieve(c context.Context, objId string) (*models.MTObject, error) {
	db := opentracing.SetSpanToGorm(c, r.db)

	object := &models.MTObject{ObjId: objId}
	err := db.Model(object).Take(object).Error
	if err != nil {
		return nil, err
	}

	if err := db.Model(object).Association("Fields").Find(&object.Fields).Error; err != nil {
		return nil, err
	}

	return object, err
}

func (r *customObjectRepository) Delete(c context.Context, objId string) error {
	return r.db.Delete(models.MTObject{ObjId: objId}).Error
}

func (r *customObjectRepository) FindCustomObjectByOrgAndType(c context.Context, orgId, objType string) (*models.MTObject, error) {
	db := opentracing.SetSpanToGorm(c, r.db)

	m := &models.MTObject{}
	err := db.Where("org_id = ? AND obj_name = ?", orgId, objType).Take(m).Error
	if err != nil {
		return nil, err
	}
	if err := db.Model(m).Association("Fields").Find(&m.Fields).Error; err != nil {
		return nil, err
	}

	return m, nil
}
