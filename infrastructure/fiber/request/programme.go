package request

type CreateProgrammePayload struct {
	Name string `json:"name" validate:"required"`
}

type UpdateProgrammePayload struct {
	Name string `json:"name" validate:"required"`
}
