package dtos

type SharingReason struct {
	*Metadata

	Label string `json:"label,omitempty"`
}