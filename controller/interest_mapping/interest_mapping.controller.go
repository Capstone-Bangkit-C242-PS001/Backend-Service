package interest_mapping

import (
	"net/http"
	"strconv"

	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/interest_mapping"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/interest_mapping"
	"github.com/gin-gonic/gin"
)

type interestMappingController struct {
	service service.InterestMappingService
}

func NewInterestMappingController(service service.InterestMappingService) *interestMappingController {
	return &interestMappingController{
		service: service,
	}
}

func (imc *interestMappingController) Create(c *gin.Context) {
	id := c.Param("id")
	var interestMappingRequest dto.InterestMappingRequest
	if err := c.ShouldBindJSON(&interestMappingRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid input"})
		return
	}

	idNumber, err := strconv.Atoi(id)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid user ID format"})
	}

	if err := imc.service.Create(idNumber, &interestMappingRequest); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Mapping Successfully",
	}))
}

func (imc *interestMappingController) GetByUserID(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid user ID format"})
	}

	result, err := imc.service.GetByUserID(idNumber)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Get mapping successfully with user id " + id,
		Data:       result,
	}))
}
