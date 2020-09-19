package dtos

type FilterOperation string

const(
	FilterOperationEQ FilterOperation = "EQ"
	FilterOperationNE FilterOperation = "NE"
	FilterOperationLT FilterOperation = "LT"
	FilterOperationGT FilterOperation = "GT"
	FilterOperationLE FilterOperation = "LE"
	FilterOperationGE FilterOperation = "GE"
	FilterOperationContains FilterOperation = "contains"
	FilterOperationNotContain FilterOperation = "notContain"
	FilterOperationStartsWith FilterOperation = "startsWith"
	FilterOperationIncludes FilterOperation = "includes"
	FilterOperationExcludes FilterOperation = "excludes"
	FilterOperationWithin FilterOperation = "within"
)