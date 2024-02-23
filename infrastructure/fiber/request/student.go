package request

type CreateStudentPayload struct {
	KmuttId   string  `json:"kmuttId" validate:"required"`
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	GPAX      float64 `json:"gpax" `
	MathGPA   float64 `json:"mathGPA" `
	EngGPA    float64 `json:"engGPA" `
	SciGPA    float64 `json:"sciGPA"`
	School    string  `json:"school"`
	City      string  `json:"city"`
	Email     string  `json:"email"`
	Year      string  `json:"year" validate:"required"`
	Admission string  `json:"admission"`
	Remark    string  `json:"remark"`

	ProgrammeId    string `json:"programmeId" validate:"required"`
	DepartmentName string `json:"departmentName" validate:"required"`
}

type GetStudentsByParamsPayload struct {
	Year           string `json:"year"`
	ProgrammeId    string `json:"programmeId"`
	DepartmentName string `json:"departmentId"`
}

type CreateBulkStudentsPayload struct {
	Students []CreateStudentPayload `json:"students" validate:"dive"`
}

type UpdateStudentPayload struct {
	Name      string  `json:"name"`
	KmuttId   string  `json:"kmuttId"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	GPAX      float64 `json:"gpax" `
	MathGPA   float64 `json:"mathGPA" `
	EngGPA    float64 `json:"engGPA" `
	SciGPA    float64 `json:"sciGPA"`
	School    string  `json:"school"`
	Year      string  `json:"year"`
	Admission string  `json:"admission"`
	Remark    string  `json:"remark"`

	ProgrammeId    string `json:"programmeId"`
	DepartmentName string `json:"departmentName"`
}
