package auth

import (
	repository "cashback_info/internal/repository/user"
	"cashback_info/internal/service/password"
	tokenservice "cashback_info/internal/service/token"
	"time"
)

type AuthService interface {
	Login(email string, password string) (string, time.Time, error)
}

type authService struct {
	userService     repository.UserRepository
	tokenService    tokenservice.TokenService
	passwordService password.PasswordService
}

func NewAuthService(userService repository.UserRepository, tokenService tokenservice.TokenService, passwordService password.PasswordService) AuthService {
	return &authService{userService: userService, tokenService: tokenService, passwordService: passwordService}
}

func (a *authService) Login(email string, password string) (string, time.Time, error) {
	user, err := a.userService.GetByEmail(email)
	if err != nil {
		return "", time.Time{}, err
	}

	if err := a.passwordService.VerifyPassword(password, user.PasswordHash); err != nil {
		return "", time.Time{}, err
	}

	token, expirationTime, err := a.tokenService.GenerateToken(user.ID)
	if err != nil {
		return "", time.Time{}, err
	}

	return token, expirationTime, nil
}
