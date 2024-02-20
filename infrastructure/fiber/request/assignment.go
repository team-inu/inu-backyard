package request

type CreateAssignmentPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Score       *int   `json:"score" validate:"required"`
	Weight      *int   `json:"weight" validate:"required"`

	CourseLearningOutcomeID string `json:"courseLearningOutcomeId" validate:"required"`
}

type GetAssignmentsByParamsPayload struct {
	CourseLearningOutcomeID string `json:"courseLearningOutcomeId"`
}

type GetAssignmentsByCourseIDPayload struct {
	CourseID string `json:"courseId"`
}

type CreateBulkAssignmentsPayload struct {
	Assignments []CreateAssignmentPayload
}

type UpdateAssignmentRequestPayload struct {
	ID string `json:"id" validate:"required"`

	Name                    string `json:"name"`
	Description             string `json:"description"`
	Score                   int    `json:"score"`
	Weight                  int    `json:"weight"`
	CourseLearningOutcomeID string `json:"courseLearningOutcomeId"`
}

type DeleteAssignmentRequestPayload struct {
	ID string `json:"id" validate:"required"`
}
