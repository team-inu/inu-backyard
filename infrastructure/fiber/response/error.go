package response

import (
	"github.com/gofiber/fiber/v2"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

var DomainErrCodeToHttpStatus = map[int]int{
	errs.ErrInternal:   fiber.StatusInternalServerError,
	errs.ErrRoute:      fiber.StatusNotFound,
	errs.ErrFileSystem: fiber.StatusInternalServerError,

	errs.ErrAuthHeader:       fiber.StatusBadRequest,
	errs.ErrPayloadValidator: fiber.StatusBadRequest,
	errs.ErrBodyParser:       fiber.StatusUnprocessableEntity,
	errs.ErrQueryParser:      fiber.StatusUnprocessableEntity,
	errs.ErrParamsParser:     fiber.StatusUnprocessableEntity,

	errs.ErrStudentNotFound: fiber.StatusNotFound,
	errs.ErrQueryStudent:    fiber.StatusInternalServerError,
	errs.ErrCreateStudent:   fiber.StatusInternalServerError,
	errs.ErrUpdateStudent:   fiber.StatusInternalServerError,
	errs.ErrDeleteStudent:   fiber.StatusInternalServerError,

	errs.ErrCourseNotFound: fiber.StatusNotFound,
	errs.ErrQueryCourse:    fiber.StatusInternalServerError,
	errs.ErrCreateCourse:   fiber.StatusInternalServerError,
	errs.ErrUpdateCourse:   fiber.StatusInternalServerError,
	errs.ErrDeleteCourse:   fiber.StatusInternalServerError,

	errs.ErrCLONotFound: fiber.StatusNotFound,
	errs.ErrQueryCLO:    fiber.StatusInternalServerError,
	errs.ErrCreateCLO:   fiber.StatusInternalServerError,
	errs.ErrUpdateCLO:   fiber.StatusInternalServerError,
	errs.ErrDeleteCLO:   fiber.StatusInternalServerError,

	errs.ErrPLONotFound: fiber.StatusNotFound,
	errs.ErrQueryPLO:    fiber.StatusInternalServerError,
	errs.ErrCreatePLO:   fiber.StatusInternalServerError,
	errs.ErrUpdatePLO:   fiber.StatusInternalServerError,
	errs.ErrDeletePLO:   fiber.StatusInternalServerError,

	errs.ErrPONotFound: fiber.StatusNotFound,
	errs.ErrQueryPO:    fiber.StatusInternalServerError,
	errs.ErrCreatePO:   fiber.StatusInternalServerError,
	errs.ErrUpdatePO:   fiber.StatusInternalServerError,
	errs.ErrDeletePO:   fiber.StatusInternalServerError,

	errs.ErrSubPLONotFound: fiber.StatusNotFound,
	errs.ErrQuerySubPLO:    fiber.StatusInternalServerError,
	errs.ErrCreateSubPLO:   fiber.StatusInternalServerError,
	errs.ErrUpdateSubPLO:   fiber.StatusInternalServerError,
	errs.ErrDeleteSubPLO:   fiber.StatusInternalServerError,
}
