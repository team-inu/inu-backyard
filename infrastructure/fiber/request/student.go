package request

type CreateStudentPayload struct {
	KmuttId   string  `json:"kmuttId" validate:"required"`
	FirstName string  `json:"firstName" validate:"required"`
	LastName  string  `json:"lastName" validate:"required"`
	GPAX      float64 `json:"gpax" `
	MathGPA   float64 `json:"mathGpa" `
	EngGPA    float64 `json:"engGpa" `
	SciGPA    float64 `json:"sciGpa"`
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
	Students []CreateStudentPayload `json:"students"`
}

type UpdateStudentPayload struct {
	Name      string  `json:"name"`
	KmuttId   string  `json:"kmuttId"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	GPAX      float64 `json:"gpax" `
	MathGPA   float64 `json:"mathGpa" `
	EngGPA    float64 `json:"engGpa" `
	SciGPA    float64 `json:"sciGpa"`
	School    string  `json:"school"`
	Year      string  `json:"year"`
	Admission string  `json:"admission"`
	Remark    string  `json:"remark"`

	ProgrammeId    string `json:"programmeId"`
	DepartmentName string `json:"departmentName"`
}
