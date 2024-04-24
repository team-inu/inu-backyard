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
	GetById(id string) (*Prediction, error)
	GetAll() ([]Prediction, error)
	GetLatest() (*Prediction, error)
	CreatePrediction(prediction *Prediction) error
	Update(id string, prediction *Prediction) error
}

type PredictionUseCase interface {
	GetById(id string) (*Prediction, error)
	GetAll() ([]Prediction, error)
	GetLatest() (*Prediction, error)
	CreatePrediction() (*string, error)
	Update(id string, status PredictionStatus, result string) error
}
