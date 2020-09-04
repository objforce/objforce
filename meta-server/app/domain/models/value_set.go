package models

type ValueSet struct {
	/**
		The fullname of the controlling field if this is a dependent picklist. 
		A controlling field can be a checkbox or picklist field, but in this case it’s a picklist. 
		The controlling picklist filters the available values in the dependent picklist.
	 */
	ControllingField string `json:"controllingField"`

	/**
		Whether the picklist’s values are limited to only the values defined by a Salesforce admin. Values are true or false.
	 */
	Restricted bool `json:"restricted"`

	/**
		Defines value-specific settings for a custom dependent picklist. Indicates whether the value set of the custom picklist field is sorted alphabetically.
	 */
	ValueSetDefinition ValueSetValuesDefinition `json:"valueSetDefinition"`
	
	/**
		The masterLabel of the global value set to be used for this picklist field.
	 */
	ValueSetName string `json:"valueSetName"`

	/**
		Used for the settings that describe a value in a custom picklist field. 
		The picklist can have its own unique value set, or inherit the values from a global value set. 
		You can add field dependency values via Metadata API but not remove them.
	 */
	ValueSettings ValueSettings `json:"valueSettings"`	
}