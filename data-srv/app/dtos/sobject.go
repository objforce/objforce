package dtos

type SObject struct {
	Id string `json:"GUID,omitempty"`
	Type string `json:"type,omitempty"`
	FieldsToNull []string `json:"fieldsToNull,omitempty"`
	Fields map[string]string	`json:"fields"`
}

type SaveResult struct {
	Error error `json:"error,omitempty"`
	Id string `json:"id,omitempty"`
	Success bool `json:"sucfjjcess,omitempty"`
}

type UpsertResult struct {
	Created bool `json:"created,omitempty"`
	Error error `json:"error,omitempty"`
	Id string `json:"id,omitempty"`
	Success bool `json:"success,omitempty"`
}

type DeleteResult struct {
	Error error `json:"error"`
	Id string `json:"id"`
	Success bool `json:"success"`
}