package request

type CreateSubProgramLearningOutcomePayload struct {
	Code                     string `json:"code" validate:"required"`
	DescriptionThai          string `json:"descriptionThai" validate:"required"`
	DescriptionEng           string `json:"descriptionEng" validate:"required"`
	ProgramLearningOutcomeId string `json:"programLearningOutcomeId"`
}

type UpdateSubProgramLearningOutcomePayload struct {
	Code                     string `json:"code"`
	DescriptionThai          string `json:"descriptionThai"`
	DescriptionEng           string `json:"descriptionEng"`
	ProgramLearningOutcomeId string `json:"programLearningOutcomeId"`
}
