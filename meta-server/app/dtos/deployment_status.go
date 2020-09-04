package dtos

type DeploymentStatus string

const(
	DeploymentStatus_InDevelopment DeploymentStatus = "InDevelopment"
	DeploymentStatus_Deployed DeploymentStatus = "Deployed"
)