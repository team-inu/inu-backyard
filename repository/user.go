package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type lecturerRepositoryGorm struct {
	gorm *gorm.DB
}

func NewLecturerRepositoryGorm(gorm *gorm.DB) entity.UserRepository {
	return &lecturerRepositoryGorm{gorm: gorm}
}

func (r lecturerRepositoryGorm) GetAll() ([]entity.User, error) {
	var lecturers []entity.User

	err := r.gorm.Find(&lecturers).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get lecturers: %w", err)
	}

	return lecturers, nil
}

func (r lecturerRepositoryGorm) GetBySessionId(sessionId string) (*entity.User, error) {
	var lecturer *entity.User

	err := r.gorm.Joins("JOIN session ON session.lecturer_id = lecturer_id").Where("session.id = ?", sessionId).First(&lecturer).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get lecturer by session id: %w", err)
	}

	return lecturer, nil
}

func (r lecturerRepositoryGorm) GetById(id string) (*entity.User, error) {
	var lecturer *entity.User

	err := r.gorm.Where("id = ?", id).First(&lecturer).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get lecturer by id: %w", err)
	}

	return lecturer, nil
}

func (r lecturerRepositoryGorm) GetByEmail(email string) (*entity.User, error) {
	var lecturer *entity.User

	err := r.gorm.Where("email = ?", email).First(&lecturer).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get lecturer by email: %w", err)
	}

	return lecturer, nil
}

func (r lecturerRepositoryGorm) GetByParams(params *entity.User, limit int, offset int) ([]entity.User, error) {
	var lecturers []entity.User

	err := r.gorm.Where(params).Limit(limit).Offset(offset).Find(&lecturers).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get lecturers by params: %w", err)
	}

	return lecturers, nil
}

func (r lecturerRepositoryGorm) Create(lecturer *entity.User) error {
	err := r.gorm.Create(&lecturer).Error
	if err != nil {
		return fmt.Errorf("cannot create lecturer: %w", err)
	}

	return nil
}

func (r lecturerRepositoryGorm) CreateMany(lecturers []entity.User) error {
	err := r.gorm.Create(&lecturers).Error
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("cannot create lecturers: %w", err)
	}

	return nil
}

func (r lecturerRepositoryGorm) Update(id string, lecturer *entity.User) error {
	err := r.gorm.Model(&entity.User{}).Where("id = ?", id).Updates(lecturer).Error
	if err != nil {
		return fmt.Errorf("cannot update lecturer: %w", err)
	}

	return nil
}

func (r lecturerRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.User{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete lecturer: %w", err)
	}

	return nil
}
