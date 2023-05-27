package admin

import (
	at "github.com/berrylradianh/ecowave-go/modules/entity/admin"
	ar "github.com/berrylradianh/ecowave-go/modules/repository/admin"
)

type AdminUsecase interface {
	LoginAdmin(admin *at.Admin) (string, error)
}

type adminUsecase struct {
	adminRepo ar.AdminRepo
}

func New(adminRepo ar.AdminRepo) *adminUsecase {
	return &adminUsecase{
		adminRepo,
	}
}
