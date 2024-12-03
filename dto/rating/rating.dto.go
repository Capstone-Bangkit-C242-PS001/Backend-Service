package rating

import (
	"time"

	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/course"
)

type CreateRatingRequest struct {
	UserID   int `json:"user_id" binding:"required"`
	CourseID int `json:"course_id" binding:"required"`
	Rating   int `json:"rating" binding:"required"`
}

type RatingResponse struct {
	UserID        int             `json:"user_id"`
	RatingCourses []RatingCourses `json:"rating_courses"`
}

type RatingDefaultResponse struct {
	ID        string    `json:"id"`
	UserID    int       `json:"user_id"`
	CourseID  int       `json:"course_id"`
	Rating    int       `json:"rating"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RatingCourses struct {
	Course dto.CourseResponse `json:"course"`
	Rating int                `json:"rating"`
}
