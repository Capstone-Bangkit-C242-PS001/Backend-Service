package errorhandler

import (
	"net/http"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, err error) {
	var statusCode int

	switch err.(type) {
	case *NotFoundError:
		statusCode = http.StatusNotFound
	case *BadRequestError:
		statusCode = http.StatusBadRequest
	case *InternalServerError:
		statusCode = http.StatusInternalServerError
	case *UnauthorizedError:
		statusCode = http.StatusUnauthorized
	}

	response := response.Response(response.ResponseParam{
		StatusCode: statusCode,
		Message:    err.Error(),
	})

	c.JSON(statusCode, response)
}
