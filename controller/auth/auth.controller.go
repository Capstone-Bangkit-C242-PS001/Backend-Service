package auth

import (
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

func (ac *authController) Register(c *gin.Context) {
	var registerRequest dto.RegisterRequest

	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid input"})
		return
	}

	if err := ac.service.Register(&registerRequest); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusCreated,
		Message:    "User registered successfully",
	}))
}

func (ac *authController) Login(c *gin.Context) {
	var loginRequest dto.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: err.Error()})
		return
	}

	resp, err := ac.service.Login(&loginRequest)
	if err != nil {
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
