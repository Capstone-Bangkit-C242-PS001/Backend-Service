package prediction

type PredictRequest struct {
	Skillsets []string `json:"skillsets"`
}

type PredictResponse struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Provider    string `json:"provider"`
	ProviderURL string `json:"provider_url"`
}
