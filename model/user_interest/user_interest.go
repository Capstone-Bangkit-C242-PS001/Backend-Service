package user_interest

import "time"

type UserInterest struct {
	ID           string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	InterestName string    `gorm:"size:255;not null"`
	Description  string    `gorm:"size:255;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func (UserInterest) TableName() string {
	return "user_interest"
}
