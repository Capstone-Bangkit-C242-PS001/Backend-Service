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
