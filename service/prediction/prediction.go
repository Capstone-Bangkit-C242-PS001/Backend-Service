package prediction

import (
	"fmt"
	externalDTO "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/external"
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/prediction"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/external"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/course"
)

type PredictionService interface {
	Predict(userID int, data dto.PredictRequest) (*[]dto.PredictResponse, error)
}

type predictionService struct {
	courseRepository course.CourseRepository
	mlService        external.MLService
}

func NewPredictionService(courseRepository course.CourseRepository, mlService external.MLService) *predictionService {
	return &predictionService{
		courseRepository: courseRepository,
		mlService:        mlService,
	}
}

func (ps *predictionService) Predict(userID int, req dto.PredictRequest) (*[]dto.PredictResponse, error) {
	data := externalDTO.PredictRequest{
		UserID:   userID,
		Skillset: req.Skillsets,
	}

	predictedCourseIDs, err := ps.mlService.Predict(data)
	if err != nil {
		return nil, fmt.Errorf("failed to predict: %w", err)
	}

	courses, err := ps.courseRepository.GetByIds(*predictedCourseIDs)
	if err != nil {
		return nil, fmt.Errorf("failed to get courses: %w", err)
	}

	var response []dto.PredictResponse
	for _, c := range *courses {
		response = append(response, dto.PredictResponse{
			ID:          c.ID,
			Title:       c.Title,
			Description: c.Description,
			Provider:    c.Provider,
			ProviderURL: c.ProviderURL,
		})
	}

	return &response, nil
}
