package models

import "time"

type MTData struct {
	Id string `json:"Id,omitempty" gorm:"primary_key"`
	ObjId string `json:"objId,omitempty"`
	OrgId string `json:"orgId,omitempty"`
	Name string `json:"name"`
	Fields map[string]interface{}	`json:"fields"`
	CreatedAt time.Time `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
	UpdatedAt time.Time	`json:"updatedAt"`
	UpdatedBy string `json:"updatedBy"`
}

func (m *MTData) Unique() interface{} {
	return map[string]interface{}{
		"id": m.Id,
	}
}