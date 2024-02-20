package request

type CreateScoreRequestPayload struct {
	StudentId    string  `json:"studentId" validate:"required"`
	Score        float64 `json:"score" validate:"required"`
	LecturerId   string  `json:"lecturerId" validate:"required"`
	AssignmentId string  `json:"assignmentId" validate:"required"`
}

type UpdateScoreRequestPayload struct {
	Score float64 `json:"score" validate:"required"`
}
