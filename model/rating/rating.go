package rating

import "time"

type Rating struct {
	ID        string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	UserID    int       `gorm:"not null"`
	CourseID  int       `gorm:"not null"`
	Rating    int       `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdatedTime"`
}

func (Rating) TableName() string {
	return "rating"
}
