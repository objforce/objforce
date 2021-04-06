package entities

type FilterItem struct {
	Field      string          `json:"field"`      // Represents the field specified in the filter
	Operation  FilterOperation `json:"operation"`  // Represents the filter operation for this filter item. Valid values are enumerated in FilterOperation
	Value      string          `json:"value"`      // Represents the value of the filter item being operated upon, for example, if the filter is my_number_field__c > 1, the value of value is 1
	ValueField string          `json:"valueField"` // Specifies if the final column in the filter contains a field or a field value. Approval processes don’t support valueField entries in filter criteria
}
