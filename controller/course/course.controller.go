package course

import (
	"net/http"
	"strconv"

	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/course"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/course"
	"github.com/gin-gonic/gin"
)

type courseController struct {
	service service.CourseService
}

func NewCourseController(service service.CourseService) *courseController {
	return &courseController{
		service: service,
	}
}

func (cc *courseController) Create(c *gin.Context) {
	var courseRequest dto.CreateRequest
	if err := c.ShouldBind(&courseRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid Input"})
		return
	}

	if err := cc.service.Create(&courseRequest); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Course added successfully " + courseRequest.Title,
	}))
}

func (cc *courseController) GetByID(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid user ID format"})
	}

	result, err := cc.service.GetByID(idNumber)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Request Succeed",
		Data:       result,
	}))
}

func (cc *courseController) GetAll(c *gin.Context) {
	result, err := cc.service.GetAll()
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Course Found",
		Data:       result,
	}))
}
