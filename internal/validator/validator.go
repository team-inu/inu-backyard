package validator

import (
	goValidator "github.com/go-playground/validator/v10"
)

type Validator interface {
	Struct(data interface{}) []ErrorResponse
}

type validator struct {
	goValidator *goValidator.Validate
}

func NewPayloadValidator() Validator {
	return &validator{
		goValidator: goValidator.New(),
	}
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func (v *validator) Struct(data interface{}) []ErrorResponse {
	var validationErrors []ErrorResponse
	errs := v.goValidator.Struct(data)
	if errs != nil {
		for _, err := range errs.(goValidator.ValidationErrors) {
			var elem ErrorResponse

			elem.FailedField = err.Field()
			elem.Tag = err.Tag()
			elem.Value = err.Value()
			elem.Error = true

			validationErrors = append(validationErrors, elem)
		}
	}

	return validationErrors
}
