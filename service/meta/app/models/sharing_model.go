package models

type SharingModel string

const (
	SharingModelPrivate                   SharingModel = "Private"
	SharingModelRead                      SharingModel = "Read"
	SharingModelReadWrite                 SharingModel = "ReadWrite"
	SharingModelReadWriteTransfer         SharingModel = "ReadWriteTransfer"
	SharingModelFullAccess                SharingModel = "FullAccess"
	SharingModelControlledByParent        SharingModel = "ControlledByParent"
	SharingModelControlledByCampaign      SharingModel = "ControlledByCampaign"
	SharingModelControlledByLeadOrContact SharingModel = "ControlledByLeadOrContact"
)
