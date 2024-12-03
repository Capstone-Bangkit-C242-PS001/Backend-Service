package prediction

type PredictRequest struct {
	Skillsets []string `json:"skillsets"`
}

type PredictResponse struct {
	ID          int    `gorm:"primaryKey;autoIncrement"`
	Title       string `gorm:"size:255;not null"`
	Description string `gorm:"size:255;not null"`
	Provider    string `gorm:"size:255;not null"`
	ProviderURL string `gorm:"size:255;not null"`
}
