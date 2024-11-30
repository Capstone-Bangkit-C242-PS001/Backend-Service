package user

import (
	"errors"

	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/user"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
	GetByID(id string) (*model.User, error)
	Update(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (ur *userRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User

	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Create(user *model.User) error {
	result := ur.db.Create(&user)

	return result.Error
}

func (ur *userRepository) GetByID(id string) (*model.User, error) {
	var user model.User

	if err := ur.db.Where("id = ?", id).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

func (ur *userRepository) Update(user *model.User) error {
	result := ur.db.Save(user)

	return result.Error
}
