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

func (s scoreUseCase) GetAll() ([]entity.Score, error) {
	scores, err := s.scoreRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrQueryScore, "cannot get all scores", err)
	}

	return scores, nil
}

func (s scoreUseCase) GetByID(id string) (*entity.Score, error) {
	score, err := s.scoreRepo.GetByID(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryScore, "cannot get score by id", err)
	}

	return score, nil
}

func (s scoreUseCase) Create(score float64, studentID string, assessmentID string, lecturerID string) (*entity.Score, error) {
	createdScore := entity.Score{
		ID:           ulid.Make().String(),
		Score:        score,
		StudentID:    studentID,
		LecturerID:   lecturerID,
		AssessmentID: assessmentID,
	}

	err := s.scoreRepo.Create(&createdScore)

	if err != nil {
		return nil, errs.New(errs.ErrCreateScore, "cannot create score", err)
	}

	return &createdScore, nil
}

func (s scoreUseCase) Update(scoreID string, score float64) error {
	oldScore, err := s.scoreRepo.GetByID(scoreID)
	if err != nil {
		return errs.New(errs.ErrQueryScore, "cannot get score by id", err)
	}
	err = s.scoreRepo.Update(&entity.Score{
		ID:           oldScore.ID,
		Score:        score,
		StudentID:    oldScore.StudentID,
		LecturerID:   oldScore.LecturerID,
		AssessmentID: oldScore.AssessmentID,
	})
	if err != nil {
		return errs.New(errs.ErrUpdateScore, "cannot update score", err)
	}

	return nil
}

func (s scoreUseCase) Delete(id string) error {
	_, err := s.scoreRepo.GetByID(id)

	if err != nil {
		return errs.New(errs.ErrQueryScore, "cannot get score by id", err)
	}
	err = s.scoreRepo.Delete(id)

	if err != nil {
		return errs.New(errs.ErrDeleteScore, "cannot delete score", err)
	}

	return nil
}
