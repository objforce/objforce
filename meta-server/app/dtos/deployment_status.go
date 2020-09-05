package dtos

type DeploymentStatus string

const(
	DeploymentStatusInDevelopment DeploymentStatus = "InDevelopment"
	DeploymentStatusDeployed DeploymentStatus = "Deployed"
)