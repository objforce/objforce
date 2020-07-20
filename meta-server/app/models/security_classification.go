package models

type SecurityClassification string

const(
	Public SecurityClassification = "Public"
	Internal SecurityClassification = "Internal"
	Confidential SecurityClassification = "Confidential"
	Restricted SecurityClassification = "Restricted"
	MissionCritical SecurityClassification = "MissionCritical"
)