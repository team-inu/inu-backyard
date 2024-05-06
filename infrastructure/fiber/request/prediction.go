package request

type PredictPayload struct {
	ProgrammeName string   `json:"programmeName" validate:"required"`
	GPAX          *float64 `json:"gpax" validate:"required"`
	MathGPA       *float64 `json:"mathGPA" validate:"required"`
	EngGPA        *float64 `json:"engGPA" validate:"required"`
	SciGPA        *float64 `json:"sciGPA" validate:"required"`
	School        string   `json:"school" validate:"required"`
	Admission     string   `json:"admission" validate:"required"`
}
