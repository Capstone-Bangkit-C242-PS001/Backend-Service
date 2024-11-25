package routes

import (
	controller "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/controller/auth"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/auth"
	"github.com/gin-gonic/gin"
)

func AuthRouter(api *gin.RouterGroup) {
	userRepository := user.NewUserRepository(database.DB)
	authService := service.NewAuthService(userRepository)
	authController := controller.NewAuthController(authService)

	api.POST("/register", authController.Register)
}