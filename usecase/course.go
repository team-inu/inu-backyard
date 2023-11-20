package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type courseUsecase struct {
	courseRepo entity.CourseRepository
}

func NewCourseUsecase(courseRepo entity.CourseRepository) entity.CourseUsecase {
	return &courseUsecase{courseRepo: courseRepo}
}

func (c courseUsecase) GetAll() ([]entity.Course, error) {
	courses, err := c.courseRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get all courses", err)
	}

	return courses, nil
}

func (c courseUsecase) GetByID(id string) (*entity.Course, error) {
	course, err := c.courseRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get course by id %s", id, err)
	}

	return course, nil
}

func (c courseUsecase) Create(name string, code string, year int, lecturerId string) (*entity.Course, error) {
	course := entity.Course{
		ID:         ulid.Make().String(),
		Name:       name,
		Code:       code,
		LecturerID: lecturerId,
	}

	err := c.courseRepo.Create(&course)

	if err != nil {
		return nil, errs.New(errs.ErrCreateCourse, "cannot create course", err)
	}

	return &course, nil
}

func (c courseUsecase) Delete(id string) error {
	err := c.courseRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCourse, "cannot delete course", err)
	}

	return nil
}
