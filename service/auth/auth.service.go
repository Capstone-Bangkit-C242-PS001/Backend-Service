package auth

import (
	"path/filepath"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/auth"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/user"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/utils"
	"github.com/google/uuid"
)

type AuthService interface {
	Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error)
	Login(req *auth.LoginRequest) (*auth.LoginResponse, error)
}

type authService struct {
	repository repository.UserRepository
}

func NewAuthService(repository repository.UserRepository) *authService {
	return &authService{repository: repository}
}

func (as *authService) Register(req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	user, err := as.repository.FindByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, &errorhandler.BadRequestError{Message: "email already exists"}
	}

	hashedPassword, err := utils.HashPassword(req.Password)

	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	registerUser := model.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if req.ProfilePic != nil {
		ext := filepath.Ext(req.ProfilePic.Filename)
		fileName := "profile/" + uuid.New().String() + ext

		url, err := utils.UploadToGCS(req.ProfilePic, fileName)
		if err != nil {
			return nil, &errorhandler.BadRequestError{Message: err.Error()}
		}

		registerUser.ProfilePic = &url
	}

	if err := as.repository.Create(&registerUser); err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	emptyString := ""
	if registerUser.ProfilePic == nil {
		registerUser.ProfilePic = &emptyString
	}

	return &auth.RegisterResponse{
		Name:       registerUser.Name,
		Email:      registerUser.Email,
		ProfilePic: *registerUser.ProfilePic,
	}, nil
}

func (as *authService) Login(req *auth.LoginRequest) (*auth.LoginResponse, error) {
	user, err := as.repository.FindByEmail(req.Email)

	if err != nil {
		return nil, err
	}

	if !utils.VerifyPassword(user.Password, req.Password) {
		return nil, &errorhandler.NotFoundError{Message: "Invalid email or password"}
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	return &auth.LoginResponse{
		ID:    user.ID,
		Name:  user.Name,
		Token: token,
	}, nil
}
