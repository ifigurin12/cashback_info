package password

import "golang.org/x/crypto/bcrypt"

type PasswordService interface {
	HashPassword(password string) (string, error)
	VerifyPassword(password string, hashedPassword string) error
}

type passwordService struct {
}

func NewPasswordService() PasswordService {
	return &passwordService{}
}

func (p *passwordService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// VerifyPassword проверяет пароль
func (p *passwordService) VerifyPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
