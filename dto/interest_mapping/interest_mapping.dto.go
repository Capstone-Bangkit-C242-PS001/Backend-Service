package interest_mapping

import dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/user_interest"

type InterestMappingRequest struct {
	InterestIDs []string `json:"interest_ids" binding:"required"`
}

type InterestMappingResponse struct {
	UserID    string                     `json:"user_id"`
	Interests []dto.UserInterestResponse `json:"interests"`
}
