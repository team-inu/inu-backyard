package request

type CreateProgramOutcomePayload struct {
	SemesterId  string `json:"semesterId" validate:"required"`
	Code        string `json:"code" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateProgramOutcomePayload struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
