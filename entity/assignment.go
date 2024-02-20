package entity

type Assignment struct {
	ID                      string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name                    string `json:"name"`
	Description             string `json:"description"`
	Score                   int    `json:"score"`
	Weight                  int    `json:"weight"`
	CourseLearningOutcomeID string `json:"courseLearningOutcomeID"`

	CourseLearningOutcome CourseLearningOutcome ` gorm:"references:Code"`
}

type AssignmentRepository interface {
	GetByID(id string) (*Assignment, error)
	GetByParams(params *Assignment, limit int, offset int) ([]Assignment, error)
	Create(assignment *Assignment) error
	CreateMany(assignment []Assignment) error
	Update(id string, assignment *Assignment) error
	Delete(id string) error
}

type AssignmentUseCase interface {
	GetByID(id string) (*Assignment, error)
	GetByParams(params *Assignment, limit int, offset int) ([]Assignment, error)
	GetByCourseID(courseID string, limit int, offset int) ([]Assignment, error)
	Create(assignment *Assignment) error
	CreateMany(assignment []Assignment) error
	Update(id string, assignment *Assignment) error
	Delete(id string) error
}
