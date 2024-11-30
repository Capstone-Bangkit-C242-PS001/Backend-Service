package like

import (
	"errors"
	"fmt"
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/like"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/like"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/like"
	"gorm.io/gorm"
)

type LikeService interface {
	Like(userID string, req dto.LikeRequest) error
	Unlike(userID string, req dto.LikeRequest) error
}

type likeService struct {
	repository repository.LikeRepository
}

func NewLikeService(repository repository.LikeRepository) *likeService {
	return &likeService{
		repository: repository,
	}
}

func (ls *likeService) Like(userID string, req dto.LikeRequest) error {
	// Get all liked courses
	likedCourses, err := ls.repository.Get(userID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to get liked courses: %w", err)
		}
	}

	// Check if the course is already liked
	for _, likedCourse := range *likedCourses {
		if likedCourse.CourseID == req.CourseID {
			return nil
		}
	}

	// Create like
	if err := ls.repository.Create(model.Like{
		UserID:   userID,
		CourseID: req.CourseID,
	}); err != nil {
		return fmt.Errorf("failed to like course: %w", err)
	}

	return nil
}

var LikeNotFoundError = errors.New("like not found")

func (ls *likeService) Unlike(userID string, req dto.LikeRequest) error {
	// Check if the course is already liked
	likedCourses, err := ls.repository.Get(userID)
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return fmt.Errorf("failed to get liked courses: %w", err)
		}
	}

	// Find like ID to be deleted
	var likeID *string
	for _, likedCourse := range *likedCourses {
		if likedCourse.CourseID == req.CourseID {
			likeID = &likedCourse.ID
		}
	}

	if likeID == nil {
		return LikeNotFoundError
	}

	// Unlike
	if err := ls.repository.Delete(*likeID); err != nil {
		return fmt.Errorf("failed to unlike course: %w", err)
	}

	return nil
}
