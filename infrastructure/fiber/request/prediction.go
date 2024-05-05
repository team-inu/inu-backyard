package request

type PredictPayload struct {
	ProgrammeName string  `validate:"required"`
	OldGPAX       float64 `validate:"required"`
	MathGPA       float64 `validate:"required"`
	EngGPA        float64 `validate:"required"`
	SciGPA        float64 `validate:"required"`
	School        string  `validate:"required"`
	Admission     string  `validate:"required"`
}
