package entities

type LookupFilter struct {
	Active        bool         `json:"active"` // 代表是否激活
	BooleanFilter string       `json:"booleanFilter"`
	Description   string       `json:"description"`
	ErrorMessage  string       `json:"errorMessage"`
	FilterItems   []FilterItem `json:"filterItems"`
	InfoMessage   string       `json:"infoMessage"` // The information message displayed on the page. Use to describe things the user might not understand, such as why certain items are excluded in the lookup filter
}
