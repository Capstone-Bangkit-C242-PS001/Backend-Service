package course

import (
	"mime/multipart"
	"time"
)

type CreateRequest struct {
	Title       string                `form:"title" binding:"required"`
	Description string                `form:"description" binding:"required"`
	Provider    string                `form:"provider" binding:"required"`
	ProviderURL string                `form:"provider_url" binding:"required"`
	CoursePic   *multipart.FileHeader `form:"course_pic" binding:"required"`
}

type CourseResponse struct {
	ID          int       `json:"id"`
	Title       string    `form:"title"`
	Description string    `form:"description"`
	Provider    string    `form:"provider"`
	ProviderURL string    `form:"provider_url"`
	CoursePic   string    `form:"course_pic"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
