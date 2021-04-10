package models

type SearchLayouts struct {
	CustomTabListAdditionalFields      []string `json:"customTabListAdditionalFields,omitempty"`
	ExcludedStandardButtons            []string `json:"excludedStandardButtons,omitempty"`
	ListViewButtons                    []string `json:"listViewButtons,omitempty"`
	LookupDialogsAdditionalFields      []string `json:"lookupDialogsAdditionalFields,omitempty"`
	LookupFilterFields                 []string `json:"lookupFilterFields,omitempty"`
	LookupPhoneDialogsAdditionalFields []string `json:"lookupPhoneDialogsAdditionalFields,omitempty"`
	SearchFilterFields                 []string `json:"searchFilterFields,omitempty"`
	SearchResultsAdditionalFields      []string `json:"searchResultsAdditionalFields,omitempty"`
	SearchResultsCustomButtons         []string `json:"searchResultsCustomButtons,omitempty"`
}
