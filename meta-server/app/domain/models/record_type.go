package models

type RecordType struct {
	*Metadata

	Active                  bool                       `json:"active,omitempty"`
	BusinessProcess         string                     `json:"businessProcess,omitempty"`
	CompactLayoutAssignment string                     `json:"compactLayoutAssignment,omitempty"`
	Description             string                     `json:"description,omitempty"`
	Label                   string                     `json:"label,omitempty"`
	PicklistValues          []*RecordTypePicklistValue `json:"picklistValues,omitempty"`
}

type RecordTypePicklistValue struct {
	Picklist string           `json:"picklist,omitempty"`
	Values   []*PicklistValue `json:"values,omitempty"`
}