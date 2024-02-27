package request

type CreateCourseLearningOutcomePayload struct {
	Code                                string   `json:"code" validate:"required"`
	Description                         string   `json:"description" validate:"required"`
	ExpectedPassingAssignmentPercentage float64  `json:"expectedPassingAssignmentPercentage" validate:"required"`
	ExpectedScorePercentage             float64  `json:"expectedScorePercentage" validate:"required"`
	ExpectedPassingStudentPercentage    float64  `json:"expectedPassingStudentPercentage" validate:"required"`
	Status                              string   `json:"status" validate:"required"`
	ProgramOutcomeId                    string   `json:"programOutcomeId" validate:"required"`
	CourseId                            string   `json:"courseId" validate:"required"`
	SubProgramLearningOutcomeIds        []string `json:"subProgramLearningOutcomeId" validate:"required"`
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
