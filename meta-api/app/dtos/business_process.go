package dtos

/**
The BusinessProcess metadata type enables you to display different picklist values for users based on their profile.This type extends the Metadata metadata type and inherits its fullName field.
Multiple business processes allow you to track separate sales, support, and lead lifecycles. A sales, support, lead, or solution process is assigned to a record type. The record type determines the user profiles that are associated with the business process. For more information, see “ Managing Multiple Business Processes ” in Salesforce Help.
 */

type BusinessProcess struct {
	*Metadata
	Description string `json:"description"`
	IsActive bool `json:"isActive"`
	NamespacePrefix string `json:"namespacePrefix"`
	Values []PicklistValue `json:"values"`
}