package request

type CreateSubProgramLeaningOutcome struct {
	Code                     string `validate:"required"`
	DescriptionThai          string `validate:"required"`
	DescriptionEng           string `validate:"required"`
	ProgramLearningOutcomeId string ``
}

type CreateSubProgramLearningOutcomePayload struct {
	SubProgramLearningOutcomes []CreateSubProgramLeaningOutcome `validate:"required,dive"`
}

type UpdateSubProgramLearningOutcomePayload struct {
	Code                     string `json:"code"`
	DescriptionThai          string `json:"descriptionThai"`
	DescriptionEng           string `json:"descriptionEng"`
	ProgramLearningOutcomeId string `json:"programLearningOutcomeId"`
}
