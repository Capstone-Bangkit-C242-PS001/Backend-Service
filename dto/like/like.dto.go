package like

type LikeRequest struct {
	CourseID string `json:"course_id"`
}

type UnlikeRequest struct {
	CourseID string `json:"course_id"`
}
