package auth

import (
	entityuser "cashback_info/internal/model/user"
	repository "cashback_info/internal/repository/user"
	"cashback_info/internal/service/password"
	tokenservice "cashback_info/internal/service/token"
)

type AuthService interface {
	Login(email string, password string) (*entityuser.Token, error)
}

type authService struct {
	userService     repository.UserRepository
	tokenService    tokenservice.TokenService
	passwordService password.PasswordService
}

func NewAuthService(userService repository.UserRepository, tokenService tokenservice.TokenService, passwordService password.PasswordService) AuthService {
	return &authService{userService: userService, tokenService: tokenService, passwordService: passwordService}
}

func (a *authService) Login(email string, password string) (*entityuser.Token, error) {
	user, err := a.userService.GetByEmail(email)
	if err != nil {
		return nil, err
	}

	if err := a.passwordService.VerifyPassword(password, user.PasswordHash); err != nil {
		return nil, err
	}

	token, expirationTime, err := a.tokenService.GenerateToken(user.ID)
	if err != nil {
		return nil, err
	}

	return &entityuser.Token{Token: token, UserID: user.ID, ExpirationTime: expirationTime}, nil
}
