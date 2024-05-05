package entity

type Prediction struct {
	PredictedGPAX float64 `json:"predictedGPAX"`
}

type PredictionRequirements struct {
	ProgrammeName string
	OldGPAX       float64
	MathGPA       float64
	EngGPA        float64
	SciGPA        float64
	School        string
	Admission     string
}

type PredictionUseCase interface {
	CreatePrediction(requirements PredictionRequirements) (*Prediction, error)
}
