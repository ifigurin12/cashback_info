package user

import (
	model "cashback_info/internal/model/user"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id string) (*model.User, error)
	Create(user model.User) error
	Update(user model.User) error
	Delete(id string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id string) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) Create(user model.User) error {
	if err := r.db.Create(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Update(user model.User) error {
	if err := r.db.Save(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) Delete(id string) error {
	if err := r.db.Delete(&model.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
