package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/objforce/objforce/meta-srv/app/domain/models"
)

const (
	BASIC_TABLE_FIELDS = 30
)

type FieldLocation struct {
	FieldNum int		`json:"fieldNum"`
	TableName string `json:"tableName"`
}

type FieldAllocator interface {
	Allocate(db *gorm.DB, objId string) (*FieldLocation, error)
}

type fieldAllocator struct {
	extendTableMetaRepository ExtendTableMetaRepository
}

func NewFieldAllocator(extendTableMetaRepository ExtendTableMetaRepository) FieldAllocator {
	return &fieldAllocator{
		extendTableMetaRepository,
	}
}

func (a *fieldAllocator) Allocate(db *gorm.DB, objId string) (*FieldLocation, error) {
	var fields []*models.MTField
	err := db.Model(&models.MTField{}).Where("obj_id = ?", objId).Find(&fields).Error
	if err != nil {
		return nil, err
	}

	// 优先从 basic table 中找到可用的字段

	// 从当前在使用的扩展表中找到可用的字段，如果当前可用扩展表无字段可用，要驱动扩展表迁移

	fieldLocation := &FieldLocation{}


	// basic table fields
	if len(fields) < BASIC_TABLE_FIELDS {
		fieldLocation.TableName = "mt_basic_objects"
		fieldLocation.FieldNum = len(fields)
		return fieldLocation, nil
	}

	_, err = a.extendTableMetaRepository.All()
	if err != nil {
		return nil, err
	}



	return nil, err
}
