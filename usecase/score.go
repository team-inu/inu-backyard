package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type scoreUseCase struct {
	scoreRepo entity.ScoreRepository
}

func NewScoreUseCase(scoreRepo entity.ScoreRepository) entity.ScoreUsecase {
	return &scoreUseCase{scoreRepo: scoreRepo}
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

func (u scoreUseCase) Create(score float64, studentId string, assignmentId string, lecturerId string) (*entity.Score, error) {
	createdScore := entity.Score{
		Id:           ulid.Make().String(),
		Score:        score,
		StudentId:    studentId,
		LecturerId:   lecturerId,
		AssignmentId: assignmentId,
	}

	err := u.scoreRepo.Create(&createdScore)

	if err != nil {
		return nil, errs.New(errs.ErrCreateScore, "cannot create score", err)
	}

	return &createdScore, nil
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
