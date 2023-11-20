package request

type CreateCourseRequestPayload struct {
	Name       string `json:"name" validate:"required"`
	Code       string `json:"code" validate:"required"`
	Year       int    `json:"year" validate:"required"`
	LecturerID string `json:"lecturerId" validate:"required"`
}
