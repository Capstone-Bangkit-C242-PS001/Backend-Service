package routes

import (
	controller "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/controller/user_interest"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user_interest"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/user_interest"
	"github.com/gin-gonic/gin"
)

func UserInterestRoute(api gin.IRoutes) {
	interest_repository := repository.NewUserInterestRepository(database.DB)
	interest_service := service.NewUserInterestService(interest_repository)
	interest_controller := controller.NewUserInterestController(interest_service)

	api.POST("/interest", interest_controller.Create)
	api.GET("/interest", interest_controller.GetAll)
	api.GET("/interest/:id", interest_controller.GetByID)
}
