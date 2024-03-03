package request

type CreateSemesterPayload struct {
	Year             int `json:"year" validate:"required"`
	SemesterSequence int `json:"semesterSequence" validate:"required"`
}

type UpdateSemesterPayload struct {
	Year             int `json:"year" validate:"required"`
	SemesterSequence int `json:"semesterSequence" validate:"required"`
}
