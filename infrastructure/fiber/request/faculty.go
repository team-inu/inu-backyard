package request

type CreateFacultyRequestPayload struct {
	Name string `json:"name" validate:"required"`
}

type UpdateFacultyRequestPayload struct {
	Name    string `json:"name" validate:"required"`
	NewName string `json:"newName" validate:"required"`
}

type DeleteFacultyRequestPayload struct {
	Name string `json:"name" validate:"required"`
}
