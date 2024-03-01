package repository

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	"gorm.io/gorm"
)

type sessionRepository struct {
	gorm *gorm.DB
}

func NewSessionRepository(gorm *gorm.DB) entity.SessionRepository {
	return &sessionRepository{gorm: gorm}
}

func (r *sessionRepository) Create(session *entity.Session) error {
	fmt.Println(session.UserId)

	err := r.gorm.Create(session).Error

	if err != nil {
		return fmt.Errorf("cannot query to create session: %w", err)
	}
	return nil
}

func (r *sessionRepository) Get(id string) (*entity.Session, error) {
	var session entity.Session
	err := r.gorm.Where("id = ?", id).First(&session).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("cannot query to get session: %w", err)
	}

	return &session, nil
}

func (r *sessionRepository) Delete(id string) error {
	err := r.gorm.Delete(&entity.Session{Id: id}).Error
	if err != nil {
		return fmt.Errorf("cannot query to delete session: %w", err)
	}
	return nil
}

func (r *sessionRepository) DeleteByUserId(userId string) error {
	err := r.gorm.Delete(&entity.Session{UserId: userId}).Error
	if err != nil {
		return fmt.Errorf("cannot query to delete session: %w", err)
	}
	return nil
}

func (r *sessionRepository) DeleteDuplicates(userId string, ipAddress string, userAgent string) error {
	result := r.gorm.Where("user_id = ? AND ip_address = ? AND user_agent = ?", userId, ipAddress, userAgent).Delete(&entity.Session{})

	if result.Error != nil {
		return fmt.Errorf("cannot query to delete session: %w", result.Error)
	}

	// if result.RowsAffected == 0 {
	// 	return fmt.Errorf("no session to delete")
	// }

	return nil
}
