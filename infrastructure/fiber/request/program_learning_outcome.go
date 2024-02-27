package request

type CrateProgramLearningOutcome struct {
	Code                       string `validate:"required"`
	DescriptionThai            string `validate:"required"`
	DescriptionEng             string
	ProgramYear                int                              `validate:"required"`
	ProgrammeName              string                           `validate:"required"`
	SubProgramLearningOutcomes []CreateSubProgramLeaningOutcome `validate:"dive"`
}

type CreateProgramLearningOutcomePayload struct {
	ProgramLearningOutcomes []CrateProgramLearningOutcome `json:"programLearningOutcomes" validate:"required,dive"`
}

type UpdateProgramLearningOutcome struct {
	Code            string  `validate:"required"`
	DescriptionThai string  `validate:"required"`
	DescriptionEng  *string `validate:"required"`
	ProgramYear     int     `validate:"required"`
	ProgrammeName   string  `validate:"required"`
}

type UpdateProgramLearningOutcomePayload struct {
	ProgramLearningOutcomes []UpdateProgramLearningOutcome `json:"programLearningOutcomes" validate:"required,dive"`
}
