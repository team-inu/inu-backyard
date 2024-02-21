package request

import "github.com/team-inu/inu-backyard/entity"

type CreateEnrollmentsPayload struct {
	CourseId   string                  `json:"courseId" validate:"required"`
	StudentIds []string                `json:"studentIds" validate:"required"`
	Status     entity.EnrollmentStatus `json:"status" validate:"required"`
}

type UpdateEnrollmentPayload struct {
	CourseId  string `json:"courseId"`
	StudentId string `json:"studentId"`
}
