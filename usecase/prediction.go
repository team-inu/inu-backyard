package usecase

import (
	"os/exec"

	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"github.com/team-inu/inu-backyard/internal/config"
)

type predictionUseCase struct {
	predictionRepo entity.PredictionRepository
	config         config.FiberServerConfig
}

func NewPredictionUseCase(predictionRepo entity.PredictionRepository, config config.FiberServerConfig) entity.PredictionUseCase {
	return &predictionUseCase{
		predictionRepo: predictionRepo,
		config:         config,
	}
}

func (u predictionUseCase) GetAll() ([]entity.Prediction, error) {

	predictions, err := u.predictionRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrPredictionNotFound, "cannot get all predictions", err)
	}

	return predictions, nil
}

func (u predictionUseCase) GetLatest() (*entity.Prediction, error) {
	prediction, err := u.predictionRepo.GetLatest()
	if err != nil {
		return nil, errs.New(errs.ErrPredictionNotFound, "cannot get latest prediction", err)
	}

	return prediction, nil
}

func (u predictionUseCase) GetById(id string) (*entity.Prediction, error) {
	prediction, err := u.predictionRepo.GetById(id)
	if err != nil {
		return nil, errs.New(errs.ErrQueryPrediction, "cannot get prediction by id %s", id, err)
	}

	return prediction, nil
}

// TODO: add visibility
func (u predictionUseCase) runTask(predictionId string) error {
	cmd := exec.Command(
		"python3",
		"predict.py",
		u.config.Database.User,
		u.config.Database.Password,
		u.config.Database.Host,
		u.config.Database.Port,
		u.config.Database.DatabaseName,
		predictionId,
	)

	err := cmd.Run()
	if err != nil {
		if err = u.Update(predictionId, entity.PredictionStatusFailed, ""); err != nil {
			return errs.New(errs.ErrUpdatePrediction, "cannot update prediction status while facing unexpected error from python script", err)
		}

		return errs.New(errs.ErrUpdatePrediction, "found unexpected error when running python script", err)
	}

	return nil
}

func (u predictionUseCase) Update(predictionId string, status entity.PredictionStatus, result string) error {
	existedPrediction, err := u.GetById(predictionId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot get prediction id %s to update", predictionId, err)
	} else if existedPrediction == nil {
		return errs.New(errs.ErrPredictionNotFound, "prediction id %s not found to update", predictionId)
	}

	err = u.predictionRepo.Update(predictionId, &entity.Prediction{
		Status: status,
		Result: result,
	})

	if err != nil {
		return errs.New(errs.ErrUpdatePrediction, "cannot update prediction by id %s", predictionId, err)
	}

	return nil
}

func (u predictionUseCase) CreatePrediction() (*string, error) {
	id := ulid.Make().String()

	prediction := &entity.Prediction{
		Id:     id,
		Status: entity.PredictionStatusPending,
		Result: "",
	}

	err := u.predictionRepo.CreatePrediction(prediction)
	if err != nil {
		return nil, errs.New(errs.ErrCreatePrediction, "cannot create prediction", err)
	}

	go u.runTask(id)

	return &id, nil
}
