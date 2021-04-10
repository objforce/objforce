package entities

import "time"

type MTData struct {
	GUID  string `json:"guid,omitempty" gorm:"primary_key"`
	OrgId string `json:"orgId,omitempty"`
	ObjId string `json:"objId,omitempty"`
	// Name      string            `json:"name"`
	Fields    map[string][]byte `json:"fields"`
	CreatedAt time.Time         `json:"createdAt"`
	CreatedBy *string           `json:"createdBy"`
	UpdatedAt time.Time         `json:"updatedAt"`
	UpdatedBy *string           `json:"updatedBy"`
	IsDeleted bool              `json:"isDeleted"`
}

type SaveResult struct {
	Error   error  `json:"error,omitempty"`
	Id      string `json:"id,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type DeleteResult struct {
	Error   error  `json:"error"`
	Id      string `json:"id"`
	Success bool   `json:"success"`
}

type UpsertResult struct {
	Created bool    `json:"created,omitempty"`
	Errors  []error `json:"error,omitempty"`
	Id      string  `json:"id,omitempty"`
	Success bool    `json:"success,omitempty"`
}
