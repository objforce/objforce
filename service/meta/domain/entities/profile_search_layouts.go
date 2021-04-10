package entities

/**
Represents a user profile’s search results layouts for an object. ProfileSearchLayouts are similar to SearchLayouts. However, with profile-specific layouts, each user profile can have a different search results layout for an object.
*/

type ProfileSearchLayouts struct {
	ProfileName []string `json:"profileName"`
	Fields      []string `json:"fields"`
}
