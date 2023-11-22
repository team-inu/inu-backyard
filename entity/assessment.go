package entity

type Assessment struct {
	ID                      string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Score                   int    `json:"score"`
	Weight                  int    `json:"weight"`
	CourseLearningOutcomeID string `json:"courseLearningOutcomeID"`

	CourseLearningOutcome CourseLearningOutcome ` gorm:"references:Code"`
}

type AssessmentRepository interface {
	GetByID(id string) (*Assessment, error)
	GetByParams(params *Assessment, limit int, offset int) ([]Assessment, error)
	Create(assessment *Assessment) error
	CreateMany(assessment []Assessment) error
	Update(id string, assessment *Assessment) error
	Delete(id string) error
}

type AssessmentUseCase interface {
	GetByID(id string) (*Assessment, error)
	GetByParams(params *Assessment, limit int, offset int) ([]Assessment, error)
	GetByCourseID(courseID string, limit int, offset int) ([]Assessment, error)
	Create(assessment *Assessment) error
	CreateMany(assessment []Assessment) error
	Update(id string, assessment *Assessment) error
	Delete(id string) error
}
