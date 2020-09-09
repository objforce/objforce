package models

type Document struct {
	Index string	`json:"index"`
	Type string		`json:"type"`
	Id string `json:"id"`
	Fields map[string]interface{}
}
