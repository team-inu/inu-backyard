package request

type CreateAssignmentGroupPayload struct {
	Name     string `json:"name" validate:"required"`
	CourseId string `json:"courseId" validate:"required"`
}

type UpdateAssignmentGroupPayload struct {
	Name string `json:"name" validate:"required"`
}
