package entity

type SubProgramLearningOutcome struct {
	ID                       string `json:"id" gorm:"primaryKey;type:char(255)"`
	Code                     string `json:"code"`
	DescriptionThai          string `json:"descriptionThai"`
	DescriptionEng           string `json:"descriptionEng"`
	ProgramLearningOutcomeID string `json:"programLearningOutcomeID"`

	CourseLearningOutcome  []*CourseLearningOutcome `gorm:"many2many:clo_subplo"`
	ProgramLearningOutcome ProgramLearningOutcome
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
	Create(code string, descriptionThai string, descriptionEng string, programLearningOutcomeID string) error
	Update(id string, programLearningOutcome *SubProgramLearningOutcome) error
	Delete(id string) error
}
