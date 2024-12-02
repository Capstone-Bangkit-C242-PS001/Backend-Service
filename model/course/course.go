package course

import "time"

type Course struct {
	ID          int       `gorm:"primaryKey;autoIncrement"`
	Title       string    `gorm:"size:255;not null"`
	Description string    `gorm:"size:255;not null"`
	Provider    string    `gorm:"size:255;not null"`
	ProviderURL *string   `gorm:"size:255;not null"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdatedTime"`
}
