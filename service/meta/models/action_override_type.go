package models

type ActionOverrideType string

const (
	ActionOverrideTypeDefault            ActionOverrideType = "default"
	ActionOverrideTypeFlexipage          ActionOverrideType = "flexipage"
	ActionOverrideTypeLightningcomponent ActionOverrideType = "lightningcomponent"
	ActionOverrideTypeScontrol           ActionOverrideType = "scontrol"
	ActionOverrideTypeStandard           ActionOverrideType = "standard"
	ActionOverrideTypeVisualforce        ActionOverrideType = "visualforce"
)
