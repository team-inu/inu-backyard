package request

type CreateStudentRequestBody struct {
	Name      string  `json:"name" validate:"required"`
	KmuttID   string  `json:"kmuttId" validate:"required"`
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	GPAX      float64 `json:"gpax" `
	MathGPA   float64 `json:"mathGpa" `
	EngGPA    float64 `json:"engGpa" `
	SciGPA    float64 `json:"sciGpa"`
	School    string  `json:"school" validate:"required"`
	Year      string  `json:"year" validate:"required"`
	Admission string  `json:"admission"`
	Remark    string  `json:"remark"`

	ProgrammeID    string `json:"programmeId" validate:"required"`
	DepartmentName string `json:"departmentName" validate:"required"`
}

type CreateBulkStudentsRequestBody struct {
	Students []CreateStudentRequestBody
}
