package request

type CreateStudentRequestBody struct {
	Name      string `json:"name" validate:"required"`
	KmuttID   string `json:"kmuttId" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
}
