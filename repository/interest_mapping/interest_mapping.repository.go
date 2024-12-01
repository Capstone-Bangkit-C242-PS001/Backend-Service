package interest_mapping

import (
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/interest_mapping"
	"gorm.io/gorm"
)

type InterestMappingRepository interface {
	Create(interestMapping *model.InterestMapping) error
	GetByUserID(id string) ([]model.InterestMapping, error)
}

type interestMappingRepository struct {
	db *gorm.DB
}

func NewInterestMappingRepository(db *gorm.DB) *interestMappingRepository {
	return &interestMappingRepository{
		db: db,
	}
}

func (imr *interestMappingRepository) Create(interestMapping *model.InterestMapping) error {
	result := imr.db.Create(interestMapping)

	return result.Error
}

func (imr *interestMappingRepository) GetByUserID(id string) ([]model.InterestMapping, error) {
	var interestMapping []model.InterestMapping

	result := imr.db.Where("user_id = ?", id).Find(&interestMapping)

	if err := result.Error; err != nil {
		return nil, &errorhandler.NotFoundError{Message: err.Error()}
	}

	if len(interestMapping) == 0 {
		return nil, &errorhandler.NotFoundError{Message: "Interest Mapping Not Found"}
	}

	return interestMapping, nil
}
