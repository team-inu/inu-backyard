package entity

type CourseLearningOutcome struct {
	ID                                  string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Code                                string  `json:"code"`
	Description                         string  `json:"description"`
	ExpectedPassingAssignmentPercentage float64 `json:"expectedPassingAssignment"`
	ExpectedScorePercentage             float64 `json:"expectedScorePercentage"`
	ExpectedPassingStudentPercentage    float64 `json:"expectedPassingStudentPercentage"`
	SubProgramLearningOutcomeID         string  `json:"subProgramLearningOutcomeID"`
	ProgramOutcomeID                    string  `json:"programOutcomeID"`
	CourseID                            string  `json:"courseID"`
	Status                              string  `json:"status"`

	SubProgramLearningOutcomes []*SubProgramLearningOutcome `gorm:"many2many:clo_subplo"`
	Assignments                []*Assignment                `gorm:"many2many:clo_assignment"`
	ProgramOutcome             ProgramOutcome
	Course                     Course
}

type CourseLearningOutcomeRepository interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetByID(id string) (*CourseLearningOutcome, error)
	GetByCourseID(courseID string) ([]CourseLearningOutcome, error)
	Create(courseLearningOutcome *CourseLearningOutcome) error
	Update(id string, courseLearningOutcome *CourseLearningOutcome) error
	Delete(id string) error
}

type CourseLearningOutcomeUsecase interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetByID(id string) (*CourseLearningOutcome, error)
	GetByCourseID(courseID string) ([]CourseLearningOutcome, error)
	Create(code string, description string, weight int, subProgramLearningOutcomeID string, programOutcomeID string, courseID string, status string) error
	Update(id string, courseLearningOutcome *CourseLearningOutcome) error
	Delete(id string) error
}
