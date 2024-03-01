package request

type CreateUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Role      string `json:"role" validate:"required"`
	Password  string `json:"password" validate:"required"`
}

type UpdateUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Role      string `json:"role" validate:"required"`
}

type CreateBulkUserPayload struct {
	Users []CreateUserPayload `json:"users" validate:"dive"`
}
