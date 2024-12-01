package routes

import (
	controller "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/controller/interest_mapping"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/interest_mapping"
	userRepository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
	userInterestRepository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user_interest"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/interest_mapping"
	"github.com/gin-gonic/gin"
)

func InterestMappingRoutes(api gin.IRoutes) {
	userRepository := userRepository.NewUserRepository(database.DB)
	interestRepository := userInterestRepository.NewUserInterestRepository(database.DB)
	repository := repository.NewInterestMappingRepository(database.DB)
	service := service.NewInterestMappingService(repository, interestRepository, userRepository)
	controller := controller.NewInterestMappingController(service)

	api.POST("/interest-mapping/:id", controller.Create)
	api.GET("/interest-mapping/:id", controller.GetByUserID)
}
