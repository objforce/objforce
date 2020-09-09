package models

/**
Represents the metadata associated with a lookup filter. This metadata type is used to create, update, or delete lookup filter definitions. This component has been removed as of API version 30.0 and is only available in previous API versions. The metadata associated with a lookup filter is now represented by the lookupFilter field in the CustomField component.
This type extends the Metadata metadata type and inherits its fullName field. You can also use this metadata type to work with customizations of lookup filters on standard fields.
 */

type NamedFilter struct {
	*Metadata
	Active bool `json:"active"`
	BooleanFilter string `json:"booleanFilter"`
	Description string `json:"description"`
	ErrorMessage string `json:"errorMessage"`
	Field string `json:"field"`
	FilterItems []*FilterItem `json:"filterItems"`
	InfoMessage string `json:"infoMessage"`
	IsOptional bool `json:"isOptional"`
	Name string `json:"name"`
	SourceObject string `json:"sourceObject"`
}
