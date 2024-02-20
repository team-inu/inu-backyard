package request

type CreateScoreRequestPayload struct {
	StudentID    string  `json:"studentID" validate:"required"`
	Score        float64 `json:"score" validate:"required"`
	LecturerID   string  `json:"lecturerID" validate:"required"`
	AssignmentID string  `json:"assignmentID" validate:"required"`
}

type UpdateScoreRequestPayload struct {
	Score float64 `json:"score" validate:"required"`
}
