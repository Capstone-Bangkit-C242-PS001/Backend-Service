package auth

import "mime/multipart"

type RegisterRequest struct {
	Name       string                `form:"name" binding:"required"`
	Email      string                `form:"email" binding:"required"`
	Password   string                `form:"password" binding:"required"`
	ProfilePic *multipart.FileHeader `form:"profile_pic"`
}

type RegisterResponse struct {
	Name       string `json:"name"`
	Email      string `json:"email"`
	ProfilePic string `json:"profile_pic"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Token string `json:"token"`
}
