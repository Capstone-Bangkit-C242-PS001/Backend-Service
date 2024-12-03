package routes

import (
	controller "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/controller/prediction"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/external"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/course"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/prediction"
	"github.com/gin-gonic/gin"
)

func PredictionRoutes(api gin.IRoutes) {
	courseRepository := repository.NewCourseRepository(database.DB)
	mlService := external.NewMLService()
	predictionService := service.NewPredictionService(courseRepository, mlService)
	predictionController := controller.NewPredictionController(predictionService)

	api.POST("", predictionController.Predict)
}
