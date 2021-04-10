package models

type DeleteConstraint int32

const (
	DeleteConstraint_SetNull  DeleteConstraint = 0
	DeleteConstraint_Restrict DeleteConstraint = 1
	DeleteConstraint_Cascade  DeleteConstraint = 2
)
