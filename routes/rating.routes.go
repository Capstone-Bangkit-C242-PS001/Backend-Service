package routes

import (
	controller "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/controller/rating"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	courseRepository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/course"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/rating"
	userRepository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/user"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/rating"
	"github.com/gin-gonic/gin"
)

func RatingRoutes(api gin.IRoutes) {
	courseRepository := courseRepository.NewCourseRepository(database.DB)
	userRepository := userRepository.NewUserRepository(database.DB)
	repository := repository.NewRatingRepository(database.DB)
	service := service.NewRatingService(repository, courseRepository, userRepository)
	controller := controller.NewRatingController(service)

	api.POST("/rating", controller.Create)
	api.GET("/rating/:id", controller.GetByUserID)
	api.GET("/rating", controller.GetRating)
}
