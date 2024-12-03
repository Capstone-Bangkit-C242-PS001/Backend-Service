package rating

import (
	"errors"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/rating"
	"gorm.io/gorm"
)

type RatingRepository interface {
	Create(rating *model.Rating) error
	GetByUserID(id int) ([]model.Rating, error)
	GetRating(user_id, course_id int) (*model.Rating, error)
}

type ratingRepository struct {
	db *gorm.DB
}

func NewRatingRepository(db *gorm.DB) *ratingRepository {
	return &ratingRepository{
		db: db,
	}
}

func (rp *ratingRepository) Create(rating *model.Rating) error {
	return rp.db.Create(rating).Error
}

func (rp *ratingRepository) GetByUserID(id int) ([]model.Rating, error) {
	var ratings []model.Rating
	result := rp.db.Where("user_id = ?", id).Find(&ratings)

	if err := result.Error; err != nil {
		return nil, &errorhandler.NotFoundError{Message: err.Error()}
	}

	if len(ratings) == 0 {
		return nil, &errorhandler.NotFoundError{Message: "Rating Not Found"}
	}

	return ratings, nil
}

func (rp *ratingRepository) GetRating(user_id, course_id int) (*model.Rating, error) {
	var rating model.Rating

	result := rp.db.Where("user_id = ? and course_id = ?", user_id, course_id).Find(&rating)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: "User interest not found"}
		}
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	return &rating, nil
}
