package course

import (
	"errors"

	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/errorhandler"
	model "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/model/course"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(course *model.Course) error
	GetByID(id int) (*model.Course, error)
	GetByIds(ids []int) (*[]model.Course, error)
	GetAll() ([]model.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) *courseRepository {
	return &courseRepository{
		db: db,
	}
}

func (cr *courseRepository) Create(course *model.Course) error {
	return cr.db.Create(course).Error
}

func (cr *courseRepository) GetByID(id int) (*model.Course, error) {
	var course model.Course

	result := cr.db.Where("id = ?", id).First(&course)
	if err := result.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &errorhandler.NotFoundError{Message: err.Error()}
		}
		return nil, &errorhandler.InternalServerError{Message: err.Error()}
	}

	return &course, nil
}

func (cr *courseRepository) GetByIds(ids []int) (*[]model.Course, error) {
	var courses []model.Course
	result := cr.db.Where("id IN ?", ids).Find(&courses)
	if err := result.Error; err != nil {
		return nil, err
	}

	return &courses, nil
}

func (cr *courseRepository) GetAll() ([]model.Course, error) {
	var courses []model.Course
	result := cr.db.Find(&courses)

	return courses, result.Error
}
