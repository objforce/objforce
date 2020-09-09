package dtos


type FilterScope string

const (
	FilterScopeEverything FilterScope = "Everything"

	FilterScopeMine FilterScope = "Mine"

	FilterScopeQueue FilterScope = "Queue"

	FilterScopeDelegated FilterScope = "Delegated"

	FilterScopeMyTerritory FilterScope = "MyTerritory"

	FilterScopeMyTeamTerritory FilterScope = "MyTeamTerritory"

	FilterScopeTeam FilterScope = "Team"
)