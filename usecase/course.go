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

func (c courseUsecase) GetById(id string) (*entity.Course, error) {
	course, err := c.courseRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get course by id %s", id, err)
	}

	return course, nil
}

func (c courseUsecase) Create(name string, code string, semesterId string, lecturerId string) error {
	course := entity.Course{
		Id:         ulid.Make().String(),
		Name:       name,
		Code:       code,
		SemesterId: semesterId,
		LecturerId: lecturerId,
	}

	err := c.courseRepo.Create(&course)

	if err != nil {
		return errs.New(errs.ErrCreateCourse, "cannot create course", err)
	}

	return nil
}

func (u courseUsecase) Update(id string, course *entity.Course) error {
	existCourse, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get course id %s to update", id, err)
	} else if existCourse == nil {
		return errs.New(errs.ErrCourseNotFound, "cannot get course id %s to update", id)
	}

	err = u.courseRepo.Update(id, course)
	if err != nil {
		return errs.New(errs.ErrUpdateCourse, "cannot update course by id %s", course.Id, err)
	}

	return nil
}

func (c courseUsecase) Delete(id string) error {
	err := c.courseRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCourse, "cannot delete course", err)
	}

	return nil
}
