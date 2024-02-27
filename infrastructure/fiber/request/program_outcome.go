package request

type CreateProgramOutcome struct {
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateProgramOutcomePayload struct {
	ProgramOutcomes []CreateProgramOutcome `json:"programOutcomes" validate:"required,dive"`
}

type UpdateProgramOutcomePayload struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
