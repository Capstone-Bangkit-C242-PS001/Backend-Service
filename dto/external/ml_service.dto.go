package external

type PredictRequest struct {
	UserID   int      `json:"user_id"`
	Skillset []string `json:"skillset"`
}

type PredictResponse []int
