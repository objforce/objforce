package models

type SharedTo struct {
	AllCustomerPortalUsers      string   `json:"allCustomerPortalUsers,omitempty"`
	AllInternalUsers            string   `json:"allInternalUsers,omitempty"`
	AllPartnerUsers             string   `json:"allPartnerUsers,omitempty"`
	Group                       []string `json:"group,omitempty"`
	Groups                      []string `json:"groups,omitempty"`
	ManagerSubordinates         []string `json:"managerSubordinates,omitempty"`
	Managers                    []string `json:"managers,omitempty"`
	PortalRole                  []string `json:"portalRole,omitempty"`
	PortalRoleAndSubordinates   []string `json:"portalRoleAndSubordinates,omitempty"`
	Queue                       []string `json:"queue,omitempty"`
	Role                        []string `json:"role,omitempty"`
	RoleAndSubordinates         []string `json:"roleAndSubordinates,omitempty"`
	RoleAndSubordinatesInternal []string `json:"roleAndSubordinatesInternal,omitempty"`
	Roles                       []string `json:"roles,omitempty"`
	RolesAndSubordinates        []string `json:"rolesAndSubordinates,omitempty"`
	Territories                 []string `json:"territories,omitempty"`
	TerritoriesAndSubordinates  []string `json:"territoriesAndSubordinates,omitempty"`
	Territory                   []string `json:"territory,omitempty"`
	TerritoryAndSubordinates    []string `json:"territoryAndSubordinates,omitempty"`
}