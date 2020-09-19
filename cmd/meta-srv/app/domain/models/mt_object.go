package models

import (
	_gorm "github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

type MTObject struct {
	ObjId string `json:"objId,omitempty" gorm:"primary_key"`
	OrgId string `json:"orgId,omitempty"`
	ObjName string `json:"objName,omitempty"`
	ActionOverrides []*ActionOverride	`json:"actionOverrides,omitempty"`
	AllowInChatterGroups bool `json:"allowInChatterGroups,omitempty"`
	BusinessProcesses []*BusinessProcess
	CompactLayoutAssignment string
	CompactLayouts []*CompactLayout
	CustomHelp string `json:"customHelp"`
	CustomHelpPage string `json:"customHelpPage"`
	CustomSettingsType string `json:"customSettingsType"`
	CustomSettingsVisibility string `json:"customSettingsVisibility"`
	DataStewardGroup string `json:"dataStewardGroup"`
	DataStewardUser string `json:"dataStewardUser"`
	DeploymentStatus string `json:"deploymentStatus"`
	Deprecated bool `json:"deprecated"`
	Description string `json:"description"`
	EnableActivities bool `json:"enableActivities"`
	EnableBulkApi bool `json:"enableBulkApi"`
	EnableDivisions bool `json:"enableDivisions"`
	EnableEnhancedLookup bool `json:"enableEnhancedLookup"`
	EnableFeeds bool `json:"enableFeeds"`
	EnableHistory              bool                    `json:"enableHistory"`
	EnableReports              bool                    `json:"enableReports"`
	EnableSearch               bool                    `json:"enableSearch"`
	EnableSharing              bool                    `json:"enableSharing"`
	EnableStreamingApi         bool                    `json:"enableStreamingApi"`
	EventType                  string                  `json:"eventType"`
	ExternalDataSource         string                  `json:"externalDataSource"`
	ExternalName               string                  `json:"externalName"`
	ExternalRepository         string                  `json:"externalRepository"`
	ExternalSharingModel       string                  `json:"externalSharingModel"`
	Fields                     []*MTField              `json:"fields" gorm:"FOREIGNKEY:ObjId;ASSOCIATION_FOREIGNKEY:ObjId"`
	FieldSets                  *FieldSet               `json:"fieldSets"`
	Gender                     string                  `json:"gender"`
	Household                  bool                    `json:"household"`
	HistoryRetentionPolicy     *HistoryRetentionPolicy `json:"historyRetentionPolicy"`
	Indexes                    []*Index                `json:"indexes"`
	Label                      string                  `json:"label"`
	ListViews                  []*ListView             `json:"listViews"`
	NamedFilter                []*NamedFilter          `json:"namedFilter"`
	NameField                  MTField                 `json:"nameField"`
	PluralLabel                string                  `json:"pluralLabel"`
	ProfileSearchLayouts       *ProfileSearchLayouts   `json:"profileSearchLayouts"`
	PublishBehavior            string                  `json:"publishBehavior"`
	RecordTypes                []*RecordType           `json:"recordTypes"`
	RecordTypeTrackFeedHistory bool                    `json:"recordTypeTrackFeedHistory"`
	RecordTypeTrackHistory     bool                    `json:"recordTypeTrackHistory"`
	SearchLayouts              *SearchLayouts          `json:"searchLayouts"`
	SharingModel               string                  `json:"sharingModel"`
	SharingReasons             *SharingReason          `json:"sharingReason"`
	SharingRecalculations      []*SharingRecalculation `json:"sharingRecalculations"`
	StartsWith string `json:"startsWith"`
	ValidationRules []*ValidationRule `json:"validationRules"`
	Visibility string `json:"visibility"`
	WebLinks []*WebLink `json:"webLinks"`
}

func (c *MTObject) BeforeCreate(scope *_gorm.Scope) error {
	scope.SetColumn("obj_id", uuid.NewV4().String())
	return nil
}