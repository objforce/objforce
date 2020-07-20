package models

type SummaryOperation string

const(
	Count SummaryOperation = "Count"
	Min SummaryOperation = "Min"
	Max SummaryOperation = "Max"
	Sum SummaryOperation = "Sum"
)