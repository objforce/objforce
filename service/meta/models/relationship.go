package models

type Relationship struct {
	GUID        string `json:"guid,omitempty"`
	OrgId       string `json:"orgId,omitempty"`
	ObjId       string `json:"objId,omitempty"`
	RelationId  string `json:"relationId,omitempty"`
	TargetObjId string `json:"targetObjId,omitempty"`
}
