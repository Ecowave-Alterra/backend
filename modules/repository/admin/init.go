package admin

import (
	at "github.com/berrylradianh/ecowave-go/modules/entity/admin"
	"gorm.io/gorm"
)

type AdminRepo interface {
	LoginAdmin(admin *at.Admin) (string, error)
}

type adminRepo struct {
	db *gorm.DB
}

func New(db *gorm.DB) AdminRepo {
	return &adminRepo{
		db,
	}
}
