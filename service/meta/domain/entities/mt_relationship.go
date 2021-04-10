package entities

type MTRelationship struct {
	GUID        string `json:"guid"`
	OrgId       string `json:"orgId"`
	ObjId       string `json:"objId"`
	RelationId  string `json:"relationId"`
	TargetObjId string `json:"targetObjId"`
}
