package entity

type CourseLearningOutcome struct {
	ID                          string `json:"id" gorm:"primaryKey;type:char(255)"`
	Code                        string `json:"code"`
	Description                 string `json:"description"`
	Weight                      int    `json:"weight"`
	SubProgramLearningOutcomeID string `json:"subProgramLearningOutcomeId"`
	ProgramOutcomeID            string `json:"programOutcomeId"`
	CourseId                    string `json:"courseId"`
	Status                      string `json:"status"`

	SubProgramLearningOutcome SubProgramLearningOutcome `gorm:"references:Code"`
	ProgramOutcome            ProgramOutcome            `gorm:"references:Code"`
	// Course                 Course
}

type CourseLearningOutcomeRepository interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetByID(id string) (*CourseLearningOutcome, error)
	GetByCourseID(courseId string) ([]CourseLearningOutcome, error)
	Create(courseLearningOutcome *CourseLearningOutcome) error
	Update(id string, courseLearningOutcome *CourseLearningOutcome) error
	Delete(id string) error
}

type CourseLearningOutcomeUsecase interface {
	GetAll() ([]CourseLearningOutcome, error)
	GetByID(id string) (*CourseLearningOutcome, error)
	GetByCourseID(courseId string) ([]CourseLearningOutcome, error)
	Create(code string, description string, weight int, subProgramLearningOutcomeId string, programOutcomeId string, courseId string, status string) error
	Update(id string, courseLearningOutcome *CourseLearningOutcome) error
	Delete(id string) error
}
