package request

type CreateFacultyRequestBody struct {
	Name string `json:"name" validate:"required"`
}

type UpdateFacultyRequestBody struct {
	Name    string `json:"name" validate:"required"`
	NewName string `json:"newName" validate:"required"`
}

type DeleteFacultyRequestBody struct {
	Name string `json:"name" validate:"required"`
}
