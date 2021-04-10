package models

type FieldSet struct {
	AvailableFields []FieldSetItem `json:"availableFields"`
	Description     string         `json:"description"`
	DisplayedFields []FieldSetItem `json:"displayedFields"`
	Label           string         `json:"label"`
}

type FieldSetItem struct {
	Field          string `json:"field"`
	IsFieldManaged bool   `json:"isFieldManaged"`
	IsRequired     bool   `json:"isRequired"`
}
