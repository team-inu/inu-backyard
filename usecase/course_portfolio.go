package usecase

import (
	"fmt"
	"math"

	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type coursePortfolioUseCase struct {
	CoursePortfolioRepository    entity.CoursePortfolioRepository
	CourseUseCase                entity.CourseUseCase
	UserUseCase                  entity.UserUseCase
	EnrollmentUseCase            entity.EnrollmentUseCase
	AssignmentUseCase            entity.AssignmentUseCase
	ScoreUseCase                 entity.ScoreUseCase
	CourseLearningOutcomeUseCase entity.CourseLearningOutcomeUseCase
	CourseStreamUseCase          entity.CourseStreamsUseCase
}

func NewCoursePortfolioUseCase(
	coursePortfolioRepository entity.CoursePortfolioRepository,
	courseUseCase entity.CourseUseCase,
	userUseCase entity.UserUseCase,
	enrollmentUseCase entity.EnrollmentUseCase,
	assignmentUseCase entity.AssignmentUseCase,
	scoreUseCase entity.ScoreUseCase,
	courseLearningOutcomeUseCase entity.CourseLearningOutcomeUseCase,
	courseStreamUseCase entity.CourseStreamsUseCase,
) entity.CoursePortfolioUseCase {
	return &coursePortfolioUseCase{
		CoursePortfolioRepository:    coursePortfolioRepository,
		CourseUseCase:                courseUseCase,
		UserUseCase:                  userUseCase,
		EnrollmentUseCase:            enrollmentUseCase,
		AssignmentUseCase:            assignmentUseCase,
		ScoreUseCase:                 scoreUseCase,
		CourseLearningOutcomeUseCase: courseLearningOutcomeUseCase,
		CourseStreamUseCase:          courseStreamUseCase,
	}
}

func (u coursePortfolioUseCase) Generate(courseId string) (*entity.CoursePortfolio, error) {
	course, err := u.CourseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course id %s while generate course portfolio", courseId, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "course id %s not found while generate course portfolio", courseId)
	}

	lecturer, err := u.UserUseCase.GetById(course.UserId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get lecturer id %s while generate course portfolio", course.UserId, err)
	} else if lecturer == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "user id %s not found while generate course portfolio", course.UserId)
	}

	courseInfo := entity.CourseInfo{
		Name:      course.Name,
		Code:      course.Code,
		Lecturers: []string{fmt.Sprintf("%s %s", lecturer.FirstName, lecturer.LastName)},
	}

	gradeDistribution, err := u.CalculateGradeDistribution(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot calculate grade distribution while generate course portfolio", err)
	}

	tabeeOutcomes, err := u.EvaluateTabeeOutcomes(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate tabee outcomes while generate course portfolio", err)
	}

	courseResult := entity.CourseResult{
		GradeDistribution: *gradeDistribution,
		TabeeOutcomes:     tabeeOutcomes,
	}

	courseStreams, err := u.CourseStreamUseCase.GetByTargetCourseId(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course streams while generate course portfolio", err)
	}

	upstreamSubject := make([]entity.Subject, 0)
	downStreamSubject := make([]entity.Subject, 0)

	for _, stream := range courseStreams {
		switch stream.StreamType {
		case entity.DownCourseStreamType:
			fmt.Println(stream.FromCourse)
			upstreamSubject = append(upstreamSubject, entity.Subject{
				CourseName: fmt.Sprintf("%s %s", stream.FromCourse.Code, stream.FromCourse.Name),
				Comment:    stream.Comment,
			})

		case entity.UpCourseStreamType:
			downStreamSubject = append(downStreamSubject, entity.Subject{
				CourseName: fmt.Sprintf("%s %s", stream.TargetCourse.Code, stream.TargetCourse.Name),
				Comment:    stream.Comment,
			})

		}
	}

	courseDevelopment := entity.CourseDevelopment{
		Plans:       make([]string, 0),
		DoAndChecks: make([]string, 0),
		Acts:        make([]string, 0),
		SubjectComments: entity.SubjectComments{
			UpstreamSubjects:   upstreamSubject,
			DownstreamSubjects: downStreamSubject,
		},
	}

	courseSummary := entity.CourseSummary{
		TeachingMethods: make([]string, 0),
		Objectives:      make([]string, 0),
	}

	coursePortfolio := &entity.CoursePortfolio{
		CourseInfo:        courseInfo,
		CourseResult:      courseResult,
		CourseSummary:     courseSummary,
		CourseDevelopment: courseDevelopment,
	}

	return coursePortfolio, nil
}

func (u coursePortfolioUseCase) CalculateGradeDistribution(courseId string) (*entity.GradeDistribution, error) {
	type studentScore struct {
		studentId string
		score     float64
		weight    int
	}

	course, err := u.CourseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course by id %s while calculate grade distribution", courseId, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "course id %s not found while calculate grade distribution", courseId)
	}

	assignments, err := u.AssignmentUseCase.GetByCourseId(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get assignments by course id %s while calculate grade distribution", courseId, err)
	}

	cumulativeWeight := 0
	for _, assignment := range assignments {
		cumulativeWeight += assignment.Weight
	}

	cumulativeWeightedMaxScore := 0
	for _, assignment := range assignments {
		cumulativeWeightedMaxScore += assignment.MaxScore * assignment.Weight / cumulativeWeight
	}

	studentScoresByAssignmentId := make(map[string][]studentScore, 0)
	for _, assignment := range assignments {
		scores, err := u.ScoreUseCase.GetByAssignmentId(assignment.Id)
		if err != nil {
			return nil, errs.New(errs.SameCode, "cannot get scores by assignment id %s while calculate grade distribution", assignment.Id, err)
		}

		for _, score := range scores.Scores {
			studentScoresByAssignmentId[assignment.Id] = append(studentScoresByAssignmentId[assignment.Id], studentScore{
				studentId: score.StudentId,
				score:     score.Score,
				weight:    assignment.Weight,
			})
		}
	}

	studentScoreByStudentId := make(map[string]float64)

	//calculate student scores
	for _, assignmentScores := range studentScoresByAssignmentId {
		for _, score := range assignmentScores {
			studentScoreByStudentId[score.studentId] += (score.score * float64(score.weight) / float64(cumulativeWeight))
		}
	}

	for studentId, studentScore := range studentScoreByStudentId {
		studentScoreByStudentId[studentId] = studentScore * 100 / float64(cumulativeWeightedMaxScore)
	}

	frequencyByScore := make(map[int]int)
	for _, score := range studentScoreByStudentId {
		roundedScore := int(math.Round(score))
		frequencyByScore[roundedScore]++
	}

	scoreFrequencies := make([]entity.ScoreFrequency, 0, len(frequencyByScore))
	for score, frequency := range frequencyByScore {
		scoreFrequencies = append(scoreFrequencies, entity.ScoreFrequency{
			Score:     score,
			Frequency: frequency,
		})
	}

	weightedCriteriaGrade := course.CriteriaGrade.CalculateCriteriaWeight(float64(cumulativeWeightedMaxScore))

	frequenciesByGrade := make(map[string]int)
	for _, studentScore := range studentScoreByStudentId {
		switch {
		case studentScore >= weightedCriteriaGrade.A:
			frequenciesByGrade["A"] += 1
		case studentScore >= weightedCriteriaGrade.BP:
			frequenciesByGrade["BP"] += 1
		case studentScore >= weightedCriteriaGrade.B:
			frequenciesByGrade["B"] += 1
		case studentScore >= weightedCriteriaGrade.CP:
			frequenciesByGrade["CP"] += 1
		case studentScore >= weightedCriteriaGrade.C:
			frequenciesByGrade["C"] += 1
		case studentScore >= weightedCriteriaGrade.DP:
			frequenciesByGrade["DP"] += 1
		case studentScore >= weightedCriteriaGrade.D:
			frequenciesByGrade["D"] += 1
		default:
			frequenciesByGrade["F"] += 1
		}
	}

	gradeFrequencies := []entity.GradeFrequency{
		{
			Name:       "A",
			GradeScore: weightedCriteriaGrade.A,
			Frequency:  frequenciesByGrade["A"],
		},
		{
			Name:       "BP",
			GradeScore: weightedCriteriaGrade.BP,
			Frequency:  frequenciesByGrade["BP"],
		},
		{
			Name:       "B",
			GradeScore: weightedCriteriaGrade.B,
			Frequency:  frequenciesByGrade["B"],
		},
		{
			Name:       "CP",
			GradeScore: weightedCriteriaGrade.CP,
			Frequency:  frequenciesByGrade["CP"],
		},
		{
			Name:       "C",
			GradeScore: weightedCriteriaGrade.C,
			Frequency:  frequenciesByGrade["C"],
		},
		{
			Name:       "DP",
			GradeScore: weightedCriteriaGrade.DP,
			Frequency:  frequenciesByGrade["DP"],
		},
		{
			Name:       "D",
			GradeScore: weightedCriteriaGrade.D,
			Frequency:  frequenciesByGrade["D"],
		},
		{
			Name:       "F",
			GradeScore: weightedCriteriaGrade.F,
			Frequency:  frequenciesByGrade["F"],
		},
	}

	studentAmount := len(studentScoreByStudentId)

	totalStudentGPA := 0.0
	for grade, frequency := range frequenciesByGrade {
		totalStudentGPA += float64(frequency) * course.CriteriaGrade.GradeToGPA(grade)
	}

	gpa := 0.0
	if totalStudentGPA != 0.0 {
		gpa = totalStudentGPA / float64(studentAmount)
	}

	return &entity.GradeDistribution{
		GradeFrequencies: gradeFrequencies,
		StudentAmount:    studentAmount,
		GPA:              gpa,
		ScoreFrequencies: scoreFrequencies,
	}, nil
}

func (u coursePortfolioUseCase) EvaluateTabeeOutcomes(courseId string) ([]entity.TabeeOutcome, error) {
	assignmentPercentages, err := u.CoursePortfolioRepository.EvaluatePassingAssignmentPercentage(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate passing assignment percentage by course id %s while evaluate tabee outcome", courseId, err)
	}

	assessmentsByCloId := make(map[string][]entity.Assessment, len(assignmentPercentages))
	for _, assignmentPercentage := range assignmentPercentages {

		cloId := assignmentPercentage.CourseLearningOutcomeId

		assessmentsByCloId[cloId] = append(assessmentsByCloId[cloId], entity.Assessment{
			AssessmentTask:        assignmentPercentage.Name,
			PassingCriteria:       assignmentPercentage.ExpectedScorePercentage,
			StudentPassPercentage: assignmentPercentage.PassingPercentage,
		})
	}

	clos, err := u.CourseLearningOutcomeUseCase.GetByCourseId(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get clo while evaluate tabee outcome", err)
	}

	courseOutcomeByPoId := make(map[string][]entity.CourseOutcome, 0)
	for _, clo := range clos {
		courseOutcomeByPoId[clo.ProgramOutcomeId] = append(courseOutcomeByPoId[clo.ProgramOutcomeId], entity.CourseOutcome{
			Name:        clo.Description,
			Assessments: assessmentsByCloId[clo.Id],
		})
	}

	passingPoPercentages, err := u.CoursePortfolioRepository.EvaluatePassingPoPercentage(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate passing po percentage by course id %s while evaluate tabee outcome", courseId, err)
	}

	passingPoPercentageByPoId := make(map[string]float64, len(passingPoPercentages))
	for _, passingPoPercentage := range passingPoPercentages {
		passingPoPercentageByPoId[passingPoPercentage.ProgramOutcomeId] = passingPoPercentage.PassingPercentage
	}

	tabeeOutcomesByPoId := make(map[string][]entity.TabeeOutcome, 0)
	for _, clo := range clos {
		tabeeOutcomesByPoId[clo.ProgramOutcomeId] = append(tabeeOutcomesByPoId[clo.ProgramOutcomeId], entity.TabeeOutcome{
			Name:              clo.ProgramOutcomeName,
			CourseOutcomes:    courseOutcomeByPoId[clo.ProgramOutcomeId],
			MinimumPercentage: passingPoPercentageByPoId[clo.ProgramOutcomeId],
		})
	}

	tabeeOutcomes := make([]entity.TabeeOutcome, 0, len(tabeeOutcomesByPoId))
	for _, tabeeOutcome := range tabeeOutcomesByPoId {
		tabeeOutcomes = append(tabeeOutcomes, tabeeOutcome...)
	}

	return tabeeOutcomes, nil
}

func (u coursePortfolioUseCase) GetCloPassingStudentsByCourseId(courseId string) ([]entity.CloPassingStudent, error) {
	course, err := u.CourseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course id %s while getting clo passing students", course, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "course id %s not found while getting clo passing students", courseId, err)
	}

	records, err := u.CoursePortfolioRepository.EvaluatePassingCloStudents(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate passing clo student by course id %s", courseId, err)
	}

	closMap := make(map[string][]entity.StudentData)

	for _, record := range records {
		closMap[record.CourseLearningOutcomeId] = append(closMap[record.CourseLearningOutcomeId], entity.StudentData{
			FirstName: record.FirstName,
			LastName:  record.LastName,
			StudentId: record.StudentId,
			Pass:      record.Pass,
		})
	}

	clos := make([]entity.CloPassingStudent, 0)

	for cloId := range closMap {
		clos = append(clos, entity.CloPassingStudent{
			CourseLearningOutcomeId: cloId,
			Students:                closMap[cloId],
		})
	}

	return clos, nil
}
