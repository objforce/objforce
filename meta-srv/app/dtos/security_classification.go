package dtos

type SecurityClassification string

const(
	SecurityClassificationPublic SecurityClassification = "Public"
	SecurityClassificationInternal SecurityClassification = "Internal"
	SecurityClassificationConfidential SecurityClassification = "Confidential"
	SecurityClassificationRestricted SecurityClassification = "Restricted"
	SecurityClassificationMissionCritical SecurityClassification = "MissionCritical"
)