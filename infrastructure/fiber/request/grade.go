package request

type CreateGradePayload struct {
	StudentID string `json:"studentId" validate:"required"`
	Year      string `json:"year" validate:"required"`
	Grade     string `json:"grade" validate:"required"`
}

type UpdateGradePayload struct {
	StudentID string `json:"studentId"`
	Year      string `json:"year"`
	Grade     string `json:"grade"`
}
