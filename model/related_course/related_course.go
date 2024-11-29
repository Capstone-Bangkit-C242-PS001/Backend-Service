package related_course

import "time"

type RelatedCourse struct {
	ID              string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	CourseID        string    `gorm:"type:char(36);not null"`
	RelatedCourseID string    `gorm:"type:char(36);not null"`
	RelationTypeID  string    `gorm:"type:char(36);not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdatedTime"`
}
