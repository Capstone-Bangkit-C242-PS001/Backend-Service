package like

import (
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/like"
	"gorm.io/gorm"
)

type LikeRepository interface {
	Create(like model.Like) error
	Get(userID string) (*[]model.Like, error)
	Delete(ID string) error
}

type likeRepository struct {
	db *gorm.DB
}

func NewLikeRepository(db *gorm.DB) *likeRepository {
	return &likeRepository{db: db}
}

func (lr *likeRepository) Create(like model.Like) error {
	if err := lr.db.Create(&like).Error; err != nil {
		return err
	}
	return nil
}

func (lr *likeRepository) Get(userID string) (*[]model.Like, error) {
	var likes []model.Like
	if err := lr.db.Where("user_id = ?", userID).Find(&likes).Error; err != nil {
		return nil, err
	}
	return &likes, nil
}

func (lr *likeRepository) Delete(ID string) error {
	if err := lr.db.Where("id = ?", ID).Delete(&model.Like{}).Error; err != nil {
		return err
	}
	return nil
}
