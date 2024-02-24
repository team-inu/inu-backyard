package request

import "github.com/team-inu/inu-backyard/entity"

type CreateProgramLearningOutcomePayload struct {
	ProgramLearningOutcomes []entity.CrateProgramLearningOutcomeDto `json:"programLearningOutcomes" validate:"required,dive"`
}

type UpdateProgramLearningOutcomePayload struct {
	Code            string `json:"code"`
	DescriptionThai string `json:"descriptionThai"`
	DescriptionEng  string `json:"descriptionEng"`
	ProgramYear     int    `json:"programYear"`
	Programme       string `json:"programme"`
}
