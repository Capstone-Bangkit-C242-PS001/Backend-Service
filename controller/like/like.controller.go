package like

import (
	"errors"
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/like"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/like"
	"github.com/gin-gonic/gin"
	"net/http"
)

type likeController struct {
	service service.LikeService
}

func NewLikeController(service service.LikeService) *likeController {
	return &likeController{
		service: service,
	}
}

func (lc *likeController) Like(c *gin.Context) {
	var likeRequest dto.LikeRequest
	if err := c.ShouldBind(&likeRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid Input"})
		return
	}

	userID := c.Value("userID").(string)

	err := lc.service.Like(userID, likeRequest)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "User liked successfully",
	})
}

func (lc *likeController) Unlike(c *gin.Context) {
	var likeRequest dto.LikeRequest
	if err := c.ShouldBind(&likeRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid Input"})
		return
	}

	userID := c.Value("userID").(string)

	err := lc.service.Unlike(userID, likeRequest)
	if err != nil {
		if errors.Is(err, service.LikeNotFoundError) {
			errorhandler.HandleError(c, &errorhandler.NotFoundError{Message: err.Error()})
			return
		}
		errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "User unliked successfully",
	})
}
