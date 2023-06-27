package auth

import (
	"errors"
	"math/rand"
	"time"

	fp "github.com/berrylradianh/ecowave-go/helper/forgotpassword"
	pw "github.com/berrylradianh/ecowave-go/helper/password"
	vld "github.com/berrylradianh/ecowave-go/helper/validator"
	"github.com/berrylradianh/ecowave-go/middleware/jwt"
	ue "github.com/berrylradianh/ecowave-go/modules/entity/user"
)

func (ac *authUsecase) Register(request *ue.RegisterRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}

	hashedPassword, err := pw.HashPassword(request.Password)
	if err != nil {
		return err
	}

	request.Password = string(hashedPassword)

	err = ac.authRepo.CreateUser(request)
	if err != nil {
		return err
	}

	return nil
}
func (ac *authUsecase) RegisterGoogle(request *ue.RegisterGoogleRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}

	err := ac.authRepo.CreateUserGoogle(request)
	if err != nil {
		return err
	}

	return nil
}

func (ac *authUsecase) Login(request *ue.LoginRequest) (interface{}, uint, error) {
	if err := vld.Validation(request); err != nil {
		return nil, 0, err
	}

	response, password, role, err := ac.authRepo.Login(request.Email)

	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, 0, errors.New("Email atau password salah")
	}

	err = pw.VerifyPassword(password, request.Password)
	if err != nil {

		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, 0, errors.New("Email atau password salah")
	}

	token, err := jwt.CreateToken(int(response.ID), response.Email)
	if err != nil {
		return nil, 0, err
	}
	response.Token = token

	return response, role, nil
}
func (ac *authUsecase) LoginGoogle(request *ue.LoginGoogleRequest) (interface{}, uint, error) {
	if err := vld.Validation(request); err != nil {
		return nil, 0, err
	}
	response, role, err := ac.authRepo.LoginGoogleId(request.GoogleId)
	if err != nil {
		//lint:ignore ST1005 Reason for ignoring this linter
		return nil, 0, errors.New("Akun tidak ditemukan")
	}

	token, err := jwt.CreateToken(int(response.ID), response.Email)
	if err != nil {
		return nil, 0, err
	}

	response.Token = token

	return response, role, nil
}

func (ac *authUsecase) ForgotPassword(request ue.ForgotPassRequest) (string, error) {

	if err := vld.Validation(request); err != nil {
		return "", err
	}

	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return "", errors.New("Email tidak ditemukan")
	}

	if user.RoleId == 1 {
		return "", errors.New("Tidak diperbolehkan merubah data admin")
	}

	var alphaNumRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	codeVerRandRune := make([]rune, 6)
	for i := 0; i < 6; i++ {
		codeVerRandRune[i] = alphaNumRunes[rand.Intn(len(alphaNumRunes)-1)]
	}
	codeVerPassword := string(codeVerRandRune)

	_, err = ac.authRepo.GetUserRecovery(user.ID)
	if err != nil {
		err = ac.authRepo.UserRecovery(user.ID, codeVerPassword)
		if err != nil {
			return "", err
		}
	} else {
		err = ac.authRepo.UpdateUserRecovery(user.ID, codeVerPassword)
		if err != nil {
			return "", err
		}
	}

	err = fp.ForgotPassword(request.Email, codeVerPassword)
	if err != nil {
		return "", err
	}

	return user.Email, nil
}
func (ac *authUsecase) VerifOtp(request ue.VerifOtp) error {
	if err := vld.Validation(request); err != nil {
		return err
	}
	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return errors.New("Email tidak ditemukan")
	}
	userRecovery, err := ac.authRepo.GetUserRecovery(user.ID)
	if err != nil {
		return errors.New("Kode verifikasi tidak ditemukan")
	}

	expTime := userRecovery.CreatedAt.Add(15 * time.Minute)

	if !time.Now().Before(expTime) {
		return errors.New("Kode otp kadaluarsa")
	}

	if request.CodeOtp != userRecovery.Code {
		return errors.New("Kode verifikasi salah")
	}
	return nil
}
func (ac *authUsecase) ChangePassword(request ue.RecoveryRequest) error {
	if err := vld.Validation(request); err != nil {
		return err
	}
	user, err := ac.authRepo.GetUserByEmail(request.Email)
	if err != nil {
		return errors.New("Email tidak ditemukan")
	}

	hashedPassword, err := pw.HashPassword(request.Password)
	if err != nil {
		return err
	}
	request.Password = string(hashedPassword)
	err = ac.authRepo.ChangePassword(request)
	if err != nil {
		return err
	}
	err = ac.authRepo.DeleteUserRecovery(user.ID)
	if err != nil {
		return err
	}
	return nil
}
