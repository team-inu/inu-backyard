package entity

type AssignmentGroup struct {
	Id          string       `json:"id" gorm:"primaryKey;type:char(255)"`
	Name        string       `json:"name"`
	CourseId    string       `json:"courseId"`
	Assignments []Assignment `gorm:"foreignKey:AssignmentGroupId" json:"assignments"`

	Course *Course `json:",omitempty"`
}

type Assignment struct {
	Id                               string                   `json:"id" gorm:"primaryKey;type:char(255)"`
	Name                             string                   `json:"name"`
	Description                      string                   `json:"description"`
	MaxScore                         int                      `json:"maxScore"`
	Weight                           int                      `json:"weight"`
	ExpectedScorePercentage          float64                  `json:"expectedScorePercentage"`
	ExpectedPassingStudentPercentage float64                  `json:"expectedPassingStudentPercentage"`
	IsIncludedInClo                  *bool                    `json:"isIncludedInClo"`
	AssignmentGroupId                string                   `json:"assignmentGroupId" gorm:"not null"`
	CourseId                         string                   `json:"courseId" gorm:"->;-:migration"`
	CourseLearningOutcomes           []*CourseLearningOutcome `gorm:"many2many:clo_assignment" json:"courseLearningOutcomes"`
}

type AssignmentRepository interface {
	GetById(id string) (*Assignment, error)
	GetByParams(params *Assignment, limit int, offset int) ([]Assignment, error)
	GetByCourseId(courseId string) ([]Assignment, error)
	GetPassingStudentPercentage(assignmentId string) (float64, error)
	Create(assignment *Assignment) error
	CreateMany(assignment []Assignment) error
	Update(id string, assignment *Assignment) error
	Delete(id string) error

	CreateLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId []string) error
	DeleteLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId string) error

	GetGroupByQuery(query AssignmentGroup) ([]AssignmentGroup, error)
	GetGroupByGroupId(assignmentGroupId string) (*AssignmentGroup, error)
	CreateGroup(assignmentGroup *AssignmentGroup) error
	UpdateGroup(assignmentGroupId string, assignmentGroup *AssignmentGroup) error
}

type AssignmentUseCase interface {
	GetById(id string) (*Assignment, error)
	GetByParams(params *Assignment, limit int, offset int) ([]Assignment, error)
	GetByCourseId(courseId string) ([]Assignment, error)
	GetPassingStudentPercentage(assignmentId string) (float64, error)
	Create(assignmentGroupId string, name string, description string, maxScore int, weight int, expectedScorePercentage float64, expectedPassingStudentPercentage float64, courseLearningOutcomeIds []string, isIncludedInClo bool) error
	Update(id string, name string, description string, maxScore int, weight int, expectedScorePercentage float64, expectedPassingStudentPercentage float64, isIncludedInClo bool) error
	Delete(id string) error

	CreateLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId []string) error
	DeleteLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId string) error

	GetGroupByGroupId(assignmentGroupId string) (*AssignmentGroup, error)
	CreateGroup(name string, courseId string) error
	UpdateGroup(assignmentGroupId string, name string) error
}
