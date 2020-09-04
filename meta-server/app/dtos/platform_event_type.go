package dtos

type PlatformEventType string

const(
	PlatformEventType_HighVolume PlatformEventType = "HighVolume"
	PlatformEventType_StandardVolume PlatformEventType = "StandardVolume"
)