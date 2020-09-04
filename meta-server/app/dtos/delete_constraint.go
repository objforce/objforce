package dtos

type DeleteConstraint string

const (
	SetNull DeleteConstraint = "SetNull"
	Restrict DeleteConstraint = "Restrict"
	Cascade DeleteConstraint = "Cascade"
)