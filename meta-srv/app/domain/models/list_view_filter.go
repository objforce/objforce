package models

type ListViewFilter struct {
	Field     string           `json:"field,omitempty"`
	Operation FilterOperation `json:"operation,omitempty"`
	Value     string           `json:"value,omitempty"`
}