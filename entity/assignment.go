package entity

type AssignmentGroup struct {
	Id       string  `json:"id" gorm:"primaryKey;type:char(255)"`
	Name     string  `json:"name"`
	CourseId string  `json:"courseId"`
	Weight   float64 `json:"weight"`

	Assignments []Assignment `gorm:"foreignKey:AssignmentGroupId" json:"assignments,omitempty"`

	Course *Course `json:",omitempty"`
}

type Assignment struct {
	Id                               string                   `json:"id" gorm:"primaryKey;type:char(255)"`
	Name                             string                   `json:"name"`
	Description                      string                   `json:"description"`
	MaxScore                         float64                  `json:"maxScore"`
	ExpectedScorePercentage          float64                  `json:"expectedScorePercentage"`
	ExpectedPassingStudentPercentage float64                  `json:"expectedPassingStudentPercentage"`
	IsIncludedInClo                  *bool                    `json:"isIncludedInClo"`
	AssignmentGroupId                string                   `json:"assignmentGroupId" gorm:"not null"`
	CourseId                         string                   `json:"courseId" gorm:"->;-:migration"`
	CourseLearningOutcomes           []*CourseLearningOutcome `gorm:"many2many:clo_assignment" json:"courseLearningOutcomes"`
}

func GenerateGroupByAssignmentId(assignmentGroups []AssignmentGroup, assignments []Assignment) map[string]*AssignmentGroup {
	weightByAssignmentGroupId := make(map[string]*AssignmentGroup, len(assignmentGroups))
	for _, assignmentGroup := range assignmentGroups {
		weightByAssignmentGroupId[assignmentGroup.Id] = &assignmentGroup
	}

	weightByAssignmentId := make(map[string]*AssignmentGroup, len(assignments))
	for _, assignment := range assignments {
		assignmentGroup, ok := weightByAssignmentGroupId[assignment.AssignmentGroupId]
		if !ok {
			continue
		}

		weightByAssignmentId[assignment.Id] = assignmentGroup
	}

	return weightByAssignmentId
}

type AssignmentRepository interface {
	GetById(id string) (*Assignment, error)
	GetByCourseId(courseId string) ([]Assignment, error)
	GetByGroupId(groupId string) ([]Assignment, error)
	GetPassingStudentPercentage(assignmentId string) (float64, error)
	Create(assignment *Assignment) error
	CreateMany(assignment []Assignment) error
	Update(id string, assignment *Assignment) error
	Delete(id string) error

	CreateLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId []string) error
	DeleteLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId string) error

	GetGroupByQuery(query AssignmentGroup, withAssignment bool) ([]AssignmentGroup, error)
	GetGroupByGroupId(assignmentGroupId string) (*AssignmentGroup, error)
	CreateGroup(assignmentGroup *AssignmentGroup) error
	UpdateGroup(assignmentGroupId string, assignmentGroup *AssignmentGroup) error
	DeleteGroup(assignmentGroupId string) error
}

type AssignmentUseCase interface {
	GetById(id string) (*Assignment, error)
	GetByCourseId(courseId string) ([]Assignment, error)
	GetByGroupId(assignmentGroupId string) ([]Assignment, error)
	GetPassingStudentPercentage(assignmentId string) (float64, error)
	Create(assignmentGroupId string, name string, description string, maxScore float64, expectedScorePercentage float64, expectedPassingStudentPercentage float64, courseLearningOutcomeIds []string, isIncludedInClo bool) error
	Update(id string, name string, description string, maxScore float64, expectedScorePercentage float64, expectedPassingStudentPercentage float64, isIncludedInClo bool) error
	Delete(id string) error

	CreateLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId []string) error
	DeleteLinkCourseLearningOutcome(assignmentId string, courseLearningOutcomeId string) error

	GetGroupByGroupId(assignmentGroupId string) (*AssignmentGroup, error)
	GetGroupByCourseId(courseId string, withAssignment bool) ([]AssignmentGroup, error)
	CreateGroup(name string, courseId string, weight float64) error
	UpdateGroup(assignmentGroupId string, name string, weight float64) error
	DeleteGroup(assignmentGroupId string) error
}
