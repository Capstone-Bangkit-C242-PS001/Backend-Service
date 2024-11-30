package user

import (
	"path/filepath"

	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/user"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/utils"
	"github.com/google/uuid"
)

type UserService interface {
	Update(req *dto.UpdateRequest, id string) (*dto.UpdateResponse, error)
}

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{
		repository: repository,
	}
}

func (us *userService) Update(req *dto.UpdateRequest, id string) (*dto.UpdateResponse, error) {
	user, err := us.repository.GetByID(id)

	if err != nil {
		return nil, &errorhandler.NotFoundError{Message: "User Not Found"}
	}

	if req.Name != nil {
		user.Name = *req.Name
	}

	if req.Password != nil {
		hashedPassword, err := utils.HashPassword(*req.Password)
		if err != nil {
			return nil, &errorhandler.InternalServerError{Message: err.Error()}
		}
		user.Password = hashedPassword
	}

	if req.ProfilePic != nil {
		ext := filepath.Ext(req.ProfilePic.Filename)
		fileName := "profile/" + uuid.New().String() + ext

		url, err := utils.UploadToGCS(req.ProfilePic, fileName)
		if err != nil {
			return nil, &errorhandler.BadRequestError{Message: err.Error()}
		}

		user.ProfilePic = &url
	}

	if err := us.repository.Update(user); err != nil {
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	return &dto.UpdateResponse{
		Name:       user.Name,
		Email:      user.Email,
		ProfilePic: *user.ProfilePic,
	}, nil
}
