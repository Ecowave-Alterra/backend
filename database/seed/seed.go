package seed

import (
	at "github.com/berrylradianh/ecowave-go/modules/entity/admin"
	"gorm.io/gorm"
)

type Seed struct {
	Seed interface{}
}

func RegisterSeed(db *gorm.DB) Seed {
	return Seed{
		Seed: CreateAdmin(db),
	}
}

func DBSeed(db *gorm.DB) error {
	var admin at.Admin

	if err := db.Find(&admin).Error; err != nil {
		return err
	}

	if admin.Email == "" && admin.Password == "" {
		if err := db.Create(RegisterSeed(db).Seed).Error; err != nil {
			return err
		}
	}

	return nil
}
