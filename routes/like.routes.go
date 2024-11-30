package routes

import (
	controller "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/controller/like"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	repository "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/repository/like"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/like"
	"github.com/gin-gonic/gin"
)

func LikeRoute(api gin.IRoutes) {
	likeRepository := repository.NewLikeRepository(database.DB)
	likeService := service.NewLikeService(likeRepository)
	likeController := controller.NewLikeController(likeService)

	api.POST("like", likeController.Like)
	api.POST("unlike", likeController.Unlike)
}
