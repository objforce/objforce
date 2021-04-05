package dtos

type CompactLayout struct {
	*Metadata
	Fields string `json:"fields"`
	Label string `json:"label"`
}
