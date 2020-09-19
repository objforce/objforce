package dtos

type SetupObjectVisibility string

const(
	SetupObjectVisibilityPublic SetupObjectVisibility = "Public"
	SetupObjectVisibilityProtected SetupObjectVisibility = "Protected"
	SetupObjectVisibilityPackageProtected SetupObjectVisibility = "PackageProtected"
)