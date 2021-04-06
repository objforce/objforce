package models

type ValidationRule struct {
	*Metadata

	Active                bool   `json:"active,omitempty"`
	Description           string `json:"description,omitempty"`
	ErrorConditionFormula string `json:"errorConditionFormula,omitempty"`
	ErrorDisplayField     string `json:"errorDisplayField,omitempty"`
	ErrorMessage          string `json:"errorMessage,omitempty"`
}
