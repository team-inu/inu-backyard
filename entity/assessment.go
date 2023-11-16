package entity

type Assessment struct {
	ID                      string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Score                   int    `json:"score"`
	Weight                  int    `json:"weight"`
	CourseLearningOutcomeID string

	CourseLearningOutcome CourseLearningOutcome
}

type AssessmentRepository interface {
	GetAll() ([]Assessment, error)
	GetByID(id string) (*Assessment, error)
	Create(assessment *Assessment) error
	Update(assessment *Assessment) error
	Delete(id string) error
}

type AssessmentUseCase interface {
	GetAll() ([]Assessment, error)
	GetByID(id string) (*Assessment, error)
	GetByCourseLearningOutcomeID(courseLearningOutcomeID string) ([]Assessment, error)
	GetByCourseID(courseID string) ([]Assessment, error)
	Create(assessment *Assessment) error
	Update(assessment *Assessment) error
	Delete(id string) error
}
