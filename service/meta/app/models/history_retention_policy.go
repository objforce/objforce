package models

/**
Represents the policy for retaining field history data. By setting a policy, you can specify the number of months you want to maintain field history in Salesforce and the number of years that you want to retain field history in the archive.
This component is only available to users with the “RetainFieldHistory” permission.
*/
type HistoryRetentionPolicy struct {
	ArchiveAfterMonths    int    `json:"archiveAfterMonths"`
	ArchiveRetentionYears int    `json:"archiveRetentionYears"`
	Description           string `json:"description"`
	GracePeriodDays       int    `json:"gracePeriodDays"`
}
