package models

type DeploymentStatus string

const (
	DeploymentStatusInDevelopment DeploymentStatus = "InDevelopment"
	DeploymentStatusDeployed      DeploymentStatus = "Deployed"
)
