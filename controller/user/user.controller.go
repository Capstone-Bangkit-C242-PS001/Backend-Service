package user

import (
	"net/http"
	"strconv"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/user"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/user"
	"github.com/gin-gonic/gin"
)

type userController struct {
	service service.UserService
}

func NewUserController(service service.UserService) *userController {
	return &userController{
		service: service,
	}
}

func (uc *userController) Update(c *gin.Context) {
	id := c.Param("id")
	var updateRequest dto.UpdateRequest
	if err := c.ShouldBind(&updateRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid Input"})
		return
	}

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid user ID format"})
		return
	}

	result, err := uc.service.Update(&updateRequest, idNumber)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "User updated successfully",
		Data:       result,
	})
}
