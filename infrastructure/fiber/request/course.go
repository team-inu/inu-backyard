package request

import "github.com/team-inu/inu-backyard/entity"

type CreateCourseRequestPayload struct {
	SemesterId    string                `json:"semesterId" validate:"required"`
	LecturerId    string                `json:"lecturerId" validate:"required"`
	Name          string                `json:"name" validate:"required"`
	Code          string                `json:"code" validate:"required"`
	Curriculum    string                `json:"curriculum" validate:"required"`
	Description   string                `json:"description" validate:"required"`
	CriteriaGrade *entity.CriteriaGrade `json:"criteriaGrade" validate:"required"`
}

type UpdateCourseRequestPayload struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	SemesterId string `json:"semesterId"`
	LecturerId string `json:"lecturerId"`
}
