package usecase

import (
	"fmt"

	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"github.com/team-inu/inu-backyard/repository"
)

type ImporterUseCase struct {
	importerRepository            repository.ImporterRepositoryGorm
	courseUseCase                 entity.CourseUseCase
	enrollmentUseCase             entity.EnrollmentUseCase
	assignmentUseCase             entity.AssignmentUseCase
	programOutcomeUseCase         entity.ProgramOutcomeUseCase
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase
	courseLearningOutcomeUseCase  entity.CourseLearningOutcomeUseCase
}

func NewImporterUseCase(
	importerRepository repository.ImporterRepositoryGorm,
	courseUseCase entity.CourseUseCase,
	enrollmentUseCase entity.EnrollmentUseCase,
	assignmentUseCase entity.AssignmentUseCase,
	programOutcomeUseCase entity.ProgramOutcomeUseCase,
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase,
	courseLearningOutcomeUseCase entity.CourseLearningOutcomeUseCase,
) ImporterUseCase {
	return ImporterUseCase{
		importerRepository:            importerRepository,
		courseUseCase:                 courseUseCase,
		enrollmentUseCase:             enrollmentUseCase,
		assignmentUseCase:             assignmentUseCase,
		programOutcomeUseCase:         programOutcomeUseCase,
		programLearningOutcomeUseCase: programLearningOutcomeUseCase,
		courseLearningOutcomeUseCase:  courseLearningOutcomeUseCase,
	}
}

// assignment group
type score struct {
	Score     *float64 `json:"score" validate:"required"`
	StudentId string   `json:"studentId" validate:"required"`
}

type assignment struct {
	Name                             string   `json:"name" validate:"required"`
	Description                      string   `json:"description" validate:"required"`
	MaxScore                         int      `json:"maxScore" validate:"required"`
	ExpectedScorePercentage          float64  `json:"expectedScorePercentage" validate:"required"`
	ExpectedPassingStudentPercentage float64  `json:"expectedPassingStudentPercentage" validate:"required"`
	CourseLearningOutcomeCodes       []string `json:"courseLearningOutcomeCodes" validate:"required,dive"`
	IsIncludedInClo                  *bool    `json:"isIncludedInClo" validate:"required"`
	Scores                           []score  `json:"scores" validate:"required,dive"`
}

type ImportAssignmentGroup struct {
	Name        string       `json:"name" validate:"required"`
	Weight      int          `json:"weight" validate:"required"`
	Assignments []assignment `json:"assignments" validate:"dive"`
}

// course learning outcome
type ImportCourseLearningOutcome struct {
	Code                                string   `json:"code" validate:"required"`
	Description                         string   `json:"description" validate:"required"`
	ExpectedPassingAssignmentPercentage float64  `json:"expectedPassingAssignmentPercentage" validate:"required"`
	ExpectedPassingStudentPercentage    float64  `json:"expectedPassingStudentPercentage" validate:"required"`
	Status                              string   `json:"status" validate:"required"`
	SubProgramLearningOutcomeCodes      []string `json:"subProgramLearningOutcomeCodes" validate:"required,dive"`
	ProgramOutcomeCode                  string   `json:"programOutcomeCode" validate:"required"`
}

func (u ImporterUseCase) UpdateOrCreate(
	courseId string,
	programYear string,
	lecturerId string,
	studentIds []string,
	clos []ImportCourseLearningOutcome,
	assignmentGroups []ImportAssignmentGroup,
) error {
	course, err := u.courseUseCase.GetById(courseId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get course id while import course", err)
	} else if course == nil {
		return errs.New(errs.SameCode, "course id %s not found while importing", err, courseId)
	}

	// prepare old data to delete
	oldAssignments, err := u.assignmentUseCase.GetByCourseId(courseId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get old assignment while import course")
	}

	OldAssignmentIds := make([]string, 0)
	for _, assignment := range oldAssignments {
		OldAssignmentIds = append(OldAssignmentIds, assignment.Id)
	}

	OldAssignmentGroups, err := u.assignmentUseCase.GetGroupByCourseId(courseId, false)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get old assignment group while import course")
	}

	oldAssignmentGroupIds := make([]string, 0)
	for _, assignmentGroup := range OldAssignmentGroups {
		oldAssignmentGroupIds = append(oldAssignmentGroupIds, assignmentGroup.Id)
	}

	oldClos, err := u.courseLearningOutcomeUseCase.GetByCourseId(courseId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get old clo while import course")
	}

	oldCloIds := make([]string, 0)
	for _, clo := range oldClos {
		oldCloIds = append(oldCloIds, clo.Id)
	}

	// beginning to prepare
	courseLearningOutcomes := make([]entity.CourseLearningOutcome, 0, len(clos))
	groupsToCreate := make([]entity.AssignmentGroup, 0)
	assignmentsToCreate := make([]entity.Assignment, 0)
	enrollmentsToCreate := make([]entity.Enrollment, 0)
	scoresToCreate := make([]entity.Score, 0)

	// prepare new course learning outcomes
	courseLearningOutcomeByCode := make(map[string]string, 0)
	for _, clo := range clos {
		programOutcome, err := u.programOutcomeUseCase.GetByCode(clo.ProgramOutcomeCode)
		if err != nil {
			return errs.New(errs.SameCode, "cannot get program outcome id while import course", err)
		} else if programOutcome == nil {
			return errs.New(errs.ErrProgrammeNotFound, "program outcome code %s in clo not found while import course", clo.ProgramOutcomeCode)
		}

		subPlos := []*entity.SubProgramLearningOutcome{}
		fmt.Println(clo.SubProgramLearningOutcomeCodes)
		for _, subPloCode := range clo.SubProgramLearningOutcomeCodes {
			subPlo, err := u.programLearningOutcomeUseCase.GetSubPloByCode(subPloCode, course.Curriculum, programYear)
			if err != nil {
				return errs.New(errs.ErrSubPLONotFound, "cannot get sub plo id %s while import course", clo.ProgramOutcomeCode, subPloCode)

			} else if subPlo == nil {
				return errs.New(errs.ErrSubPLONotFound, "sub program learning outcome code %s curriculum: %s year: %s not found while import course", clo.ProgramOutcomeCode, course.Curriculum, programYear)
			}

			subPlos = append(subPlos, &entity.SubProgramLearningOutcome{
				Id: subPlo.Id,
			})
		}

		id := ulid.Make().String()

		courseLearningOutcomeByCode[clo.Code] = id
		courseLearningOutcomes = append(courseLearningOutcomes, entity.CourseLearningOutcome{
			Id:                                  id,
			Code:                                clo.Code,
			Description:                         clo.Description,
			Status:                              clo.Status,
			ExpectedPassingAssignmentPercentage: clo.ExpectedPassingAssignmentPercentage,
			ExpectedPassingStudentPercentage:    clo.ExpectedPassingStudentPercentage,
			CourseId:                            courseId,
			ProgramOutcome:                      *programOutcome,
			SubProgramLearningOutcomes:          subPlos,
		})
	}

	// prepare new assignment groups
	for _, assignmentGroup := range assignmentGroups {
		assignmentGroupId := ulid.Make().String()
		groupsToCreate = append(groupsToCreate, entity.AssignmentGroup{
			Id:       assignmentGroupId,
			Name:     assignmentGroup.Name,
			CourseId: courseId,
			Weight:   assignmentGroup.Weight,
		})

		for _, assignment := range assignmentGroup.Assignments {
			assignmentId := ulid.Make().String()

			clos := make([]*entity.CourseLearningOutcome, 0)
			for _, clo := range assignment.CourseLearningOutcomeCodes {
				clos = append(clos, &entity.CourseLearningOutcome{
					Id: courseLearningOutcomeByCode[clo],
				})
			}
			assignmentsToCreate = append(assignmentsToCreate, entity.Assignment{
				Id:                               assignmentId,
				AssignmentGroupId:                assignmentGroupId,
				Name:                             assignment.Name,
				Description:                      assignment.Description,
				MaxScore:                         assignment.MaxScore,
				ExpectedScorePercentage:          assignment.ExpectedScorePercentage,
				ExpectedPassingStudentPercentage: assignment.ExpectedPassingStudentPercentage,
				CourseLearningOutcomes:           clos,
				IsIncludedInClo:                  assignment.IsIncludedInClo,
			})

			for _, score := range assignment.Scores {
				scoreId := ulid.Make().String()
				scoresToCreate = append(scoresToCreate, entity.Score{
					Id:           scoreId,
					AssignmentId: assignmentId,
					Score:        *score.Score,
					StudentId:    score.StudentId,
					UserId:       lecturerId,
				})
			}
		}

	}

	// prepare enrollments
	for _, studentId := range studentIds {
		enrollmentsToCreate = append(enrollmentsToCreate, entity.Enrollment{
			Id:        ulid.Make().String(),
			CourseId:  courseId,
			StudentId: studentId,
			Status:    entity.EnrollmentStatusEnroll,
		})
	}

	// let's go bro
	err = u.importerRepository.UpdateOrCreate(
		courseId,

		oldAssignmentGroupIds,
		OldAssignmentIds,
		oldCloIds,

		courseLearningOutcomes,
		groupsToCreate,
		assignmentsToCreate,
		enrollmentsToCreate,
		scoresToCreate,
	)

	return err
}
