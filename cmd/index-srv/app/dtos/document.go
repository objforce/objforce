package dtos

type Document struct {
	ObjId string `json:"objId"`
	Fields map[string]interface{}
}
