package models

import "time"

type MTBasicData struct {
	GUID      string                 `json:"guid,omitempty" gorm:"primary_key"`
	OrgId     string                 `json:"orgId,omitempty"`
	ObjId     string                 `json:"objId,omitempty"`
	Name      string                 `json:"name"`
	Fields    map[string]interface{} `json:"fields"`
	CreatedAt time.Time              `json:"createdAt"`
	CreatedBy string                 `json:"createdBy"`
	UpdatedAt time.Time              `json:"updatedAt"`
	UpdatedBy string                 `json:"updatedBy"`
	IsDeleted bool					 `json:"isDeleted"`
}

func (m *MTBasicData) Unique() interface{} {
	return map[string]interface{}{
		"guid": m.GUID,
	}
}