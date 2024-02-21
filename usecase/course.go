package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type courseUsecase struct {
	courseRepo      entity.CourseRepository
	semesterUseCase entity.SemesterUseCase
	lecturerUseCase entity.LecturerUseCase
}

func NewCourseUsecase(courseRepo entity.CourseRepository, semesterUseCase entity.SemesterUseCase, lecturerUseCase entity.LecturerUseCase) entity.CourseUsecase {
	return &courseUsecase{courseRepo: courseRepo, semesterUseCase: semesterUseCase, lecturerUseCase: lecturerUseCase}
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

func (c courseUsecase) Create(semesterId string, lecturerId string, name string, code string, curriculum string, description string, criteriaGrade entity.CriteriaGrade) error {
	semester, err := c.semesterUseCase.GetById(semesterId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get semester id %s while creating course", semesterId, err)
	} else if semester == nil {
		return errs.New(errs.ErrSemesterNotFound, "semester id %s not found while creating course", semesterId)
	}

	lecturer, err := c.lecturerUseCase.GetById(lecturerId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get lecturer id %s while creating course", lecturerId, err)
	} else if lecturer == nil {
		return errs.New(errs.ErrLecturerNotFound, "lecturer id %s not found while creating course", lecturerId)
	}

	if !criteriaGrade.IsValid() {
		return errs.New(errs.ErrCreateCourse, "invalid criteria grade")
	}

	course := entity.Course{
		Id:            ulid.Make().String(),
		SemesterId:    semesterId,
		LecturerId:    lecturerId,
		Name:          name,
		Code:          code,
		Curriculum:    curriculum,
		Description:   description,
		CriteriaGrade: criteriaGrade,
	}

	err = c.courseRepo.Create(&course)
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
