package user

import (
	entity "cashback_info/internal/model/user"
	repoentity "cashback_info/internal/repository/model/user"
	repository "cashback_info/internal/repository/user"
	service "cashback_info/internal/service/password"
	"fmt"

	"github.com/google/uuid"
)

type UserService interface {
	GetUserByID(id uuid.UUID) (*entity.User, error)
	CreateUser(email, login, password string) (*uuid.UUID, error)
	DeleteUser(id uuid.UUID) error
}

type userService struct {
	repo            repository.UserRepository
	passwordService service.PasswordService
}

func NewUserService(repo repository.UserRepository, passwordService service.PasswordService) UserService {
	return &userService{repo: repo, passwordService: passwordService}
}

func (u *userService) CreateUser(email, login, password string) (*uuid.UUID, error) {
	hashPassword, err := u.passwordService.HashPassword(password)
	if err != nil {
		return nil, err
	}
	ID, err := u.repo.Create(repoentity.UserDB{
		Login:        login,
		Email:        email,
		PasswordHash: hashPassword,
		RoleType:     repoentity.Default, // TODO: Remove hard-code
	})

	return ID, err
}

func (u *userService) GetUserByID(id uuid.UUID) (*entity.User, error) {
	item, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	roleType := entity.GenerateRoleTypeFromString(string(item.RoleType))
	if roleType == nil {
		return nil, fmt.Errorf("invalid role type: %s", roleType)
	}

	return &entity.User{
		ID:       item.ID,
		Login:    item.Login,
		Email:    item.Email,
		RoleType: *roleType,
		Phone:    item.Phone,
	}, nil
}

func (u *userService) DeleteUser(id uuid.UUID) error {
	return u.repo.Delete(id)
}
