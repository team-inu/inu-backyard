package request

type CreateDepartmentRequestBody struct {
	Name        string `json:"name" validate:"required"`
	FacultyName string `json:"faculty_name" validate:"required"`
}

type UpdateDepartmentRequestBody struct {
	Name        string `json:"name" validate:"required"`
	NewName     string `json:"new_name" validate:"required"`
	FacultyName string `json:"faculty_name" validate:"required"`
}

type DeleteDepartmentRequestBody struct {
	Name string `json:"name" validate:"required"`
}
