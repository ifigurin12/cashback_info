package user

import (
	model "cashback_info/internal/repository/model/user"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id uuid.UUID) (*model.UserDB, error)
	GetByEmail(email string) (*model.UserDB, error)
	Create(user model.UserDB) (*uuid.UUID, error)
	Delete(id uuid.UUID) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id uuid.UUID) (*model.UserDB, error) {
	var user model.UserDB
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*model.UserDB, error) {
	var user model.UserDB
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user model.UserDB) (*uuid.UUID, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return nil, err
	}
	return &user.ID, nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	if err := r.db.Delete(&model.UserDB{}, id).Error; err != nil {
		return err
	}
	return nil
}
