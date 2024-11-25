package main

import (
	"fmt"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/config"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/database"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.GetConfig()
	database.LoadDB()

	server := gin.Default()
	api := server.Group("/api")

	api.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong pong pong",
		})
	})

	routes.AuthRouter(api)

	server.Run(fmt.Sprintf(":%v", cfg.APP_PORT))
}
