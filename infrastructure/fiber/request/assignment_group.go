package request

type CreateAssignmentGroupPayload struct {
	Name     string  `json:"name" validate:"required"`
	Weight   float64 `json:"weight" validate:"required"`
	CourseId string  `json:"courseId" validate:"required"`
}

type UpdateAssignmentGroupPayload struct {
	Name   string  `json:"name" validate:"required"`
	Weight float64 `json:"weight" validate:"required"`
}
