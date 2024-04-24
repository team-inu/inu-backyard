package usecase

import (
	"os/exec"

	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type predictionUseCase struct {
	predictionRepo entity.PredictionRepository
}

func NewPredictionUseCase(predictionRepo entity.PredictionRepository) entity.PredictionUseCase {
	return &predictionUseCase{predictionRepo: predictionRepo}
}

func (u predictionUseCase) GetAll() ([]entity.Prediction, error) {
	predictions, err := u.predictionRepo.GetAll()
	if err != nil {
		return nil, errs.New(errs.ErrorPredictionNotFound, "cannot get all predictions", err)
	}

	return predictions, nil
}

func (u predictionUseCase) GetLatest() (*entity.Prediction, error) {
	prediction, err := u.predictionRepo.GetLatest()
	if err != nil {
		return nil, errs.New(errs.ErrorPredictionNotFound, "cannot get latest prediction", err)
	}

	return prediction, nil
}

func (u predictionUseCase) runTask(predictionId string) error {
	cmd := exec.Command("python3", "predict.py", predictionId)

	err := cmd.Run()
	if err != nil {
		if err = u.UpdatePrediction("", entity.PredictionStatusFailed); err != nil {
			return errs.New(errs.ErrorCreatePrediction, "xxx", err)
		}

		return errs.New(errs.ErrorUpdatePrediction, "cannot run python script", err)
	}

	out, err := cmd.CombinedOutput()
	if err != nil {
		if err = u.UpdatePrediction("", entity.PredictionStatusFailed); err != nil {
			return errs.New(errs.ErrorUpdatePrediction, "xxx", err)
		}

		return errs.New(errs.ErrorUpdatePrediction, "cannot get output from python script", err)
	}

	err = u.UpdatePrediction(string(out), entity.PredictionStatusDone)
	if err != nil {
		return errs.New(errs.ErrorUpdatePrediction, "xxx", err)
	}

	return nil
}

func (u predictionUseCase) UpdatePrediction(result string, status entity.PredictionStatus) error {
	return nil
}

func (u predictionUseCase) CreatePrediction() (*string, error) {
	id := ulid.Make().String()

	prediction := &entity.Prediction{
		Id:     id,
		Status: entity.PredictionStatusPending,
		Result: "",
	}

	go u.runTask(id)

	err := u.predictionRepo.CreatePrediction(prediction)
	if err != nil {
		return nil, errs.New(errs.ErrorCreatePrediction, "cannot create prediction", err)
	}

	return &id, nil
}
