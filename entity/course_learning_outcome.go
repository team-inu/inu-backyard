package entity

type CourseLearningOutcome struct {
	Id                                  string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Code                                string  `json:"code"`
	Description                         string  `json:"description"`
	ExpectedPassingAssignmentPercentage float64 `json:"expectedPassingAssignment"`
	ExpectedScorePercentage             float64 `json:"expectedScorePercentage"`
	ExpectedPassingStudentPercentage    float64 `json:"expectedPassingStudentPercentage"`
	SubProgramLearningOutcomeId         string  `json:"subProgramLearningOutcomeId"`
	ProgramOutcomeId                    string  `json:"programOutcomeId"`
	CourseId                            string  `json:"courseId"`
	Status                              string  `json:"status"`

	SubProgramLearningOutcomes []*SubProgramLearningOutcome `gorm:"many2many:clo_subplo"`
	Assignments                []*Assignment                `gorm:"many2many:clo_assignment"`
	ProgramOutcome             ProgramOutcome
	Course                     Course
}

type CourseLearningOutcomeRepository interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetById(id string) (*CourseLearningOutcome, error)
	GetByCourseId(courseId string) ([]CourseLearningOutcome, error)
	Create(courseLearningOutcome *CourseLearningOutcome) error
	Update(id string, courseLearningOutcome *CourseLearningOutcome) error
	Delete(id string) error
}

type CourseLearningOutcomeUsecase interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetById(id string) (*CourseLearningOutcome, error)
	GetByCourseId(courseId string) ([]CourseLearningOutcome, error)
	Create(code string, description string, weight int, subProgramLearningOutcomeId string, programOutcomeId string, courseId string, status string) error
	Update(id string, courseLearningOutcome *CourseLearningOutcome) error
	Delete(id string) error
}
