package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type lecturerUseCase struct {
	lecturerRepo entity.LecturerRepository
}

func NewLecturerUseCase(lecturerRepo entity.LecturerRepository) entity.LecturerUseCase {
	return &lecturerUseCase{lecturerRepo: lecturerRepo}
}

func (u lecturerUseCase) GetAll() ([]entity.Lecturer, error) {
	lecturers, err := u.lecturerRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get all lecturers", err)
	}

	return lecturers, nil
}

func (u lecturerUseCase) GetByEmail(email string) (*entity.Lecturer, error) {
	lecturer, err := u.lecturerRepo.GetByEmail(email)
	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get lecturer by email %s", email, err)
	}

	return lecturer, nil
}

func (u lecturerUseCase) GetByID(id string) (*entity.Lecturer, error) {
	lecturer, err := u.lecturerRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get lecturer by id %s", id, err)
	}

	return lecturer, nil
}

func (u lecturerUseCase) GetBySessionId(sessionId string) (*entity.Lecturer, error) {
	lecturer, err := u.lecturerRepo.GetBySessionId(sessionId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get lecturer by session id %s", sessionId, err)
	}

	return lecturer, nil
}

func (u lecturerUseCase) GetByParams(params *entity.Lecturer, limit int, offset int) ([]entity.Lecturer, error) {
	lecturers, err := u.lecturerRepo.GetByParams(params, limit, offset)

	if err != nil {
		return nil, errs.New(errs.ErrQueryLecturer, "cannot get lecturers by params", err)
	}

	return lecturers, nil
}

func (u lecturerUseCase) Create(name string, firstName string, lastName string) error {
	lecturer := &entity.Lecturer{
		ID:        ulid.Make().String(),
		Name:      name,
		FirstName: firstName,
		LastName:  lastName,
	}

	err := u.lecturerRepo.Create(lecturer)
	if err != nil {
		return errs.New(errs.ErrCreateLecturer, "cannot create lecturer", err)
	}

	return nil
}

func (u lecturerUseCase) Update(id string, lecturer *entity.Lecturer) error {
	existLecturer, err := u.GetByID(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get lecturer id %s to update", id, err)
	} else if existLecturer == nil {
		return errs.New(errs.ErrLecturerNotFound, "cannot get lecturer id %s to update", id)
	}

	err = u.lecturerRepo.Update(id, lecturer)
	if err != nil {
		return errs.New(errs.ErrUpdateLecturer, "cannot update lecturer by id %s", lecturer.ID, err)
	}

	return nil
}

func (u lecturerUseCase) Delete(id string) error {
	lecturer, err := u.GetByID(id)
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
