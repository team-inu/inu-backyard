package usecase

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"github.com/team-inu/inu-backyard/internal/config"
)

type sessionUsecase struct {
	sessionRepository entity.SessionRepository
	config            config.AuthConfig
}

func NewSessionUsecase(
	sessionRepository entity.SessionRepository,
	config config.AuthConfig,
) entity.SessionUseCase {
	return &sessionUsecase{
		sessionRepository: sessionRepository,
		config:            config,
	}
}

func (u sessionUsecase) Sign(id string) string {
	hmac := hmac.New(sha256.New, []byte(u.config.Session.Secret))
	hmac.Write([]byte(id))
	regex := regexp.MustCompile(`=+$`)
	signature := regex.ReplaceAllString(base64.StdEncoding.EncodeToString(hmac.Sum(nil)), "")
	return u.config.Session.Prefix + ":" + id + "." + signature
}

func (u sessionUsecase) Unsign(header string) (string, error) {
	if !strings.HasPrefix(header, u.config.Session.Prefix+":") {
		return "", errs.New(errs.ErrSessionPrefix, "prefix mismatch")
	}

	id := header[len(u.config.Session.Prefix)+1 : strings.LastIndex(header, ".")]
	expectation := u.Sign(id)

	isLengthMatch := len([]byte(header)) == len([]byte(expectation))
	isInputMatch := subtle.ConstantTimeCompare([]byte(header), []byte(expectation)) == 1

	if !isLengthMatch || !isInputMatch {
		return "", errs.New(errs.ErrSignatureMismatch, "signature mismatch")
	}
	return id, nil
}

func (u sessionUsecase) Create(
	userId string, ipAddress string, userAgent string,
) (*fiber.Cookie, error) {
	fmt.Println("=====xxxxx")
	if err := u.sessionRepository.DeleteDuplicates(userId, ipAddress, userAgent); err != nil {
		return nil, errs.New(
			errs.ErrDupSession,
			"cannot delete previous session to create a new session for user id %d", userId,
			err,
		)
	}
	fmt.Println("=====xxxxxsss")

	id := uuid.NewString()
	signedId := u.Sign(id)
	createdAt := time.Now()
	expiredAt := createdAt.Add(time.Duration(u.config.Session.MaxAge) * time.Second)

	err := u.sessionRepository.Create(&entity.Session{
		Id:        id,
		UserId:    userId,
		IpAddress: ipAddress,
		UserAgent: userAgent,
		ExpiredAt: expiredAt,
		CreatedAt: createdAt,
	})
	if err != nil {
		return nil, errs.New(errs.ErrCreateSession, "cannot create session for user id %d", userId, err)
	}

	cookie := &fiber.Cookie{
		Name:     u.config.Session.CookieName,
		Value:    signedId,
		HTTPOnly: true,
		Expires:  expiredAt,
	}
	return cookie, nil
}

func (u sessionUsecase) Get(header string) (*entity.Session, error) {
	id, err := u.Unsign(header)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot unsign session", err)
	}

	session, err := u.sessionRepository.Get(id)
	if err != nil {
		return nil, errs.New(errs.ErrGetSession, "cannot get session from header", err)
	}
	return session, nil
}

func (u sessionUsecase) Destroy(id string) (*fiber.Cookie, error) {
	return &fiber.Cookie{
		Name:     u.config.Session.CookieName,
		HTTPOnly: true,
		Expires:  time.Unix(0, 0),
	}, u.sessionRepository.Delete(id)
}

func (u sessionUsecase) DestroyByUserId(userId string) (*fiber.Cookie, error) {
	return &fiber.Cookie{
		Name:     u.config.Session.CookieName,
		HTTPOnly: true,
		Expires:  time.Unix(0, 0),
	}, u.sessionRepository.DeleteByUserId(userId)
}

func (u sessionUsecase) Validate(header string) (*entity.Session, error) {
	session, err := u.Get(header)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot validate session", err)
	} else if session == nil {
		return nil, errs.New(errs.ErrInvalidSession, "session is invalid")
	}

	if !time.Now().Before(session.ExpiredAt) {
		return nil, errs.New(errs.ErrSessionExpired, "session expired")
	}

	return session, nil
}
