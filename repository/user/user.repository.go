package user

import (
	"errors"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *user.User) (*user.User, error)
	FindByEmail(email string) (*user.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindByEmail(email string) (*user.User, error) {
	var user user.User

	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Create(user *user.User) (*user.User, error) {
	result := ur.db.Create(user)

	if result.Error != nil {
		return nil, result.Error
	}

	return user, nil
}
