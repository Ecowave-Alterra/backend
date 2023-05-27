package admin

import (
	at "github.com/berrylradianh/ecowave-go/modules/entity/admin"
)

func (ar *adminRepo) LoginAdmin(admin *at.Admin) (string, error) {
	var adminDB *at.Admin

	if err := ar.db.Where("email = ?", admin.Email).First(&adminDB).Error; err != nil {
		return "", err
	}

	return adminDB.Password, nil
}
