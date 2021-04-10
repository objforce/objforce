package models

type PlatformEventPublishBehavior string

const (
	PlatformEventPublishBehaviorPublishAfterCommit PlatformEventPublishBehavior = "PublishAfterCommit"
	PlatformEventPublishBehaviorPublishImmediately PlatformEventPublishBehavior = "PublishImmediately"
)
