package request

import "github.com/team-inu/inu-backyard/usecase"

type ImportCoursePayload struct {
	CourseId               string                                `json:"courseId" validate:"required"`
	StudentIds             []string                              `json:"studentIds" validate:"required,dive"`
	CourseLearningOutcomes []usecase.ImportCourseLearningOutcome `json:"courseLearningOutcomes" validate:"dive"`
	AssignmentGroups       []usecase.ImportAssignmentGroup       `json:"assignmentGroups" validate:"dive"`
}
