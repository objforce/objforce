package models

type MTExtendData struct {
	GUID      string                 `json:"guid,omitempty" gorm:"primary_key"`
	OrgId     string                 `json:"orgId,omitempty"`
	ObjId     string                 `json:"objId,omitempty"`
	Name      string                 `json:"name"`
	Fields    map[string]interface{} `json:"fields"`
}

func (m *MTExtendData) Unique() interface{} {
	return map[string]interface{}{
		"guid": m.GUID,
	}
}