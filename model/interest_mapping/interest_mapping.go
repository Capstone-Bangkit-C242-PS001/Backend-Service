package interest_mapping

import "time"

type InterestMapping struct {
	ID         string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	UserID     string    `gorm:"type:char(36);not null"`
	InterestID string    `gorm:"type:char(36);not null"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdatedTime"`
}

func (InterestMapping) TableName() string {
	return "interest_mapping"
}
