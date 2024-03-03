package request

type CreateDepartmentRequestPayload struct {
	Name        string `json:"name" validate:"required"`
	FacultyName string `json:"faculty_name" validate:"required"`
}

type UpdateDepartmentRequestPayload struct {
	Name        string `json:"name" validate:"required"`
	NewName     string `json:"new_name" validate:"required"`
	FacultyName string `json:"faculty_name" validate:"required"`
}

type DeleteDepartmentRequestPayload struct {
	Name string `json:"name" validate:"required"`
}
