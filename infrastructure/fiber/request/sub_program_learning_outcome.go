package request

type CreateSubProgramLeaningOutcome struct {
	Code                     string `validate:"required"`
	DescriptionThai          string `validate:"required"`
	DescriptionEng           string
	ProgramLearningOutcomeId string ``
}

type CreateSubProgramLearningOutcomePayload struct {
	SubProgramLearningOutcomes []CreateSubProgramLeaningOutcome `json:"subProgramLearningOutcomes" validate:"required,dive"`
}

type UpdateSubProgramLearningOutcome struct {
	Code                     string  `validate:"required"`
	DescriptionThai          string  `validate:"required"`
	DescriptionEng           *string `validate:"required"`
	ProgramLearningOutcomeId string  `validate:"required"`
}

type UpdateSubProgramLearningOutcomePayload struct {
	SubProgramLearningOutcomes []UpdateSubProgramLearningOutcome `json:"subProgramLearningOutcomes" validate:"required,dive"`
}
