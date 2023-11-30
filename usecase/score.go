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

func (u scoreUseCase) GetByID(id string) (*entity.Score, error) {
	score, err := u.scoreRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryScore, "cannot get score by id", err)
	}

	return score, nil
}

func (u scoreUseCase) Create(score float64, studentID string, assessmentID string, lecturerID string) (*entity.Score, error) {
	createdScore := entity.Score{
		ID:           ulid.Make().String(),
		Score:        score,
		StudentID:    studentID,
		LecturerID:   lecturerID,
		AssessmentID: assessmentID,
	}

	err := u.scoreRepo.Create(&createdScore)

	if err != nil {
		return nil, errs.New(errs.ErrCreateScore, "cannot create score", err)
	}

	return &createdScore, nil
}

func (u scoreUseCase) Update(scoreID string, score float64) error {
	existScore, err := u.GetByID(scoreID)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get score by id %s ", scoreID, err)
	} else if existScore == nil {
		return errs.New(errs.ErrScoreNotFound, "score not found", err)
	}
	err = u.scoreRepo.Update(scoreID, &entity.Score{
		Score:        score,
		StudentID:    existScore.StudentID,
		LecturerID:   existScore.LecturerID,
		AssessmentID: existScore.AssessmentID,
	})
	if err != nil {
		return errs.New(errs.ErrUpdateScore, "cannot update score", err)
	}

	return nil
}

func (u scoreUseCase) Delete(id string) error {
	existScore, err := u.GetByID(id)
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
