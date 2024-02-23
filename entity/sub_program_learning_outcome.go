package entity

type SubProgramLearningOutcome struct {
	Id                       string `json:"id" gorm:"primaryKey;type:char(255)"`
	Code                     string `json:"code"`
	DescriptionThai          string `json:"descriptionThai"`
	DescriptionEng           string `json:"descriptionEng"`
	ProgramLearningOutcomeId string `json:"programLearningOutcomeId"`

	CourseLearningOutcome  []*CourseLearningOutcome `gorm:"many2many:clo_subplo"`
	ProgramLearningOutcome ProgramLearningOutcome
}

type SubProgramLearningOutcomeRepository interface {
	GetAll() ([]SubProgramLearningOutcome, error)
	GetById(id string) (*SubProgramLearningOutcome, error)
	Create(programLearningOutcome *SubProgramLearningOutcome) error
	Update(id string, programLearningOutcome *SubProgramLearningOutcome) error
	Delete(id string) error
	FilterExisted(ids []string) ([]string, error)
}

type SubProgramLearningOutcomeUseCase interface {
	GetAll() ([]SubProgramLearningOutcome, error)
	GetById(id string) (*SubProgramLearningOutcome, error)
	Create(code string, descriptionThai string, descriptionEng string, programLearningOutcomeId string) error
	Update(id string, programLearningOutcome *SubProgramLearningOutcome) error
	Delete(id string) error
	FilterNonExisted(ids []string) ([]string, error)
}
