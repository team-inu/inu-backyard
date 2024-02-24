package request

type CrateProgramLearningOutcome struct {
	Code            string `validate:"required"`
	DescriptionThai string `validate:"required"`
	DescriptionEng  string `validate:"required"`
	ProgramYear     int    `validate:"required"`
	ProgrammeName   string `validate:"required"`
}

type CreateProgramLearningOutcomePayload struct {
	ProgramLearningOutcomes []CrateProgramLearningOutcome `json:"programLearningOutcomes" validate:"required,dive"`
}

type UpdateProgramLearningOutcomePayload struct {
	Code            string `json:"code"`
	DescriptionThai string `json:"descriptionThai"`
	DescriptionEng  string `json:"descriptionEng"`
	ProgramYear     int    `json:"programYear"`
	Programme       string `json:"programme"`
}
