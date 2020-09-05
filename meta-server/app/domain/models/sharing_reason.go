package models

type SharingReason struct {
	*Metadata

	Label string `json:"label,omitempty"`
}