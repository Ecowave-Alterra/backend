package role

import (
	ut "github.com/berrylradianh/ecowave-go/modules/entity/user"
	"gorm.io/gorm"
)

type Role struct {
	*gorm.Model

	Role  string `json:"Role" form:"Role"`
	Users []ut.User
}
