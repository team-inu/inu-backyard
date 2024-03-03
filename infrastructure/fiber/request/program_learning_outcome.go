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

type UpdateProgramLearningOutcomePayload struct {
	Code            string  `json:"code" validate:"required"`
	DescriptionThai string  `json:"descriptionThai" validate:"required"`
	DescriptionEng  *string `json:"descriptionEng" validate:"required"`
	ProgramYear     int     `json:"programYear" validate:"required"`
	ProgrammeName   string  `json:"programmeName" validate:"required"`
}
