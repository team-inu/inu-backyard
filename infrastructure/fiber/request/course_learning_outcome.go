package request

type CreateCourseLearningOutcomePayload struct {
	Code                        string `json:"code" validate:"required"`
	Description                 string `json:"description" validate:"required"`
	Weight                      int    `json:"weight"`
	SubProgramLearningOutcomeId string `json:"subProgramLearningOutcomeId"`
	ProgramOutcomeId            string `json:"programOutcomeId"`
	CourseId                    string `json:"courseId" validate:"required"`
	Status                      string `json:"status"`
}

type UpdateCourseLearningOutcomePayload struct {
	Code                        string `json:"code"`
	Description                 string `json:"description"`
	Weight                      int    `json:"weight"`
	SubProgramLearningOutcomeId string `json:"subProgramLearningOutcomeId"`
	ProgramOutcomeId            string `json:"programOutcomeId"`
	CourseId                    string `json:"courseId"`
	Status                      string `json:"status"`
}
