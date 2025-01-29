package request

type CreateAssignmentPayload struct {
	Name                             string   `json:"name" validate:"required"`
	Description                      string   `json:"description"`
	MaxScore                         *float64 `json:"maxScore" validate:"required"`
	ExpectedScorePercentage          *float64 `json:"expectedScorePercentage" validate:"required"`
	ExpectedPassingStudentPercentage *float64 `json:"expectedPassingStudentPercentage" validate:"required"`
	CourseLearningOutcomeIds         []string `json:"courseLearningOutcomeIds" validate:"required"`
	IsIncludedInClo                  *bool    `json:"isIncludedInClo" validate:"required"`
	AssignmentGroupId                string   `json:"assignmentGroupId" validate:"required"`
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
	Name                             string   `json:"name"`
	Description                      string   `json:"description"`
	MaxScore                         *float64 `json:"maxScore"`
	ExpectedPassingStudentPercentage *float64 `json:"expectedPassingStudentPercentage"`
	ExpectedScorePercentage          *float64 `json:"expectedScorePercentage"`
	IsIncludedInClo                  *bool    `json:"isIncludedInClo"`
}

type CreateLinkCourseLearningOutcomePayload struct {
	CourseLearningOutcomeIds []string `json:"courseLearningOutcomeIds" validate:"required"`
}
