package models

type Index struct {
	Fields []*IndexField `json:"fields"`
	Label  string        `json:"label"`
}

type IndexField struct {
	Name          string `json:"name"`
	SortDirection string `json:"sortDirection"`
}
