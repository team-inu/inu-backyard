package request

type CreateProgramOutcome struct {
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateProgramOutcomePayload struct {
	ProgramOutcomes []CreateProgramOutcome `json:"programOutcomes" validate:"required,dive"`
}

type UpdateProgramOutcome struct {
	Code        string `validate:"required"`
	Name        string `validate:"required"`
	Description string `validate:"required"`
}

type UpdateProgramOutcomePayload struct {
	ProgramOutcomes []UpdateProgramOutcome `json:"programOutcomes" validate:"required,dive"`
}
