package dtos

type DeleteConstraint string

const (
	DeleteConstraintSetNull DeleteConstraint = "SetNull"
	DeleteConstraintRestrict DeleteConstraint = "Restrict"
	DeleteConstraintCascade DeleteConstraint = "Cascade"
)