package request

type CreateLecturerPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Role      string `json:"role" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UpdateLecturerPayload struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type CreateBulkLecturerPayload struct {
	Lecturers []CreateLecturerPayload `json:"lecturers" validate:"dive"`
}
