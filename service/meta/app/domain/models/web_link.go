package models

type WebLink struct {
	Availability WebLinkAvailability `json:"availability"`
	Description string `json:"description"`
	DisplayType WebLinkDisplayType `json:"displayType"`
	EncodingKey Encoding `json:"encodingKey"`
	HasMenubar bool `json:"hasMenubar"`
	HasScrollbars bool `json:"hasScrollbars"`
	HasToolbar bool `json:"hasToolbar"`
	Height int `json:"height"`
	IsResizable bool `json:"isResizable"`
	LinkType WebLinkType `json:"linkType"`
	MasterLabel string `json:"masterLabel"`
	OpenType WebLinkWindowType `json:"openType"`
	Page                string               `json:"page,omitempty"`
	Position            *WebLinkPosition     `json:"position,omitempty"`
	Protected           bool                 `json:"protected,omitempty"`
	RequireRowSelection bool                 `json:"requireRowSelection,omitempty"`
	Scontrol            string               `json:"scontrol,omitempty"`
	ShowsLocation       bool                 `json:"showsLocation,omitempty"`
	ShowsStatus         bool                 `json:"showsStatus,omitempty"`
	Url                 string               `json:"url,omitempty"`
	Width               int32                `json:"width,omitempty"`
}

type WebLinkAvailability string

const(
	WebLinkAvailabilityOnline WebLinkAvailability = "online"
	WebLinkAvailabilityOffline WebLinkAvailability = "offline"
)

type WebLinkDisplayType string

const(
	WebLinkDisplayTypeLink WebLinkDisplayType = "link"
	WebLinkDisplayTypeButton WebLinkDisplayType = "button"
	WebLinkDisplayTypeMassActionButton WebLinkDisplayType = "massActionButton"
)

type Encoding string

const (
	EncodingUTF8 Encoding = "UTF8"

	EncodingISO88591 Encoding = "ISO88591"

	EncodingShiftJIS Encoding = "ShiftJIS"

	EncodingISO2022JP Encoding = "ISO2022JP"

	EncodingEUCJP Encoding = "EUCJP"

	EncodingKsc56011987 Encoding = "ksc56011987"

	EncodingBig5 Encoding = "Big5"

	EncodingGB2312 Encoding = "GB2312"

	EncodingBig5HKSCS Encoding = "Big5HKSCS"

	EncodingXSJIS0213 Encoding = "xSJIS0213"
)

type WebLinkType string
const(
	WebLinkTypeUrl WebLinkType = "url"
	WebLinkTypeSControl WebLinkType = "sControl"
	WebLinkTypeJavascript WebLinkType = "javascript"
	WebLinkTypePage WebLinkType = "page"
	WebLinkTypeFlow WebLinkType = "flow"
)

type WebLinkWindowType string
const(
	WebLinkWindowTypeNewWindow WebLinkWindowType = "newWindow"
	WebLinkWindowTypeSidebar WebLinkWindowType = "sidebar"
	WebLinkWindowTypeNoSidebar WebLinkWindowType = "noSidebar"
	WebLinkWindowTypeReplace WebLinkWindowType = "replace"
	WebLinkWindowTypeOnClickJavaScript WebLinkWindowType = "onClickJavaScript"
)

type WebLinkPosition string
const (
	WebLinkPositionFullScreen WebLinkPosition = "fullScreen"

	WebLinkPositionNone WebLinkPosition = "none"

	WebLinkPositionTopLeft WebLinkPosition = "topLeft"
)