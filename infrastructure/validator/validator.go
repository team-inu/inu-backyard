package validator

import (
	go_validator "github.com/go-playground/validator/v10"
)

type Validator interface {
	Struct(data interface{}) []ErrorResponse
}

type validator struct {
	goValidator go_validator.Validate
}

func NewPayloadValidator() Validator {
	return &validator{
		goValidator: *go_validator.New(),
	}
}

type ErrorResponse struct {
	Error       bool
	FailedField string
	Tag         string
	Value       interface{}
}

func (v validator) Struct(data interface{}) []ErrorResponse {
	validationErrors := []ErrorResponse{}

	errs := v.goValidator.Struct(data)
	if errs != nil {
		for _, err := range errs.(go_validator.ValidationErrors) {
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
