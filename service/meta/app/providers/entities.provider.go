package providers

import (
	"github.com/duolacloud/microbase/datasource"
	"github.com/objforce/objforce/service/meta/app/domain/entities"
)

type OLTPEntityMap struct {
}

func (m *OLTPEntityMap) GetEntities() []interface{} {
	return []interface{}{
		&entities.MTObject{},
		&entities.MTField{},
	}
}

func NewOLTPEntityMap() datasource.EntityMap {
	return &OLTPEntityMap{}
}
