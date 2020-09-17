package dtos

type SObject struct {
	ObjId string `json:"objId,omitempty"`
	Id string `json:"Id,omitempty"`
	OrgId string `json:"orgId,omitempty"`
	Name string `json:"name"`
	Fields map[string]interface{}	`json:"fields"`
}

type RetrieveSObjectRequest struct {
	Fields string `json:"fields"`
	Type string `json:"type,omitempty"`
	Ids []string `xml:"ids,omitempty"`
}

type RetrieveSObjectResponse struct {
	Results []*SObject `json:"results"`
}

type DeleteSObjectResponse struct {
	Results []*DeleteResult `json:"results"`
}

type CreateSObjectResponse struct {
	Results []*SaveResult `json:"results"`
}

type UpdateSObjectRequest struct {
	Objects []*SObject `json:"objects,omitempty"`
}

type UpdateSObjectResponse struct {
	Results []*SaveResult `json:"results"`
}

type UpsertSObjectRequest struct {
	ExternalIDFieldName string `json:"externalIDFieldName,omitempty"`

	Objects []*SObject `json:"objects,omitempty"`
}

type UpsertSObjectResponse struct {
	Results []*UpsertResult `xml:"results,omitempty"`
}

type UpsertResult struct {
	Created bool `json:"created,omitempty"`
	Errors []*Error `json:"errors,omitempty"`
	Id string `json:"id,omitempty"`
	Success bool `json:"success,omitempty"`
}

type Error struct {
	Message string `json:"message"`
	StatusCode StatusCode `json:"code"`
}

type SaveResult struct {
	Errors []*Error `xml:"errors,omitempty"`

	Id string `xml:"id,omitempty"`

	Success bool `xml:"success,omitempty"`
}

type DeleteResult struct {
	Errors []*Error `json:"errors,omitempty"`

	Id string `json:"id,omitempty"`
	Success bool `json:"success,omitempty"`
}

type StatusCode string

const (
	StatusCodeALL_OR_NONE_OPERATION_ROLLED_BACK StatusCode = "ALL_OR_NONE_OPERATION_ROLLED_BACK"

	StatusCodeALREADY_IN_PROCESS StatusCode = "ALREADY_IN_PROCESS"
)