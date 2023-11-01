package request

type CreateCourseLearningOutcomeBody struct {
	Code                        string `json:"code" validate:"required"`
	Description                 string `json:"description" validate:"required"`
	Weight                      int    `json:"weight"`
	SubProgramLearningOutcomeID string `json:"subProgramLearningOutcomeId"`
	ProgramOutcomeID            string `json:"programOutcomeId"`
	CourseId                    string `json:"courseId" validate:"required"`
	Status                      string `json:"status"`
}
