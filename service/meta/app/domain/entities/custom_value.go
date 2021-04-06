package entities

/**
Represents the definition of a value used in a global value set or local custom picklist.
Custom picklist fields can be local and unique, or can inherit their values from a global picklist (called a global value set in API version 38.0).
This type extends the Metadata metadata type and inherits its fullName field.
To deactivate a global picklist value, you can invoke an update() call on GlobalPicklist (API version 37.0) or GlobalValueSet (API version 38.0 and later) with the value omitted, or with the value’s isActive field set to false.
Or, you can invoke an update() call directly on GlobalPicklistValue (API version 37.0) or CustomValue (API version 38.0 and later) with the isActive field set to false.
*/
type CustomValue struct {
	Color       string `json:"color"`
	Default     bool   `json:"default"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
	Label       string `json:"label"`
}
