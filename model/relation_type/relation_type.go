package relation_type

import "time"

type RelationType struct {
	ID           string    `gorm:"type:char(36);primaryKey;default:(UUID())"`
	RelationName string    `gorm:"size:255;not null"`
	Description  string    `gorm:"size:255;not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdatedTime"`
}
