package entity

type ProgramLearningOutcome struct {
	Id              string `json:"id" gorm:"primaryKey;type:char(255)"`
	Code            string `json:"code"`
	DescriptionThai string `json:"descriptionThai"`
	DescriptionEng  string `json:"descriptionEng"`
	ProgramYear     int    `json:"programYear"`
	ProgrammeId     string `json:"programmeId"`

	Programme Programme
}

type CrateProgramLearningOutcomeDto struct {
	Code            string `validate:"required"`
	DescriptionThai string `validate:"required"`
	DescriptionEng  string `validate:"required"`
	ProgramYear     int    `validate:"required"`
	ProgrammeName   string `validate:"required"`
}

type ProgramLearningOutcomeRepository interface {
	GetAll() ([]ProgramLearningOutcome, error)
	GetById(id string) (*ProgramLearningOutcome, error)
	Create(programLearningOutcome *ProgramLearningOutcome) error
	CreateMany(programLearningOutcome []ProgramLearningOutcome) error
	Update(id string, programLearningOutcome *ProgramLearningOutcome) error
	Delete(id string) error
}

type ProgramLearningOutcomeUseCase interface {
	GetAll() ([]ProgramLearningOutcome, error)
	GetById(id string) (*ProgramLearningOutcome, error)
	Create(dto []CrateProgramLearningOutcomeDto) error
	Update(id string, programLearningOutcome *ProgramLearningOutcome) error
	Delete(id string) error
}
