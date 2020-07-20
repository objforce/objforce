package models

type FilterOperation string

const(
	EQ FilterOperation = "EQ"
	NE FilterOperation = "NE"
	LT FilterOperation = "LT"
	GT FilterOperation = "GT"
	LE FilterOperation = "LE"
	GE FilterOperation = "GE"
	Contains FilterOperation = "contains"
	NotContain FilterOperation = "notContain"
	StartsWith FilterOperation = "startsWith"
	Includes FilterOperation = "includes"
	Excludes FilterOperation = "excludes"
	Within FilterOperation = "within"
)