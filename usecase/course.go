package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type courseUseCase struct {
	courseRepo      entity.CourseRepository
	semesterUseCase entity.SemesterUseCase
	userUseCase     entity.UserUseCase
}

func NewCourseUseCase(courseRepo entity.CourseRepository, semesterUseCase entity.SemesterUseCase, userUseCase entity.UserUseCase) entity.CourseUseCase {
	return &courseUseCase{courseRepo: courseRepo, semesterUseCase: semesterUseCase, userUseCase: userUseCase}
}

func (u courseUseCase) GetAll() ([]entity.Course, error) {
	courses, err := u.courseRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get all courses", err)
	}

	return courses, nil
}

func (u courseUseCase) GetById(id string) (*entity.Course, error) {
	course, err := u.courseRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryStudent, "cannot get course by id %s", id, err)
	}

	return course, nil
}

func (u courseUseCase) GetByUserId(userId string) ([]entity.Course, error) {
	user, err := u.userUseCase.GetById(userId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get user id %s while get scores", user, err)
	} else if user == nil {
		return nil, errs.New(errs.ErrQueryCourse, "user id %s not found while getting scores", userId, err)
	}

	course, err := u.courseRepo.GetByUserId(userId)
	if err != nil {
		return nil, errs.New(errs.ErrQueryCourse, "cannot get score by user id %s", userId, err)
	}

	return course, nil
}

func (u courseUseCase) Create(user entity.User, semesterId string, userId string, name string, code string, curriculum string, description string, expectedPassingCloPercentage float64, criteriaGrade entity.CriteriaGrade) error {
	if !user.IsRoles([]entity.UserRole{entity.UserRoleHeadOfCurriculum}) {
		return errs.New(errs.ErrCreateCourse, "no permission to create course")
	}

	semester, err := u.semesterUseCase.GetById(semesterId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get semester id %s while creating course", semesterId, err)
	} else if semester == nil {
		return errs.New(errs.ErrSemesterNotFound, "semester id %s not found while creating course", semesterId)
	}

	lecturer, err := u.userUseCase.GetById(userId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get user id %s while creating course", userId, err)
	} else if lecturer == nil {
		return errs.New(errs.ErrUserNotFound, "user id %s not found while creating course", userId)
	}

	if !criteriaGrade.IsValid() {
		return errs.New(errs.ErrCreateCourse, "invalid criteria grade")
	}

	course := entity.Course{
		Id:            ulid.Make().String(),
		SemesterId:    semesterId,
		UserId:        userId,
		Name:          name,
		Code:          code,
		Curriculum:    curriculum,
		Description:   description,
		CriteriaGrade: criteriaGrade,
	}

	err = u.courseRepo.Create(&course)
	if err != nil {
		return errs.New(errs.ErrCreateCourse, "cannot create course", err)
	}

	return nil
}

func (u courseUseCase) Update(user entity.User, id string, name string, code string, curriculum string, description string, expectedPassingCloPercentage float64, criteriaGrade entity.CriteriaGrade) error {
	existCourse, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get course id %s to update", id, err)
	} else if existCourse == nil {
		return errs.New(errs.ErrCourseNotFound, "cannot get course id %s to update", id)
	}

	if !user.IsRoles([]entity.UserRole{entity.UserRoleHeadOfCurriculum}) && user.Id != existCourse.UserId {
		return errs.New(errs.ErrCreateCourse, "No permission to edit this course")
	}

	if !criteriaGrade.IsValid() {
		return errs.New(errs.ErrCreateCourse, "invalid criteria grade")
	}

	err = u.courseRepo.Update(id, &entity.Course{
		Name:                         name,
		Code:                         code,
		Curriculum:                   curriculum,
		Description:                  description,
		CriteriaGrade:                criteriaGrade,
		ExpectedPassingCloPercentage: expectedPassingCloPercentage,
	})
	if err != nil {
		return errs.New(errs.ErrUpdateCourse, "cannot update course by id %s", id, err)
	}

	return nil
}

func (u courseUseCase) Delete(user entity.User, id string) error {
	if !user.IsRoles([]entity.UserRole{entity.UserRoleHeadOfCurriculum}) {
		return errs.New(errs.ErrCreateCourse, "no permission to create course")
	}

	err := u.courseRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCourse, "cannot delete course", err)
	}

	return nil
}
