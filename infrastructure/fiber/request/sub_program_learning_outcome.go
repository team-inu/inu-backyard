package request

type CreateSubProgramLearningOutcomeBody struct {
	Code                     string `json:"code" validate:"required"`
	DescriptionThai          string `json:"descriptionThai" validate:"required"`
	DescriptionEng           string `json:"descriptionEng" validate:"required"`
	ProgramLearningOutcomeID string `json:"programLearningOutcomeId"`
}
