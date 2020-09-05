package dtos

type PlatformEventType string

const(
	PlatformEventTypeHighVolume PlatformEventType = "HighVolume"
	PlatformEventTypeStandardVolume PlatformEventType = "StandardVolume"
)