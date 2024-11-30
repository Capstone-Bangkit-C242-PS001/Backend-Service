package routes

import (
	controller "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/controller/user"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/user"
	"github.com/gin-gonic/gin"
)

func UserRoute(api gin.IRoutes) {
	userRepository := repository.NewUserRepository(database.DB)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService)

	api.PUT("/user/:id", userController.Update)
}
