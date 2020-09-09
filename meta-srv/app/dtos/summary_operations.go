package dtos

type SummaryOperation string

const(
	SummaryOperationCount SummaryOperation = "Count"
	SummaryOperationMin SummaryOperation = "Min"
	SummaryOperationMax SummaryOperation = "Max"
	SummaryOperationSum SummaryOperation = "Sum"
)