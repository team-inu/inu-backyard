package usecase

import (
	"github.com/gofiber/fiber/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	sessionUseCase  entity.SessionUseCase
	lecturerUseCase entity.UserUseCase
}

func NewAuthUseCase(
	sessionUseCase entity.SessionUseCase,
	lecturerUseCase entity.UserUseCase,
) entity.AuthUseCase {
	return &authUseCase{
		sessionUseCase:  sessionUseCase,
		lecturerUseCase: lecturerUseCase,
	}
}

func (u authUseCase) Authenticate(header string) (*entity.User, error) {
	session, err := u.sessionUseCase.Validate(header)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot authenticate user", err)
	}

	user, err := u.lecturerUseCase.GetBySessionId(session.Id)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get user to authenticate", err)
	}
	return user, nil
}

func (u authUseCase) SignIn(email string, password string, ipAddress string, userAgent string) (*fiber.Cookie, error) {

	lecturer, err := u.lecturerUseCase.GetByEmail(email)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get user data to sign in", err)
	} else if lecturer == nil {
		return nil, errs.New(errs.ErrLecturerNotFound, "account with email %s is not registered", email)
	}

	if err = bcrypt.CompareHashAndPassword([]byte(lecturer.Password), []byte(password)); err != nil {
		return nil, errs.New(errs.ErrLecturerPassword, "password is incorrect", err)
	}

	cookie, err := u.sessionUseCase.Create(lecturer.Id, ipAddress, userAgent)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot create session to sign in", err)
	}
	return cookie, nil
}

func (u authUseCase) SignOut(header string) (*fiber.Cookie, error) {
	session, err := u.sessionUseCase.Validate(header)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot validate session to sign out", err)
	}

	cookie, err := u.sessionUseCase.Destroy(session.Id)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot destroy session to sign out", err)
	}
	return cookie, nil
}
