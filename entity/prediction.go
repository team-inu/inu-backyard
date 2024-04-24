package entity

type PredictionStatus string

const (
	PredictionStatusPending PredictionStatus = "PENDING"
	PredictionStatusFailed  PredictionStatus = "FAILED"
	PredictionStatusDone    PredictionStatus = "DONE"
)

type Prediction struct {
	Id     string           `json:"id" gorm:"primaryKey;type:char(255)"`
	Status PredictionStatus `json:"status"`
	Result string           `json:"result"`
}

type PredictionRepository interface {
	GetAll() ([]Prediction, error)
	GetLatest() (*Prediction, error)
	CreatePrediction(prediction *Prediction) error
}

type PredictionUseCase interface {
	GetAll() ([]Prediction, error)
	GetLatest() (*Prediction, error)
	CreatePrediction() (*string, error)
}
