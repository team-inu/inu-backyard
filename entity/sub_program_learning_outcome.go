package entity

type SubProgramLearningOutcome struct {
	ID                       string `json:"id" gorm:"primaryKey;type:char(255)"`
	Code                     string `json:"code"`
	DescriptionThai          string `json:"descriptionThai"`
	DescriptionEng           string `json:"descriptionEng"`
	ProgramLearningOutcomeID string `json:"programLearningOutcomeId"`

	ProgramLearningOutcome ProgramLearningOutcome `gorm:"references:Code"`
}

type SubProgramLearningOutcomeRepository interface {
	GetAll() ([]SubProgramLearningOutcome, error)
	GetByID(id string) (*SubProgramLearningOutcome, error)
	Create(programLearningOutcome *SubProgramLearningOutcome) error
	Update(id string, programLearningOutcome *SubProgramLearningOutcome) error
	Delete(id string) error
}

type SubProgramLearningOutcomeUsecase interface {
	GetAll() ([]SubProgramLearningOutcome, error)
	GetByID(id string) (*SubProgramLearningOutcome, error)
	Create(code string, descriptionThai string, descriptionEng string, programLearningOutcomeId string) error
	Update(id string, programLearningOutcome *SubProgramLearningOutcome) error
	Delete(id string) error
}
