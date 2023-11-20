package request

type CreateAssessmentPayload struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	Score       int    `json:"score" validate:"required"`
	Weight      int    `json:"weight" validate:"required"`

	CourseLearningOutcomeID string `json:"courseLearningOutcomeId" validate:"required"`
}

type GetAssessmentsByParamsPayload struct {
	CourseLearningOutcomeID string `json:"courseLearningOutcomeId"`
}

type GetAssessmentsByCourseIDPayload struct {
	CourseID string `json:"courseId"`
}

type CreateBulkAssessmentsPayload struct {
	Assessments []CreateAssessmentPayload
}

type UpdateAssessmentRequestPayload struct {
	ID string `json:"id" validate:"required"`

	Name                    string `json:"name"`
	Description             string `json:"description"`
	Score                   int    `json:"score"`
	Weight                  int    `json:"weight"`
	CourseLearningOutcomeID string `json:"courseLearningOutcomeId"`
}

type DeleteAssessmentRequestPayload struct {
	ID string `json:"id" validate:"required"`
}
