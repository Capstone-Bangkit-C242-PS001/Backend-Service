package user_interest

import (
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/user_interest"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/user_interest"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user_interest"
)

type UserInterestService interface {
	Create(req *dto.UserInterestRequest) error
	GetByID(id string) (*dto.UserInterestResponse, error)
	GetAll() ([]dto.UserInterestResponse, error)
}

type userInterestService struct {
	repository repository.UserInterestRepository
}

func NewUserInterestService(repository repository.UserInterestRepository) *userInterestService {
	return &userInterestService{
		repository: repository,
	}
}

func (uis *userInterestService) Create(req *dto.UserInterestRequest) error {
	userInterest := model.UserInterest{
		InterestName: req.InterestName,
		Description:  req.Description,
	}

	return uis.repository.Create(&userInterest)
}

func (uis *userInterestService) GetByID(id string) (*dto.UserInterestResponse, error) {
	result, err := uis.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.UserInterestResponse{
		ID:           result.ID,
		InterestName: result.InterestName,
		Description:  result.Description,
		CreatedAt:    result.CreatedAt,
		UpdatedAt:    result.UpdatedAt,
	}, nil
}

func (uis *userInterestService) GetAll() ([]dto.UserInterestResponse, error) {
	userInterests, err := uis.repository.GetAll()
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	var responses []dto.UserInterestResponse
	for _, interest := range userInterests {
		responses = append(responses, dto.UserInterestResponse{
			ID:           interest.ID,
			InterestName: interest.InterestName,
			Description:  interest.Description,
			CreatedAt:    interest.CreatedAt,
			UpdatedAt:    interest.UpdatedAt,
		})
	}

	return responses, nil
}
