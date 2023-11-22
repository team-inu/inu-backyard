package request

type CreateEnrollmentPayload struct {
	CourseID  string `json:"courseId" validate:"required"`
	StudentID string `json:"studentId" validate:"required"`
}

type UpdateEnrollmentPayload struct {
	CourseID  string `json:"courseId"`
	StudentID string `json:"studentId"`
}
