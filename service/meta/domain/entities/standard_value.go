package entities

/**
referer: https://developer.salesforce.com/docs/atlas.en-us.api_meta.meta/api_meta/meta_customvalue.htm
*/
type StandardValue struct {
	AllowEmail       bool             `json:"allowEmail"`
	Closed           bool             `json:"closed"`
	Converted        bool             `json:"converted"`
	CssExposed       bool             `json:"cssExposed"`
	ForecastCategory ForecastCategory `json:"forecastCategory"`
	HighPriority     bool             `json:"highPriority"`
	Probability      int              `json:"probability"`
	ReverseRole      string           `json:"reverseRole"`
	Reviewed         bool             `json:"reviewed"`
	Won              bool             `json:"won"`
}
