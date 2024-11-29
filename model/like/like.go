package like

import "time"

type Like struct {
	ID        string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	UserID    string    `gorm:"type:char(36);not null"`
	CourseID  string    `gorm:"type:char(36);not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdatedTime"`
}
