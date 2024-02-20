package request

type CreateEnrollmentPayload struct {
	CourseId  string `json:"courseId" validate:"required"`
	StudentId string `json:"studentId" validate:"required"`
}

type UpdateEnrollmentPayload struct {
	CourseId  string `json:"courseId"`
	StudentId string `json:"studentId"`
}
