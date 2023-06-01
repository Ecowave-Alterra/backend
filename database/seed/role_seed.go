package seed

import (
	re "github.com/berrylradianh/ecowave-go/modules/entity/role"
)

func CreateRoles() []*re.Role {
	roles := []*re.Role{
		{Role: "Admin"},
		{Role: "User"},
	}

	return roles
}
