package models

type SObject struct {
	Id           string            `json:"GUID,omitempty"`
	Type         string            `json:"type,omitempty"`
	FieldsToNull []string          `json:"fieldsToNull,omitempty"`
	Fields       map[string]string `json:"fields"`
}

type GetSObjectRequest struct {
	Fields string `json:"fields"`
	Type   string `json:"type,omitempty"`
	Id     string `xml:"ids,omitempty"`
}

type DeleteSObjectRequest struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

type UpsertResult struct {
	Created bool     `json:"created,omitempty"`
	Errors  []*Error `json:"errors,omitempty"`
	Id      string   `json:"id,omitempty"`
	Success bool     `json:"success,omitempty"`
}

type Error struct {
	Message    string     `json:"message"`
	StatusCode StatusCode `json:"code"`
}

type SaveResult struct {
	Errors []*Error `xml:"errors,omitempty"`

	Id string `xml:"id,omitempty"`

	Success bool `xml:"success,omitempty"`
}

type DeleteResult struct {
	Errors []*Error `json:"errors,omitempty"`

	Id      string `json:"id,omitempty"`
	Success bool   `json:"success,omitempty"`
}

type StatusCode string

const (
	StatusCodeALL_OR_NONE_OPERATION_ROLLED_BACK StatusCode = "ALL_OR_NONE_OPERATION_ROLLED_BACK"

	StatusCodeALREADY_IN_PROCESS StatusCode = "ALREADY_IN_PROCESS"
)
