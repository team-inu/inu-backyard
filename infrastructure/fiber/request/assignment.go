package request

type CreateAssignmentPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Score       *int   `json:"score" validate:"required"`
	Weight      *int   `json:"weight" validate:"required"`

	CourseLearningOutcomeId string `json:"courseLearningOutcomeId" validate:"required"`
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
	Id string `json:"id" validate:"required"`

	Name                    string `json:"name"`
	Description             string `json:"description"`
	Score                   int    `json:"score"`
	Weight                  int    `json:"weight"`
	CourseLearningOutcomeId string `json:"courseLearningOutcomeId"`
}

type DeleteAssignmentRequestPayload struct {
	Id string `json:"id" validate:"required"`
}
