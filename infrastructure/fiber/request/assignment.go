package request

type CreateAssignmentPayload struct {
	Name                             string   `json:"name" validate:"required"`
	Description                      string   `json:"description"`
	MaxScore                         *int     `json:"maxScore" validate:"required"`
	Weight                           *int     `json:"weight" validate:"required"`
	ExpectedScorePercentage          *float64 `json:"expectedScorePercentage" validate:"required"`
	ExpectedPassingStudentPercentage *float64 `json:"expectedPassingStudentPercentage" validate:"required"`
	CourseLearningOutcomeIds         []string `json:"courseLearningOutcomeIds" validate:"required"`
}

type GetAssignmentsByParamsPayload struct {
	CourseLearningOutcomeId string `json:"courseLearningOutcomeId"`
}

type GetAssignmentsByCourseIdPayload struct {
	CourseId string `json:"courseId"`
}

type CreateBulkAssignmentsPayload struct {
	Assignments []CreateAssignmentPayload
}

type UpdateAssignmentRequestPayload struct {
	Name                             string  `json:"name"`
	Description                      string  `json:"description"`
	Weight                           int     `json:"weight"`
	MaxScore                         int     `json:"maxScore"`
	ExpectedPassingStudentPercentage float64 `json:"expectedPassingStudentPercentage"`
	ExpectedScorePercentage          float64 `json:"expectedScorePercentage"`
}

type CreateLinkCourseLearningOutcomePayload struct {
	CourseLearningOutcomeIds []string `json:"courseLearningOutcomeIds" validate:"required"`
}
