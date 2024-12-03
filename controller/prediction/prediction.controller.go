package prediction

import (
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/prediction"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/prediction"
	"github.com/gin-gonic/gin"
	"net/http"
)

type predictionController struct {
	service service.PredictionService
}

func NewPredictionController(service service.PredictionService) *predictionController {
	return &predictionController{
		service: service,
	}
}

func (pc *predictionController) Predict(c *gin.Context) {
	var courseRequest dto.PredictRequest
	if err := c.ShouldBind(&courseRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid Input"})
		return
	}

	userID := c.Value("userID").(int)

	result, err := pc.service.Predict(userID, courseRequest)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Data:       result,
	}))
}
