package entity

// [1] Info
type CourseInfo struct {
	Name      string   `json:"courseName"`
	Code      string   `json:"courseCode"`
	Lecturers []string `json:"lecturers"`
}

// [2] Summary
type CourseSummary struct {
	TeachingMethods []string
	OnlineTool      string
	Objectives      []string
}

// [3.1] Tabee Outcome
type Assessment struct {
	AssessmentTask        string
	PassingCriteria       string
	StudentPassPercentage string
}

type CourseOutcome struct {
	Name        string
	Assessments []Assessment
}

type TabeeOutcome struct {
	Name              string
	CourseOutcomes    []CourseOutcome
	MinimumPercentage string
}

// [3.2] Grade Distribution
type GradeFrequency struct {
	Name       string
	GradeScore string
	Frequency  string
}

type GradeDistribution struct {
	StudentAmount  string
	GPA            string
	GradeFrequency GradeFrequency
}

// [3] Result
type CourseResult struct {
	TabeeOutcomes     []TabeeOutcome
	GradeDistribution GradeDistribution
}

// [4.1] SubjectComments
type Subject struct {
	CourseName string
	Comment    string
}

type SubjectComments struct {
	UpstreamSubjects   []Subject
	DownstreamSubjects []Subject
}

// [4] Development
type CourseDevelopment struct {
	Plans           []string
	DoAndChecks     []string
	Acts            []string
	SubjectComments SubjectComments
	OtherComment    string
}

// Course Portfolio
type CoursePortfolio struct {
	CourseInfo        CourseInfo
	CourseSummary     CourseSummary
	CourseResult      CourseResult
	CourseDevelopment CourseDevelopment
}

type CoursePortfolioUseCase interface {
	Generate(courseId string) (*CoursePortfolio, error)
	CalculateGradeDistribution() (*GradeDistribution, error)
	EvaluateTabeeOutcomes() ([]TabeeOutcome, error)
}
