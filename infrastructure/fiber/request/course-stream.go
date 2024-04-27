package request

import (
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type CreateCourseStreamRequestPayload struct {
	FromCourseId   string                  `json:"fromCourseId" validate:"required"`
	TargetCourseId string                  `json:"targetCourseId" validate:"required"`
	StreamType     entity.CourseStreamType `json:"streamType" validate:"required"`
	Comment        string                  `json:"comment" validate:"required"`
}

type GetCourseStreamRequestPayload struct {
	FromCourseId   string `json:"fromCourseId"`
	TargetCourseId string `json:"targetCourseId" `
}

func (p GetCourseStreamRequestPayload) Validate() *errs.DomainError {
	if p.TargetCourseId != "" && p.FromCourseId != "" {
		return errs.New(errs.ErrPayloadValidator, "targetCourseId OR fromCourseId only one")
	}

	if p.TargetCourseId == "" && p.FromCourseId == "" {
		return errs.New(errs.ErrPayloadValidator, "must have query at least targetCourseId OR fromCourseId only one")
	}

	return nil
}
