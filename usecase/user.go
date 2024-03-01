package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"golang.org/x/crypto/bcrypt"
)

type lecturerUseCase struct {
	lecturerRepo entity.UserRepository
}

func NewLecturerUseCase(lecturerRepo entity.UserRepository) entity.UserUseCase {
	return &lecturerUseCase{lecturerRepo: lecturerRepo}
}

func (u lecturerUseCase) GetAll() ([]entity.User, error) {
	lecturers, err := u.lecturerRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get all lecturers", err)
	}

	return lecturers, nil
}

func (u lecturerUseCase) GetByEmail(email string) (*entity.User, error) {
	lecturer, err := u.lecturerRepo.GetByEmail(email)
	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get lecturer by email %s", email, err)
	}

	return lecturer, nil
}

func (u lecturerUseCase) GetById(id string) (*entity.User, error) {
	lecturer, err := u.lecturerRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get lecturer by id %s", id, err)
	}

	return lecturer, nil
}

func (u lecturerUseCase) GetBySessionId(sessionId string) (*entity.User, error) {
	lecturer, err := u.lecturerRepo.GetBySessionId(sessionId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get lecturer by session id %s", sessionId, err)
	}

	return lecturer, nil
}

func (u lecturerUseCase) GetByParams(params *entity.User, limit int, offset int) ([]entity.User, error) {
	lecturers, err := u.lecturerRepo.GetByParams(params, limit, offset)

	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get lecturers by params", err)
	}

	return lecturers, nil
}

func (u lecturerUseCase) Create(firstName string, lastName string, email string, password string) error {
	bcryptPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errs.New(errs.ErrCreateLecturer, "cannot create lecturer", err)
	}

	hasPassword := string(bcryptPassword)

	lecturer := &entity.User{
		Id:        ulid.Make().String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  hasPassword,
	}

	err = u.lecturerRepo.Create(lecturer)
	if err != nil {
		return errs.New(errs.ErrCreateLecturer, "cannot create lecturer", err)
	}

	return nil
}

func (u lecturerUseCase) CreateMany(lecturers []entity.User) error {
	//encrypt password
	for i := range lecturers {
		bcryptPassword, err := bcrypt.GenerateFromPassword([]byte((lecturers)[i].Password), bcrypt.DefaultCost)
		if err != nil {
			return errs.New(errs.ErrCreateLecturer, "cannot create lecturer", err)
		}
		lecturers[i].Id = ulid.Make().String()
		(lecturers)[i].Password = string(bcryptPassword)
	}

	err := u.lecturerRepo.CreateMany(lecturers)
	if err != nil {
		return errs.New(errs.ErrCreateLecturer, "cannot create lecturer", err)
	}

	return nil
}

func (u lecturerUseCase) Update(id string, lecturer *entity.User) error {
	existLecturer, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get lecturer id %s to update", id, err)
	} else if existLecturer == nil {
		return errs.New(errs.ErrLecturerNotFound, "cannot get lecturer id %s to update", id)
	}

	err = u.lecturerRepo.Update(id, lecturer)
	if err != nil {
		return errs.New(errs.ErrUpdateLecturer, "cannot update lecturer by id %s", lecturer.Id, err)
	}

	return nil
}

func (u lecturerUseCase) Delete(id string) error {
	lecturer, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get lecturer id %s to delete", id, err)
	} else if lecturer == nil {
		return errs.New(errs.ErrLecturerNotFound, "cannot get lecturer id %s to delete", id)
	}

	err = u.lecturerRepo.Delete(id)

	if err != nil {
		return errs.New(errs.ErrDeleteLecturer, "cannot delete lecturer by id %s", id, err)
	}

	return nil
}
