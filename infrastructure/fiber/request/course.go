package request

import "github.com/team-inu/inu-backyard/entity"

type CreateCourseRequestPayload struct {
	SemesterId                   string  `json:"semesterId" validate:"required"`
	UserId                       string  `json:"userId" validate:"required"`
	Name                         string  `json:"name" validate:"required"`
	Code                         string  `json:"code" validate:"required"`
	Curriculum                   string  `json:"curriculum" validate:"required"`
	Description                  string  `json:"description" validate:"required"`
	ExpectedPassingCloPercentage float64 `json:"expectedPassingCloPercentage" validate:"required"`
	entity.CriteriaGrade
}

type UpdateCourseRequestPayload struct {
	Name                         string  `json:"name" validate:"required"`
	Code                         string  `json:"code" validate:"required"`
	Curriculum                   string  `json:"curriculum" validate:"required"`
	Description                  string  `json:"description" validate:"required"`
	ExpectedPassingCloPercentage float64 `json:"expectedPassingCloPercentage" validate:"required"`
	IsPortfolioCompleted         *bool   `json:"isPortfolioCompleted" validate:"required"`
	entity.CriteriaGrade
}
