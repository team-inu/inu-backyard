package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	slice "github.com/team-inu/inu-backyard/internal/utils"
)

type scoreUseCase struct {
	scoreRepo         entity.ScoreRepository
	enrollmentUseCase entity.EnrollmentUseCase
	assignmentUseCase entity.AssignmentUseCase
	LecturerUseCase   entity.LecturerUseCase
}

func NewScoreUseCase(
	scoreRepo entity.ScoreRepository,
	enrollmentUseCase entity.EnrollmentUseCase,
	assignmentUseCase entity.AssignmentUseCase,
	lecturerUseCase entity.LecturerUseCase,
) entity.ScoreUseCase {
	return &scoreUseCase{
		scoreRepo:         scoreRepo,
		enrollmentUseCase: enrollmentUseCase,
		assignmentUseCase: assignmentUseCase,
		LecturerUseCase:   lecturerUseCase,
	}
}

func (u scoreUseCase) GetAll() ([]entity.Score, error) {
	scores, err := u.scoreRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryScore, "cannot get all scores", err)
	}

	return scores, nil
}

func (u scoreUseCase) GetById(id string) (*entity.Score, error) {
	score, err := u.scoreRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryScore, "cannot get score by id", err)
	}

	return score, nil
}

func (u scoreUseCase) CreateMany(lecturerId string, assignmentId string, studentScores []entity.StudentScore) error {
	lecturer, err := u.LecturerUseCase.GetById(lecturerId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get lecturer id %s to create score", lecturerId, err)
	} else if lecturer == nil {
		return errs.New(errs.ErrLecturerNotFound, "cannot get lecturer id %s to create score", lecturerId)
	}

	assignment, err := u.assignmentUseCase.GetById(assignmentId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get assignment id %s to create score", assignmentId, err)
	} else if assignment == nil {
		return errs.New(errs.ErrAssignmentNotFound, "cannot get assignment id %s to create score", assignmentId)
	}

	for _, studentScore := range studentScores {
		if studentScore.Score > float64(assignment.MaxScore) {
			return errs.New(errs.ErrCreateScore, "score %f of student id %s is more than max score of assignment (score: %d)", studentScore.Score, studentScore.StudentId, assignment.MaxScore)
		}
	}

	studentIds := []string{}
	for _, studentScore := range studentScores {
		studentIds = append(studentIds, studentScore.StudentId)
	}

	withStatus := entity.EnrollmentStatusEnroll
	joinedStudentIds, err := u.enrollmentUseCase.FilterJoinedStudent(studentIds, &withStatus)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get existed student ids while creating score")
	}

	nonJoinedStudentIds := slice.Subtraction(studentIds, joinedStudentIds)
	if len(nonJoinedStudentIds) != 0 {
		return errs.New(errs.ErrCreateAssignment, "there are non joined student ids")
	}

	scores := []entity.Score{}
	for _, studentScore := range studentScores {
		scores = append(scores, entity.Score{
			Id:           ulid.Make().String(),
			Score:        studentScore.Score,
			StudentId:    studentScore.StudentId,
			LecturerId:   lecturerId,
			AssignmentId: assignmentId,
		})
	}

	err = u.scoreRepo.CreateMany(scores)
	if err != nil {
		return errs.New(errs.ErrCreateScore, "cannot create score", err)
	}

	return nil
}

func (u scoreUseCase) Update(scoreId string, score float64) error {
	existScore, err := u.GetById(scoreId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get score by id %s ", scoreId, err)
	} else if existScore == nil {
		return errs.New(errs.ErrScoreNotFound, "score not found", err)
	}
	err = u.scoreRepo.Update(scoreId, &entity.Score{
		Score:        score,
		StudentId:    existScore.StudentId,
		LecturerId:   existScore.LecturerId,
		AssignmentId: existScore.AssignmentId,
	})
	if err != nil {
		return errs.New(errs.ErrUpdateScore, "cannot update score", err)
	}

	return nil
}

func (u scoreUseCase) Delete(id string) error {
	existScore, err := u.GetById(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get score by id %s ", id, err)
	} else if existScore == nil {
		return errs.New(errs.ErrScoreNotFound, "score not found", err)
	}
	err = u.scoreRepo.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteScore, "cannot delete score by id %s", id, err)
	}
	return nil
}
