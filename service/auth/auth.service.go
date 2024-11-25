package auth

import (
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/auth"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/user"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/utils"
)

type AuthService interface {
	Register(req *auth.RegisterRequest) error
}

type authService struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) *authService {
	return &authService{repository: repository}
}

func (as *authService) Register(req *auth.RegisterRequest) error {
	user, err := as.repository.FindByEmail(req.Email)

	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	if user != nil {
		return &errorhandler.BadRequestError{Message: "email already exists"}
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	registerUser := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := as.repository.Create(&registerUser); err != nil {
		return &errorhandler.InternalServerError{Message: err.Error()}
	}

	return nil
}