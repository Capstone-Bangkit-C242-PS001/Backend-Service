package auth

import (
	"errors"
	"net/http"

	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/auth"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/auth"

	"github.com/gin-gonic/gin"
)

type authController struct {
	service service.AuthService
}

func NewAuthController(service service.AuthService) *authController {
	return &authController{service: service}
}

// @Summary Register a new user
// @Description Allows a user to register with their name, email, password, and optional profile picture
// @Tags Authentication
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Name of the user"
// @Param email formData string true "Email of the user"
// @Param password formData string true "Password for the account"
// @Param profile_pic formData file false "Optional profile picture file"
// @Success 201 {object} RegisterResponse
// @Failure 400 "Invalid input"
// @Failure 500 "Internal server error"
// @Router /api/register [post]
func (ac *authController) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest

	if err := c.ShouldBind(&registerRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid input"})
		return
	}

	result, err := ac.service.Register(&registerRequest)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "User registered successfully",
		Data:       result,
	}))
}

// @Summary Login a user
// @Description Allows a user to log in by providing their email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param loginRequest body LoginRequest true "Login request payload"
// @Success 200 {object} LoginResponse
// @Failure 400 "Invalid input"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router /api/login [post]
func (ac *authController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	resp, err := ac.service.Login(&loginRequest)
	if err != nil {
		if errors.Is(err, service.UserNotFoundError) {
			errorhandler.HandleError(c, &errorhandler.NotFoundError{Message: err.Error()})
			return
		}

		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(
		response.ResponseParam{
			StatusCode: http.StatusOK,
			Message:    "User Login Successfully",
			Data:       resp,
		},
	))
}
