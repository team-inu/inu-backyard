package request

type CreateCourseRequestPayload struct {
	Name       string `json:"name" validate:"required"`
	Code       string `json:"code" validate:"required"`
	SemesterId string `json:"semesterId" validate:"required"`
	LecturerId string `json:"lecturerId" validate:"required"`
}

type UpdateCourseRequestPayload struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	SemesterId string `json:"semesterId"`
	LecturerId string `json:"lecturerId"`
}
