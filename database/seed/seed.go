package seed

import (
	"gorm.io/gorm"
)

type Seed struct {
	Seed interface{}
}

func RegisterSeed(db *gorm.DB) []Seed {
	return []Seed{
		{Seed: CreateRoles()},
		{Seed: CreateUser()},
		{Seed: CreateUserDetail()},
		{Seed: CreateInformation()},
	}
}

func DBSeed(db *gorm.DB) error {
	for _, seed := range RegisterSeed(db) {
		var count int64
		if err := db.Model(seed.Seed).Count(&count).Error; err != nil {
			return err
		}

		if count == 0 {
			err := db.Debug().Create(seed.Seed).Error
			if err != nil {
				return err
			}
		}
	}
	return nil
}
