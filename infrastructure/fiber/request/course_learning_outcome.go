package request

type CreateCourseLearningOutcomePayload struct {
	Code                                string   `json:"code" validate:"required"`
	Description                         string   `json:"description" validate:"required"`
	ExpectedPassingAssignmentPercentage float64  `json:"expectedPassingAssignmentPercentage" validate:"required"`
	ExpectedPassingStudentPercentage    float64  `json:"expectedPassingStudentPercentage" validate:"required"`
	Status                              string   `json:"status" validate:"required"`
	ProgramOutcomeId                    string   `json:"programOutcomeId" validate:"required"`
	CourseId                            string   `json:"courseId" validate:"required"`
	SubProgramLearningOutcomeIds        []string `json:"subProgramLearningOutcomeId" validate:"required"`
}

type UpdateCourseLearningOutcomePayload struct {
	Code                                string  `json:"code"`
	Description                         string  `json:"description"`
	ExpectedPassingAssignmentPercentage float64 `json:"expectedPassingAssignmentPercentage" validate:"required"`
	ExpectedPassingStudentPercentage    float64 `json:"expectedPassingStudentPercentage" validate:"required"`
	Status                              string  `json:"status" validate:"required"`
	ProgramOutcomeId                    string  `json:"programOutcomeId" validate:"required"`
}

type CreateLinkSubProgramLearningOutcomePayload struct {
	SubProgramLearningOutcomeId []string `json:"subProgramLearningOutcomeId" validate:"required"`
}
