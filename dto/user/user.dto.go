package user

import (
	"mime/multipart"
	"time"
)

type CreateRequest struct {
	Name       string                `form:"name" binding:"required"`
	Email      string                `form:"email" binding:"required"`
	Password   string                `form:"password" binding:"required"`
	ProfilePic *multipart.FileHeader `form:"profile_pic"`
}

type CreateResponse struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	ProfilePic string    `json:"profile_pic"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UpdateRequest struct {
	Name       *string               `form:"name"`
	Password   *string               `form:"password"`
	ProfilePic *multipart.FileHeader `form:"profile_pic"`
}

type UpdateResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProfilePic string `json:"profile_pic"`
}
