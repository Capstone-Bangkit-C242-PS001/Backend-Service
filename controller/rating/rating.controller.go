package rating

import (
	"net/http"
	"strconv"

	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/rating"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/response"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	service "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/service/rating"
	"github.com/gin-gonic/gin"
)

type ratingController struct {
	service service.RatingService
}

func NewRatingController(service service.RatingService) *ratingController {
	return &ratingController{
		service: service,
	}
}

// @Summary Add a new rating
// @Description Allows a user to add a rating for a specific course
// @Tags Ratings
// @Accept json
// @Produce json
// @Param userID header int true "User ID from JWT token"
// @Param ratingRequest body CreateRatingRequest true "Rating request payload"
// @Success 200 "Rating added successfully"
// @Failure 400 "Invalid input"
// @Failure 401 "Unauthorized"
// @Failure 500 "Internal server error"
// @Router /ratings [post]
func (rc *ratingController) Create(c *gin.Context) {
	var ratingRequest dto.CreateRatingRequest
	if err := c.ShouldBind(&ratingRequest); err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid Input"})
		return
	}

	userID := c.Value("userID").(int)

	if err := rc.service.Create(userID, &ratingRequest); err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Rating added successfull",
	}))
}

func (rc *ratingController) GetByUserID(c *gin.Context) {
	id := c.Param("id")
	idNumber, err := strconv.Atoi(id)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid user ID format"})
	}

	result, err := rc.service.GetByUserID(idNumber)
	if err != nil {
		errorhandler.HandleError(c, err)
		return
	}

	c.JSON(http.StatusOK, response.Response(response.ResponseParam{
		StatusCode: http.StatusOK,
		Message:    "Get rating successfully with user id " + id,
		Data:       result,
	}))
}

func (rc *ratingController) GetRating(c *gin.Context) {
	userID := c.Query("user_id")
	courseID := c.Query("course_id")

	if userID == "" || courseID == "" {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "user_id and course_id are required query parameters"})
		return
	}

	userIDNumber, err := strconv.Atoi(userID)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid user ID format"})
	}

	courseIDNumber, err := strconv.Atoi(courseID)
	if err != nil {
		errorhandler.HandleError(c, &errorhandler.BadRequestError{Message: "Invalid course ID format"})
	}

	result, err := rc.service.GetRating(userIDNumber, courseIDNumber)
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
