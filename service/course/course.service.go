package course

import (
	"path/filepath"

	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/course"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/course"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/course"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/utils"
	"github.com/google/uuid"
)

type CourseService interface {
	Create(req *dto.CreateRequest) error
	GetAll() ([]dto.CourseResponse, error)
	GetByID(id int) (*dto.CourseResponse, error)
}

type courseService struct {
	repository repository.CourseRepository
}

func NewCourseService(repository repository.CourseRepository) *courseService {
	return &courseService{
		repository: repository,
	}
}

func (cs *courseService) Create(req *dto.CreateRequest) error {
	course := model.Course{
		Title:       req.Title,
		Description: req.Description,
		Provider:    req.Provider,
		ProviderURL: req.ProviderURL,
	}

	if req.CoursePic != nil {
		ext := filepath.Ext(req.CoursePic.Filename)
		fileName := "course/" + uuid.New().String() + ext

		url, err := utils.UploadToGCS(req.CoursePic, fileName)
		if err != nil {
			return &errorhandler.BadRequestError{Message: err.Error()}
		}

		course.CoursePic = &url
	}

	if err := cs.repository.Create(&course); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (cs *courseService) GetByID(id int) (*dto.CourseResponse, error) {
	result, err := cs.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &dto.CourseResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
		Provider:    result.Provider,
		ProviderURL: result.ProviderURL,
		CoursePic:   *result.CoursePic,
		CreatedAt:   result.CreatedAt,
		UpdatedAt:   result.UpdatedAt,
	}, nil
}

func (cs *courseService) GetAll() ([]dto.CourseResponse, error) {
	courses, err := cs.repository.GetAll()
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	var responses []dto.CourseResponse
	for _, course := range courses {
		responses = append(responses, dto.CourseResponse{
			ID:          course.ID,
			Title:       course.Title,
			Description: course.Description,
			Provider:    course.Provider,
			ProviderURL: course.ProviderURL,
			CoursePic:   *course.CoursePic,
			CreatedAt:   course.CreatedAt,
			UpdatedAt:   course.UpdatedAt,
		})
	}

	return responses, nil
}
