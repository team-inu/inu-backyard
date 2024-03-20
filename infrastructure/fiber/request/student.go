package request

type CreateStudentPayload struct {
	KmuttId   string   `json:"kmuttId" validate:"required"`
	FirstName string   `json:"firstName" validate:"required"`
	LastName  string   `json:"lastName" validate:"required"`
	GPAX      *float64 `json:"gpax" validate:"required"`
	MathGPA   *float64 `json:"mathGPA" validate:"required"`
	EngGPA    *float64 `json:"engGPA" validate:"required"`
	SciGPA    *float64 `json:"sciGPA" validate:"required"`
	School    string   `json:"school" validate:"required"`
	City      string   `json:"city" validate:"required"`
	Email     string   `json:"email" validate:"required"`
	Year      string   `json:"year" validate:"required"`
	Admission string   `json:"admission" validate:"required"`
	Remark    string   `json:"remark"`

	ProgrammeName  string `json:"programmeName" validate:"required"`
	DepartmentName string `json:"departmentName" validate:"required"`
}

type GetStudentsByParamsPayload struct {
	Year           string `json:"year"`
	ProgrammeName  string `json:"programmeName"`
	DepartmentName string `json:"departmentId"`
}

type CreateBulkStudentsPayload struct {
	Students []CreateStudentPayload `json:"students" validate:"dive"`
}

type UpdateStudentPayload struct {
	KmuttId   string   `json:"kmuttId" validate:"required"`
	FirstName string   `json:"firstName" validate:"required"`
	LastName  string   `json:"lastName" validate:"required"`
	GPAX      *float64 `json:"gpax" validate:"required"`
	MathGPA   *float64 `json:"mathGPA" validate:"required"`
	EngGPA    *float64 `json:"engGPA" validate:"required"`
	SciGPA    *float64 `json:"sciGPA" validate:"required"`
	School    string   `json:"school" validate:"required"`
	City      string   `json:"city" validate:"required"`
	Email     string   `json:"email" validate:"required"`
	Year      string   `json:"year" validate:"required"`
	Admission string   `json:"admission" validate:"required"`
	Remark    *string  `json:"remark" validate:"required"`

	ProgrammeName  string `json:"programmeName" validate:"required"`
	DepartmentName string `json:"departmentName" validate:"required"`
}
