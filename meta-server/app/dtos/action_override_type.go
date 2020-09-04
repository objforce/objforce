package dtos

type ActionOverrideType string

const(
	DEFAULT ActionOverrideType = "default"
	FLEXIPAGE ActionOverrideType = "flexipage"
	LIGHTNINGCOMPONENT ActionOverrideType = "lightningcomponent"
	SCONTROL ActionOverrideType = "scontrol"
	STANDARD ActionOverrideType = "standard"
	VISUALFORCE ActionOverrideType = "visualforce"
)