package dtos

type SObject struct {
	ObjId string `json:"objId,omitempty"`
	Id string `json:"Id,omitempty"`
	OrgId string `json:"orgId,omitempty"`
	Name string `json:"name"`
	Fields map[string]interface{}	`json:"fields"`
}