package interest_mapping

import (
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/interest_mapping"
	userInterestDto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/user_interest"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/interest_mapping"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/interest_mapping"
	userRepository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
	userInterestRepository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user_interest"
)

type InterestMappingService interface {
	Create(userID int, req *dto.InterestMappingRequest) error
	GetByUserID(userID int) (*dto.InterestMappingResponse, error)
}

type interestMappingService struct {
	repository         repository.InterestMappingRepository
	interestRepository userInterestRepository.UserInterestRepository
	userRepository     userRepository.UserRepository
}

func NewInterestMappingService(
	repository repository.InterestMappingRepository,
	interestRepository userInterestRepository.UserInterestRepository,
	userRepository userRepository.UserRepository,
) *interestMappingService {
	return &interestMappingService{
		repository:         repository,
		interestRepository: interestRepository,
		userRepository:     userRepository,
	}
}

func (ims *interestMappingService) Create(userID int, req *dto.InterestMappingRequest) error {
	user, err := ims.userRepository.GetByID(userID)
	if err != nil {
		return &errorhandler.NotFoundError{Message: err.Error()}
	}

	for _, interestID := range req.InterestIDs {
		_, err := ims.interestRepository.GetByID(interestID)
		if err != nil {
			if _, ok := err.(*errorhandler.NotFoundError); ok {
				return &errorhandler.NotFoundError{Message: "Interest ID " + interestID + " is invalid"}
			}
			return err
		}
	}

	for _, interestID := range req.InterestIDs {
		mapping := model.InterestMapping{
			UserID:     user.ID,
			InterestID: interestID,
		}

		if err := ims.repository.Create(&mapping); err != nil {
			return err
		}
	}

	return nil
}

func (ims *interestMappingService) GetByUserID(userID int) (*dto.InterestMappingResponse, error) {
	mappings, err := ims.repository.GetByUserID(userID)
	if err != nil {
		return nil, err
	}

	var interests []userInterestDto.UserInterestResponse
	for _, mapping := range mappings {
		interest, err := ims.interestRepository.GetByID(mapping.InterestID)
		if err != nil {
			return nil, err
		}
		interests = append(interests, userInterestDto.UserInterestResponse{
			ID:           interest.ID,
			InterestName: interest.InterestName,
			Description:  interest.Description,
			CreatedAt:    interest.CreatedAt,
			UpdatedAt:    interest.UpdatedAt,
		})
	}

	return &dto.InterestMappingResponse{
		UserID:    userID,
		Interests: interests,
	}, nil
}
