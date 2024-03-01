package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type userRepositoryGorm struct {
	gorm *gorm.DB
}

func NewUserRepositoryGorm(gorm *gorm.DB) entity.UserRepository {
	return &userRepositoryGorm{gorm: gorm}
}

func (r userRepositoryGorm) GetAll() ([]entity.User, error) {
	var users []entity.User

	err := r.gorm.Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get users: %w", err)
	}

	return users, nil
}

func (r userRepositoryGorm) GetBySessionId(sessionId string) (*entity.User, error) {
	var user *entity.User

	err := r.gorm.Joins("JOIN session ON session.user_id = user_id").Where("session.id = ?", sessionId).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get user by session id: %w", err)
	}

	return user, nil
}

func (r userRepositoryGorm) GetById(id string) (*entity.User, error) {
	var user *entity.User

	err := r.gorm.Where("id = ?", id).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get user by id: %w", err)
	}

	return user, nil
}

func (r userRepositoryGorm) GetByEmail(email string) (*entity.User, error) {
	var user *entity.User

	err := r.gorm.Where("email = ?", email).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get user by email: %w", err)
	}

	return user, nil
}

func (r userRepositoryGorm) GetByParams(params *entity.User, limit int, offset int) ([]entity.User, error) {
	var users []entity.User

	err := r.gorm.Where(params).Limit(limit).Offset(offset).Find(&users).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get users by params: %w", err)
	}

	return users, nil
}

func (r userRepositoryGorm) Create(user *entity.User) error {
	err := r.gorm.Create(&user).Error
	if err != nil {
		return fmt.Errorf("cannot create user: %w", err)
	}

	return nil
}

func (r userRepositoryGorm) CreateMany(users []entity.User) error {
	err := r.gorm.Create(&users).Error
	fmt.Println(err)
	if err != nil {
		return fmt.Errorf("cannot create users: %w", err)
	}

	return nil
}

func (r userRepositoryGorm) Update(id string, user *entity.User) error {
	err := r.gorm.Model(&entity.User{}).Where("id = ?", id).Updates(user).Error
	if err != nil {
		return fmt.Errorf("cannot update user: %w", err)
	}

	return nil
}

func (r userRepositoryGorm) Delete(id string) error {
	err := r.gorm.Delete(&entity.User{Id: id}).Error

	if err != nil {
		return fmt.Errorf("cannot delete user: %w", err)
	}

	return nil
}
