package request

type CreateProgramLearningOutcomePayload struct {
	Code            string `json:"code" validate:"required"`
	DescriptionThai string `json:"descriptionThai" validate:"required"`
	DescriptionEng  string `json:"descriptionEng" validate:"required"`
	ProgramYear     int    `json:"programYear" validate:"required"`
}

type UpdateProgramLearningOutcomePayload struct {
	Code            string `json:"code"`
	DescriptionThai string `json:"descriptionThai"`
	DescriptionEng  string `json:"descriptionEng"`
	ProgramYear     int    `json:"programYear"`
}
