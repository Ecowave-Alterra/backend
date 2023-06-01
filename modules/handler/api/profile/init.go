package profile

import pc "github.com/berrylradianh/ecowave-go/modules/usecase/profile"

type ProfileHandler struct {
	profileUsecase pc.ProfileUsecase
}

func New(profileUsecase pc.ProfileUsecase) *ProfileHandler {
	return &ProfileHandler{
		profileUsecase,
	}
}
