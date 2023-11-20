package validator

import (
	"errors"
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type PayloadValidator interface {
	Validate(payload interface{}, ctx *fiber.Ctx) (bool, error)
}

type payloadValidator struct {
	validator *validator.Validate
}

func NewPayloadValidator() PayloadValidator {
	return &payloadValidator{
		validator: validator.New(),
	}
}

func (v *payloadValidator) Validate(payload interface{}, ctx *fiber.Ctx) (bool, error) {
	if len(ctx.Body()) != 0 {
		if err := ctx.BodyParser(payload); err != nil {
			return false, errs.New(errs.ErrBodyParser, err.Error())
		}
	}
	if err := ctx.ParamsParser(payload); err != nil {
		return false, errs.New(errs.ErrParamsParser, err.Error())
	}
	if err := ctx.QueryParser(payload); err != nil {
		return false, errs.New(errs.ErrQueryParser, err.Error())
	}
	if err := fileParser(payload, ctx); err != nil {
		return false, errs.New(errs.ErrBodyParser, err.Error())
	}

	if errors := v.validateStruct(payload); errors != nil {
		return false, errs.NewValidationErr(errs.ErrPayloadValidator, "payload is invalid", errors)
	}
	return true, nil
}

func (v *payloadValidator) validateStruct(payload interface{}) []errs.ValidationErrorDetail {
	var errDetails []errs.ValidationErrorDetail

	if errors := v.validator.Struct(payload); errors != nil {
		for _, err := range errors.(validator.ValidationErrors) {
			detail := &errs.ValidationErrorDetail{
				Field: err.Field(),
				Type:  err.Tag(),
			}
			errDetails = append(errDetails, *detail)
		}
	}
	return errDetails
}

func fileParser(payload interface{}, ctx *fiber.Ctx) error {
	v := reflect.ValueOf(payload)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return errors.New("interface must be a pointer to struct")
	}
	v = v.Elem() // Unwrap interfae or pointer

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fileKey := field.Tag.Get("file")

		if fileKey != "" {
			fileHeader, err := ctx.FormFile(fileKey)
			if err != nil {
				return err
			}
			file, err := fileHeader.Open()
			if err != nil {
				return err
			}

			// TODO: contains unsafe operation, need better error handling
			v.Field(i).Set(reflect.ValueOf(file))
		}
	}
	return nil
}
