package request

type CreateSubProgramLearningOutcomePayload struct {
	Code                     string `json:"code" validate:"required"`
	DescriptionThai          string `json:"descriptionThai" validate:"required"`
	DescriptionEng           string `json:"descriptionEng" validate:"required"`
	ProgramLearningOutcomeID string `json:"programLearningOutcomeId"`
}

type UpdateSubProgramLearningOutcomePayload struct {
	Code                     string `json:"code"`
	DescriptionThai          string `json:"descriptionThai"`
	DescriptionEng           string `json:"descriptionEng"`
	ProgramLearningOutcomeID string `json:"programLearningOutcomeId"`
}
