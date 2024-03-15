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
	AssessmentTask        string `json:"assessmentTask"`
	PassingCriteria       string `json:"passingCriteria"`
	StudentPassPercentage string `json:"studentPassPercentage"`
}

type CourseOutcome struct {
	Name        string       `json:"name"`
	Assessments []Assessment `json:"assessments"`
}

type TabeeOutcome struct {
	Name              string          `json:"name"`
	CourseOutcomes    []CourseOutcome `json:"courseOutcomes"`
	MinimumPercentage string          `json:"minimumPercentage"`
}

// [3.2] Grade Distribution
type GradeFrequency struct {
	Name       string  `json:"name"`
	GradeScore float64 `json:"gradeScore"`
	Frequency  int     `json:"frequency"`
}

type GradeDistribution struct {
	StudentAmount    int              `json:"studentAmount"`
	GPA              float64          `json:"GPA"`
	GradeFrequencies []GradeFrequency `json:"gradeFrequencies"`
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

type CoursePortfolioUseCase interface {
	Generate(courseId string) (*CoursePortfolio, error)
	CalculateGradeDistribution(courseId string) (*GradeDistribution, error)
	EvaluateTabeeOutcomes() ([]TabeeOutcome, error)
}
