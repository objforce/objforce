package models

type ValueSetValuesDefinition struct {
	/**
	Whether the picklist’s value set is displayed in alphabetical order in the user interface
	*/
	Sorted bool `json:"sorted"`

	/**
	Required. The list of values for this local, custom picklist.
	*/
	Value CustomValue `json:"value"`
}
