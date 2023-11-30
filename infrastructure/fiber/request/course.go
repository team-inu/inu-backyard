package request

type CreateCourseRequestPayload struct {
	Name       string `json:"name" validate:"required"`
	Code       string `json:"code" validate:"required"`
	SemesterID string `json:"semesterId" validate:"required"`
	LecturerID string `json:"lecturerId" validate:"required"`
}

type UpdateCourseRequestPayload struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	SemesterID string `json:"semesterId"`
	LecturerID string `json:"lecturerId"`
}
