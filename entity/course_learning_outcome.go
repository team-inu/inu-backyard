package entity

type CourseLearningOutcome struct {
	Id                                  string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Code                                string  `json:"code"`
	Description                         string  `json:"description"`
	ExpectedPassingAssignmentPercentage float64 `json:"expectedPassingAssignmentPercentage"`
	ExpectedPassingStudentPercentage    float64 `json:"expectedPassingStudentPercentage"`
	Status                              string  `json:"status"`
	ProgramOutcomeId                    string  `json:"programOutcomeId"`
	CourseId                            string  `json:"courseId"`

	SubProgramLearningOutcomes []*SubProgramLearningOutcome `gorm:"many2many:clo_subplo" json:"subProgramLearningOutcomes"`
	Assignments                []*Assignment                `gorm:"many2many:clo_assignment" json:"-"`
	ProgramOutcome             ProgramOutcome               `json:"-"`
	Course                     Course                       `json:"-"`
}

type CreateCourseLearningOutcomeDto struct {
	Code                                string
	Description                         string
	ExpectedPassingAssignmentPercentage float64
	ExpectedPassingStudentPercentage    float64
	Status                              string
	SubProgramLearningOutcomeIds        []string
	ProgramOutcomeId                    string
	CourseId                            string
}

type UpdateCourseLeaningOutcomeDto struct {
	Code                                string
	Description                         string
	ExpectedPassingAssignmentPercentage float64
	ExpectedPassingStudentPercentage    float64
	Status                              string
	ProgramOutcomeId                    string
}

type CourseLearningOutcomeRepository interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetById(id string) (*CourseLearningOutcome, error)
	GetByCourseId(courseId string) ([]CourseLearningOutcome, error)
	Create(courseLearningOutcome *CourseLearningOutcome) error
	CreateLinkSubProgramLearningOutcome(id string, subProgramLearningOutcomeId []string) error
	Update(id string, courseLearningOutcome *CourseLearningOutcome) error
	Delete(id string) error
	DeleteLinkSubProgramLearningOutcome(id string, subProgramLearningOutcomeId string) error
	FilterExisted(ids []string) ([]string, error)
}

type CourseLearningOutcomeUseCase interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetById(id string) (*CourseLearningOutcome, error)
	GetByCourseId(courseId string) ([]CourseLearningOutcome, error)
	Create(dto CreateCourseLearningOutcomeDto) error
	CreateLinkSubProgramLearningOutcome(id string, subProgramLearningOutcomeId []string) error
	Update(id string, dto UpdateCourseLeaningOutcomeDto) error
	Delete(id string) error
	DeleteLinkSubProgramLearningOutcome(id string, subProgramLearningOutcomeId string) error
	FilterNonExisted(ids []string) ([]string, error)
}
