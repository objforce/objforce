package models

type ListView struct {
	BooleanFilter string `json:"booleanFilter"`
	Columns []string `json:"columns"`
	Division string `json:"division"`
	FilterScope string `json:"filterScope"`
	Filters []*ListViewFilter `json:"filters"`
	Label string `json:"label"`
	Language string `json:"language"`
	Queue string `json:"queue"`
	sharedTo *SharedTo `json:"sharedTo"`
}
