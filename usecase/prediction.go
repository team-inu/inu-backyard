package usecase

import (
	"os/exec"
	"strconv"
	"strings"

	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
	"github.com/team-inu/inu-backyard/internal/config"
)

type predictionUseCase struct {
	config config.FiberServerConfig
}

func NewPredictionUseCase(config config.FiberServerConfig) entity.PredictionUseCase {
	return &predictionUseCase{
		config: config,
	}
}

func (u predictionUseCase) CreatePrediction(requirement entity.PredictionRequirements) (*entity.Prediction, error) {
	cmd := exec.Command(
		"python3",
		"predict.py",
		u.config.Database.User,
		u.config.Database.Password,
		u.config.Database.Host,
		u.config.Database.Port,
		u.config.Database.DatabaseName,
		requirement.ProgrammeName,
		strconv.FormatFloat(requirement.OldGPAX, 'f', 2, 64),
		strconv.FormatFloat(requirement.MathGPA, 'f', 2, 64),
		strconv.FormatFloat(requirement.EngGPA, 'f', 2, 64),
		strconv.FormatFloat(requirement.SciGPA, 'f', 2, 64),
		requirement.School,
		requirement.Admission,
	)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, errs.New(errs.ErrUpdatePrediction, "found unexpected error when running python script", err)
	}

	outputValue, err := strconv.ParseFloat(strings.TrimSpace(string(output[:])), 64)
	if err != nil {
		return nil, errs.New(errs.ErrUpdatePrediction, "Output from python is in unexpected type", err)
	}

	prediction := entity.Prediction{
		PredictedGPAX: outputValue,
	}

	return &prediction, nil
}
