package user_interest

import "time"

type UserInterestRequest struct {
	InterestName string `json:"interest_name" binding:"required"`
	Description  string `json:"description"`
}

type UserInterestResponse struct {
	ID           string    `json:"id"`
	InterestName string    `json:"interest_name"`
	Description  string    `json:"description"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
