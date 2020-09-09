package models

type Index struct {
	OrgId string `json:"orgName"`
	Name  string `json:"name"`
	Mapping map[string]interface{} `json:"mapping"`
}