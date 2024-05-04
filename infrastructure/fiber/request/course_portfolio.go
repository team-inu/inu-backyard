package request

import "github.com/team-inu/inu-backyard/entity"

type SaveCoursePortfolioPayload struct {
	CourseSummary     entity.CourseSummary     `json:"summary" validate:"required"`
	CourseDevelopment entity.CourseDevelopment `json:"development" validate:"required"`
}
