package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	useRepo entity.UserRepository
}

func NewUserUseCase(userRepo entity.UserRepository) entity.UserUseCase {
	return &userUseCase{useRepo: userRepo}
}

func (u userUseCase) GetAll() ([]entity.User, error) {
	users, err := u.useRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryUser, "cannot get all users", err)
	}

	return users, nil
}

func (u userUseCase) GetByEmail(email string) (*entity.User, error) {
	user, err := u.useRepo.GetByEmail(email)
	if err != nil {
		return nil, errs.New(errs.ErrQueryUser, "cannot get user by email %s", email, err)
	}

	return user, nil
}

func (u userUseCase) GetById(id string) (*entity.User, error) {
	user, err := u.useRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryUser, "cannot get user by id %s", id, err)
	}

	return user, nil
}

func (u userUseCase) GetBySessionId(sessionId string) (*entity.User, error) {
	user, err := u.useRepo.GetBySessionId(sessionId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryUser, "cannot get user by session id %s", sessionId, err)
	}

	return user, nil
}

func (u userUseCase) GetByParams(params *entity.User, limit int, offset int) ([]entity.User, error) {
	users, err := u.useRepo.GetByParams(params, limit, offset)

	if err != nil {
		return nil, errs.New(errs.ErrQueryUser, "cannot get users by params", err)
	}

	return users, nil
}

func (u userUseCase) Create(firstName string, lastName string, email string, password string) error {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errs.New(errs.ErrCreateUser, "cannot create user", err)
	}

	hasPassword := string(bcryptPassword)

	user := &entity.User{
		Id:        ulid.Make().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hasPassword,
	}

	err = u.useRepo.Create(user)
	if err != nil {
		return errs.New(errs.ErrCreateUser, "cannot create user", err)
	}

	return nil
}

func (u userUseCase) CreateMany(users []entity.User) error {
	//encrypt password
	for i := range users {
		bcryptPassword, err := bcrypt.GenerateFromPassword([]byte((users)[i].Password), bcrypt.DefaultCost)
		if err != nil {
			return errs.New(errs.ErrCreateUser, "cannot create user", err)
		}
		users[i].Id = ulid.Make().String()
		(users)[i].Password = string(bcryptPassword)
	}

	err := u.useRepo.CreateMany(users)
	if err != nil {
		return errs.New(errs.ErrCreateUser, "cannot create user", err)
	}

	return nil
}

func (u userUseCase) Update(id string, user *entity.User) error {
	existedUser, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get user id %s to update", id, err)
	} else if existedUser == nil {
		return errs.New(errs.ErrUserNotFound, "cannot get user id %s to update", id)
	}

	err = u.useRepo.Update(id, user)
	if err != nil {
		return errs.New(errs.ErrUpdateUser, "cannot update user by id %s", user.Id, err)
	}

	return nil
}

func (u userUseCase) Delete(id string) error {
	user, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get user id %s to delete", id, err)
	} else if user == nil {
		return errs.New(errs.ErrUserNotFound, "cannot get user id %s to delete", id)
	}

	err = u.useRepo.Delete(id)

	if err != nil {
		return errs.New(errs.ErrDeleteUser, "cannot delete user by id %s", id, err)
	}

	return nil
}
