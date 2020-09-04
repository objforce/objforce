package dtos

type PicklistValue struct {
	AllowEmail bool `json:"allowEmail"`
	Closed bool `json:"closed"`
	ControllingFieldValues []string `json:"controllingFieldValues"`
	Converted bool `json:"converted"`
	CssExposed bool `json:"cssExposed"`
	ForecastCategory ForecastCategories `json:"forecastCategory"`
	HighPriority bool `json:"highPriority"`
	Probability int `json:"probability"`
	ReverseRole string `json:"reverseRole"`
	Reviewed bool `json:"reviewed"`
	Won bool `json:"won"`
}