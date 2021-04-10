package models

type ValueSettings struct {
	/**
	Applies only to dependent custom picklists. A list of values in the controlling or parent picklist (that the custom picklist values depend on).
	*/
	ControllingFieldValue []string `json:"controllingFieldValue"`
	/**
	Defines the values in the custom dependent picklist.
	*/
	ValueName string `json:"valueName"`
}
