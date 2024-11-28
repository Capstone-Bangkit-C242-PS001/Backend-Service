package user

import "time"

type User struct {
	ID         string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	Name       string    `gorm:"size:255;not null"`
	Email      string    `gorm:"size:255;unique;not null"`
	Password   string    `gorm:"size:255;not null"`
	ProfilePic *string   `gorm:"size:255"`
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdatedTime"`
}
