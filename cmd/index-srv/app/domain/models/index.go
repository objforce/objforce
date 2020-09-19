package models

import(
	"time"
)

type Index struct {
	OrgId string `json:"orgId"`
	Name  string `json:"name"`
	Mapping map[string]interface{} `json:"mapping"`
}

type MTIndex struct {
	GUID string `json:"guid"`
	OrgId string `json:"orgId"`
	ObjId string `json:"objId"`
	FieldNum int `json:"fieldNum"`
	StringValue string `json:"stringValue"`
	NumValue float64 `json:"numValue"`
	DateValue time.Time `json:"dateValue"`
}

type MTUniqueIndex struct {
	GUID string `json:"guid"`
	OrgId string `json:"orgId"`
	ObjId string `json:"objId"`
	FieldNum int `json:"fieldNum"`
	StringValue string `json:"stringValue"`
	NumValue float64 `json:"numValue"`
	DateValue time.Time `json:"dateValue"`
}

// MT_Fallback_Indexes
type MTFallbackIndex struct {

}