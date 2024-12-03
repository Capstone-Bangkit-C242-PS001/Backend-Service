package user_interest

import (
	"errors"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/user_interest"
	"gorm.io/gorm"
)

type UserInterestRepository interface {
	Create(userInterest *model.UserInterest) error
	GetByID(id string) (*model.UserInterest, error)
	GetAll() ([]model.UserInterest, error)
}

type userInterestRepository struct {
	db *gorm.DB
}

func NewUserInterestRepository(db *gorm.DB) *userInterestRepository {
	return &userInterestRepository{
		db: db,
	}
}

func (uir *userInterestRepository) Create(userInterest *model.UserInterest) error {
	result := uir.db.Create(userInterest)

	return result.Error
}

func (uir *userInterestRepository) GetByID(id string) (*model.UserInterest, error) {
	var userInterest model.UserInterest
	result := uir.db.Where("id = ?", id).First(&userInterest)

	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "User interest not found"}
		}
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	return &userInterest, nil
}

func (uir *userInterestRepository) GetAll() ([]model.UserInterest, error) {
	var userInterest []model.UserInterest
	result := uir.db.Find(&userInterest)

	return userInterest, result.Error
}
