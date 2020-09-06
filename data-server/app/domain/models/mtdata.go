package models

type MTData struct {
	ObjId string `json:"objId,omitempty" gorm:"primary_key"`
	Id string `json:"Id,omitempty"`
	OrgId string `json:"orgId,omitempty"`
	Name string `json:"name"`
	Fields map[string]interface{}	`json:"fields"`
}