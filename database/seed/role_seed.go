package seed

import (
	rt "github.com/berrylradianh/ecowave-go/modules/entity/role"
)

func CreateRoles() []*rt.Role {
	roles := []*rt.Role{
		{Role: "Admin"},
		{Role: "User"},
	}

	return roles
}
