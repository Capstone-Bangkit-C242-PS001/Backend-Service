package routes

import (
	controller "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/controller/course"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/course"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/course"
	"github.com/gin-gonic/gin"
)

func CourseRoutes(api gin.IRoutes) {
	courseRepository := repository.NewCourseRepository(database.DB)
	courseService := service.NewCourseService(courseRepository)
	courseController := controller.NewCourseController(courseService)

	api.POST("/course", courseController.Create)
	api.GET("/course/:id", courseController.GetByID)
	api.GET("/courses", courseController.GetAll)
}
