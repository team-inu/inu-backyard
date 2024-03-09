package usecase

import (
	"fmt"

	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type coursePortfolioUseCase struct {
	CourseUseCase     entity.CourseUseCase
	UserUseCase       entity.UserUseCase
	EnrollmentUseCase entity.EnrollmentUseCase
}

func NewCoursePortfolioUseCase(
	courseUseCase entity.CourseUseCase,
	userUseCase entity.UserUseCase,
	enrollmentUseCase entity.EnrollmentUseCase,
) entity.CoursePortfolioUseCase {
	return &coursePortfolioUseCase{
		CourseUseCase:     courseUseCase,
		UserUseCase:       userUseCase,
		EnrollmentUseCase: enrollmentUseCase,
	}
}

func (u coursePortfolioUseCase) Generate(courseId string) (*entity.CoursePortfolio, error) {
	course, err := u.CourseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course id %s while generate course portfolio", courseId, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "course id %s not found while generate course portfolio", courseId)
	}

	lecturer, err := u.UserUseCase.GetById(course.UserId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get lecturer id %s while generate course portfolio", course.UserId, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "user id %s not found while generate course portfolio", course.UserId)
	}

	courseInfo := entity.CourseInfo{
		Name:      course.Name,
		Code:      course.Code,
		Lecturers: []string{fmt.Sprintf("%s %s", lecturer.FirstName, lecturer.LastName)},
	}

	gradeDistribution, err := u.CalculateGradeDistribution()
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot calculate grade distribution while generate course portfolio", err)
	}

	tabeeOutcomes, err := u.EvaluateTabeeOutcomes()
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate tabee outcomes while generate course portfolio", err)
	}

	courseResult := entity.CourseResult{
		GradeDistribution: *gradeDistribution,
		TabeeOutcomes:     tabeeOutcomes,
	}

	coursePortfolio := &entity.CoursePortfolio{
		CourseInfo:   courseInfo,
		CourseResult: courseResult,
	}

	return coursePortfolio, nil
}

// TODO: implement
func (u coursePortfolioUseCase) CalculateGradeDistribution() (*entity.GradeDistribution, error) {
	return &entity.GradeDistribution{}, nil
}

// TODO: implement
func (u coursePortfolioUseCase) EvaluateTabeeOutcomes() ([]entity.TabeeOutcome, error) {
	return nil, nil
}
