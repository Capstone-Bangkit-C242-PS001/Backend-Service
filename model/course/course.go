package course

import "time"

type Course struct {
	ID          string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Title       string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:255;not null"`
	Provider    string    `gorm:"size:255;not null"`
	ProviderURL *string   `gorm:"size:255;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdatedTime"`
}
