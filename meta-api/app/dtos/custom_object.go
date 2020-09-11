package dtos


type CustomObject struct {
	*Metadata
	ObjId                    string            `json:"objId"`
	ObjName					string				`json:"objName"`
	OrgId					string				`json:"orgId"`
	ActionOverrides          []*ActionOverride `json:"actionOverrides,omitempty"`
	AllowInChatterGroups     bool              `json:"allowInChatterGroups,omitempty"`
	BusinessProcesses        []*BusinessProcess
	CompactLayoutAssignment  string
	CompactLayouts           []*CompactLayout
	CustomHelp               string `json:"customHelp"`
	CustomHelpPage           string `json:"customHelpPage"`
	CustomSettingsType       CustomSettingsType `json:"customSettingsType"`
	CustomSettingsVisibility CustomSettingsVisibility `json:"customSettingsVisibility"`
	DataStewardGroup         string `json:"dataStewardGroup"`
	DataStewardUser string `json:"dataStewardUser"`
	DeploymentStatus DeploymentStatus `json:"deploymentStatus"`
	Deprecated bool `json:"deprecated"`
	Description string `json:"description"`
	EnableActivities bool `json:"enableActivities"`
	EnableBulkApi bool `json:"enableBulkApi"`
	EnableDivisions bool `json:"enableDivisions"`
	EnableEnhancedLookup bool `json:"enableEnhancedLookup"`
	EnableFeeds bool `json:"enableFeeds"`
	EnableHistory bool `json:"enableHistory"`
	EnableReports bool `json:"enableReports"`
	EnableSearch bool `json:"enableSearch"`
	EnableSharing bool `json:"enableSharing"`
	EnableStreamingApi bool `json:"enableStreamingApi"`
	EventType PlatformEventType `json:"eventType"`
	ExternalDataSource string `json:"externalDataSource"`
	ExternalName string `json:"externalName"`
	ExternalRepository string `json:"externalRepository"`
	ExternalSharingModel SharingModel `json:"externalSharingModel"`
	Fields []CustomField `json:"fields"`
	FieldSets FieldSet `json:"fieldSets"`
	Gender Gender `json:"gender"`
	Household bool `json:"household"`
	HistoryRetentionPolicy HistoryRetentionPolicy `json:"historyRetentionPolicy"`
	Indexes []Index `json:"indexes"`
	Label string `json:"label"`
	ListViews []*ListView `json:"listViews"`
	NamedFilter []*NamedFilter `json:"namedFilter"`
	NameField CustomField `json:"nameField"`
	PluralLabel string `json:"pluralLabel"`
	ProfileSearchLayouts *ProfileSearchLayouts `json:"profileSearchLayouts"`
	PublishBehavior *PlatformEventPublishBehavior `json:"publishBehavior"`
	RecordTypes []*RecordType `json:"recordTypes"`
	RecordTypeTrackFeedHistory bool `json:"recordTypeTrackFeedHistory"`
	RecordTypeTrackHistory bool `json:"recordTypeTrackHistory"`
	SearchLayouts *SearchLayouts `json:"searchLayouts"`
	SharingModel SharingModel `json:"sharingModel"`
	SharingReasons *SharingReason `json:"sharingReason"`
	SharingRecalculations []*SharingRecalculation `json:"sharingRecalculations"`
	StartsWith *StartsWith `json:"startsWith"`
	ValidationRules []*ValidationRule `json:"validationRules"`
	Visibility SetupObjectVisibility `json:"visibility"`
	WebLinks []*WebLink `json:"webLinks"`
}