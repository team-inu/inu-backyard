package entity

// [1] Info
type CourseInfo struct {
	Name      string   `json:"courseName"`
	Code      string   `json:"courseCode"`
	Lecturers []string `json:"lecturers"`
}

// [2] Summary
type CourseSummary struct {
	TeachingMethods []string `json:"teachingMethod"`
	OnlineTool      string   `json:"onlineTool"`
	Objectives      []string `json:"objectives"`
}

// [3.1] Tabee Outcome
type Assessment struct {
	AssessmentTask        string  `json:"assessmentTask"`
	PassingCriteria       float64 `json:"passingCriteria"`
	StudentPassPercentage float64 `json:"studentPassPercentage"`
}

type CourseOutcome struct {
	Name        string       `json:"name"`
	Assessments []Assessment `json:"assessments"`
}

type TabeeOutcome struct {
	Name              string          `json:"name"`
	CourseOutcomes    []CourseOutcome `json:"courseOutcomes"`
	MinimumPercentage float64         `json:"minimumPercentage"`
}

// [3.2] Grade Distribution
type GradeFrequency struct {
	Name       string  `json:"name"`
	GradeScore float64 `json:"gradeScore"`
	Frequency  int     `json:"frequency"`
}

type ScoreFrequency struct {
	Score     int `json:"score"`
	Frequency int `json:"frequency"`
}

type GradeDistribution struct {
	StudentAmount    int              `json:"studentAmount"`
	GPA              float64          `json:"GPA"`
	GradeFrequencies []GradeFrequency `json:"gradeFrequencies"`
	ScoreFrequencies []ScoreFrequency `json:"scoreFrequencies"`
}

// [3] Result
type CourseResult struct {
	TabeeOutcomes     []TabeeOutcome    `json:"tabeeOutcomes"`
	GradeDistribution GradeDistribution `json:"gradeDistribution"`
}

// [4.1] SubjectComments
type Subject struct {
	CourseName string `json:"courseName"`
	Comment    string `json:"comments"`
}

type SubjectComments struct {
	UpstreamSubjects   []Subject `json:"upstreamSubjects"`
	DownstreamSubjects []Subject `json:"downstreamSubjects"`
	Other              string    `json:"other"`
}

// [4] Development
type CourseDevelopment struct {
	Plans           []string        `json:"plans"`
	DoAndChecks     []string        `json:"doAndChecks"`
	Acts            []string        `json:"acts"`
	SubjectComments SubjectComments `json:"subjectComments"`
	OtherComment    string          `json:"otherComment"`
}

// Course Portfolio
type CoursePortfolio struct {
	CourseInfo        CourseInfo        `json:"info"`
	CourseSummary     CourseSummary     `json:"summary"`
	CourseResult      CourseResult      `json:"result"`
	CourseDevelopment CourseDevelopment `json:"development"`
}

type AssignmentPercentage struct {
	AssignmentId            string `gorm:"column:a_id"`
	Name                    string
	ExpectedScorePercentage float64
	PassingPercentage       float64
	CourseLearningOutcomeId string `gorm:"column:c_id"`
}

type PoPercentage struct {
	PassingPercentage float64
	ProgramOutcomeId  string `gorm:"column:p_id"`
}

type CloPercentage struct {
	PassingPercentage       float64
	CourseLearningOutcomeId string `gorm:"column:c_id"`
}

type StudentData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	StudentId string `json:"studentId"`
	Pass      bool   `json:"pass"`
}

type CloPassingStudent struct {
	CourseLearningOutcomeId string        `json:"courseLearningOutcomeId"`
	Students                []StudentData `json:"students"`
}

type CloPassingStudentGorm struct {
	FirstName               string
	LastName                string
	StudentId               string
	Pass                    bool
	CourseLearningOutcomeId string `gorm:"column:clo_id"`
}

type PloData struct {
	Id              string `json:"id"`
	Code            string `json:"code"`
	DescriptionThai string `json:"descriptionThai"`
	ProgramYear     int    `json:"programYear"`
	Pass            bool   `json:"pass"`
}

type PloPassingStudentGorm struct {
	Code                     string
	DescriptionThai          string
	ProgramYear              int
	StudentId                string
	Pass                     bool
	ProgramLearningOutcomeId string `gorm:"column:plo_id"`
}

type PoData struct {
	Id   string `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
	Pass bool   `json:"pass"`
}

type PoPassingStudentGorm struct {
	Code             string
	Name             string
	StudentId        string
	Pass             bool
	ProgramOutcomeId string `gorm:"column:p_id"`
}

type StudentOutcomeStatus struct {
	StudentId               string    `json:"studentId"`
	ProgramLearningOutcomes []PloData `json:"programLearningOutcomes"`
	ProgramOutcomes         []PoData  `json:"programOutcomes"`
}

type CoursePortfolioRepository interface {
	EvaluatePassingAssignmentPercentage(courseId string) ([]AssignmentPercentage, error)
	EvaluatePassingPoPercentage(courseId string) ([]PoPercentage, error)
	EvaluatePassingCloPercentage(courseId string) ([]CloPercentage, error)
	EvaluatePassingCloStudents(courseId string) ([]CloPassingStudentGorm, error)
	EvaluatePassingPloStudents(courseId string) ([]PloPassingStudentGorm, error)
	EvaluatePassingPoStudents(courseId string) ([]PoPassingStudentGorm, error)
}

type CoursePortfolioUseCase interface {
	Generate(courseId string) (*CoursePortfolio, error)
	CalculateGradeDistribution(courseId string) (*GradeDistribution, error)
	EvaluateTabeeOutcomes(courseId string) ([]TabeeOutcome, error)
	GetCloPassingStudentsByCourseId(courseId string) ([]CloPassingStudent, error)
	GetStudentOutcomesStatusByCourseId(courseId string) ([]StudentOutcomeStatus, error)
}
