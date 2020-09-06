package models

import "time"

type MTData struct {
	Id string `json:"Id,omitempty" gorm:"primary_key"`
	ObjId string `json:"objId,omitempty"`
	OrgId string `json:"orgId,omitempty"`
	Name string `json:"name"`
	Fields map[string]interface{}	`json:"fields"`
	Created time.Time `json:"created"`
	CreateBy string `json:"createuser"`
	LastModified time.Time	`json:"lastModified"`
	LastModifiedBy string `json:"lastModifiedBy"`
}