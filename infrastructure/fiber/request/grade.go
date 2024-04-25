package request

import "github.com/team-inu/inu-backyard/entity"

type CreateGradePayload struct {
	StudentId string `json:"studentId" validate:"required"`
	Year      string `json:"year" validate:"required"`
	Grade     string `json:"grade" validate:"required"`
}

type CreateManyGradesPayload struct {
	StudentGrade     []entity.StudentGrade `json:"studentGrade" validate:"dive"`
	Year             int                   `json:"year" validate:"required"`
	SemesterSequence string                `json:"semesterSequence" validate:"required"`
}

type UpdateGradePayload struct {
	StudentId string `json:"studentId"`
	Year      string `json:"year"`
	Grade     string `json:"grade"`
}
