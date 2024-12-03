package user_interest

import (
	"net/http"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/user_interest"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/user_interest"
	"github.com/gin-gonic/gin"
)

type userInterestController struct {
	service service.UserInterestService
}

func NewUserInterestController(service service.UserInterestService) *userInterestController {
	return &userInterestController{
		service: service,
	}
}

func (uic *userInterestController) Create(c *gin.Context) {
	var userInterestRequest dto.UserInterestRequest
	if err := c.ShouldBindJSON(&userInterestRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid Input"})
		return
	}

	if err := uic.service.Create(&userInterestRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.InternalServerError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Interest added successfully " + userInterestRequest.InterestName,
	}))
}

func (uic *userInterestController) GetByID(c *gin.Context) {
	id := c.Param("id")

	result, err := uic.service.GetByID(id)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	if result == nil {
		errorhandler.HandleError(c, &errorhandler.NotFoundError{Message: "There is no interest in the db"})
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Request Succeed",
		Data:       result,
	}))
}

func (uic *userInterestController) GetAll(c *gin.Context) {
	result, err := uic.service.GetAll()
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Interest Found",
		Data:       result,
	}))
}
