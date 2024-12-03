package rating

import (
	courseDto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/course"
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/rating"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/rating"
	courseRepository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/course"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/rating"
	userRepository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
)

type RatingService interface {
	Create(userID int, req *dto.CreateRatingRequest) error
	GetByUserID(id int) (*dto.RatingResponse, error)
	GetRating(user_id, course_id int) (*dto.RatingDefaultResponse, error)
}

type ratingService struct {
	repository       repository.RatingRepository
	courseRepository courseRepository.CourseRepository
	userRepository   userRepository.UserRepository
}

func NewRatingService(
	repository repository.RatingRepository,
	courseRepository courseRepository.CourseRepository,
	userRepository userRepository.UserRepository,
) *ratingService {
	return &ratingService{
		repository:       repository,
		courseRepository: courseRepository,
		userRepository:   userRepository,
	}
}

func (rs *ratingService) Create(userID int, req *dto.CreateRatingRequest) error {
	course, err := rs.courseRepository.GetByID(req.CourseID)
	if err != nil {
		return err
	}

	rating := model.Rating{
		UserID:   userID,
		CourseID: course.ID,
		Rating:   req.Rating,
	}

	if err := rs.repository.Create(&rating); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}

func (rs *ratingService) GetByUserID(id int) (*dto.RatingResponse, error) {
	ratings, err := rs.repository.GetByUserID(id)
	if err != nil {
		return nil, err
	}

	var ratingCourses []dto.RatingCourses
	for _, rating := range ratings {
		course, err := rs.courseRepository.GetByID(rating.CourseID)
		if err != nil {
			return nil, err
		}
		ratingCourse := dto.RatingCourses{
			Course: courseDto.CourseResponse{
				ID:          course.ID,
				Title:       course.Title,
				Description: course.Description,
				Provider:    course.Provider,
				ProviderURL: course.ProviderURL,
				CoursePic:   *course.CoursePic,
				CreatedAt:   course.CreatedAt,
				UpdatedAt:   course.UpdatedAt,
			},
			Rating: rating.Rating,
		}

		ratingCourses = append(ratingCourses, ratingCourse)
	}

	return &dto.RatingResponse{
		UserID:        id,
		RatingCourses: ratingCourses,
	}, nil
}

func (rs *ratingService) GetRating(user_id, course_id int) (*dto.RatingDefaultResponse, error) {
	user, err := rs.userRepository.GetByID(user_id)
	if err != nil {
		return nil, err
	}

	course, err := rs.courseRepository.GetByID(course_id)
	if err != nil {
		return nil, err
	}

	result, err := rs.repository.GetRating(user.ID, course.ID)
	if err != nil {
		return nil, err
	}

	return &dto.RatingDefaultResponse{
		ID:        result.ID,
		CourseID:  result.CourseID,
		UserID:    result.UserID,
		Rating:    result.Rating,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}
