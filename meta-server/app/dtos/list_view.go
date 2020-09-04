package dtos

type ListView struct {
	*Metadata
	BooleanFilter string `json:"booleanFilter"`
	Columns []string `json:"columns"`
	Division string `json:"division"`
	FilterScope *FilterScope `json:"filterScope"`
	Filters []*ListViewFilter `json:"filters"`
	Label string `json:"label"`
	Language *Language `json:"language"`
	Queue string `json:"queue"`
	sharedTo *SharedTo `json:"sharedTo"`
}
