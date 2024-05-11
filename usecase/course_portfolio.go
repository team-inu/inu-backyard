package usecase

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

// TODO: refactor (real)
type coursePortfolioUseCase struct {
	CoursePortfolioRepository    entity.CoursePortfolioRepository
	CourseUseCase                entity.CourseUseCase
	UserUseCase                  entity.UserUseCase
	EnrollmentUseCase            entity.EnrollmentUseCase
	AssignmentUseCase            entity.AssignmentUseCase
	ScoreUseCase                 entity.ScoreUseCase
	StudentUseCase               entity.StudentUseCase
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
	studentUscase entity.StudentUseCase,
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
		StudentUseCase:               studentUscase,
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

	closWithPos, err := u.CourseLearningOutcomeUseCase.GetByCourseId(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get clo while evaluate tabee outcome", err)
	}

	generateOutcome := func(cloWithPos []entity.CourseLearningOutcomeWithPO) ([]entity.NestedOutcome, []entity.Outcome, []entity.Outcome) {
		addedSubPlo := make(map[string]bool, 0)

		plosByPloId := make(map[string]entity.NestedOutcome, 0)
		closByCloId := make(map[string]entity.Outcome, 0)
		posByPoId := make(map[string]entity.Outcome, 0)

		plos := make([]entity.NestedOutcome, 0)
		clos := make([]entity.Outcome, 0)
		pos := make([]entity.Outcome, 0)

		for _, c := range cloWithPos {
			plosFromMap, found := plosByPloId[c.ProgramLearningOutcomeCode]

			if !found {
				addedSubPlo[c.SubProgramLearningOutcomeCode] = true

				plosByPloId[c.ProgramLearningOutcomeCode] = entity.NestedOutcome{
					Code: c.ProgramLearningOutcomeCode,
					Name: c.ProgramLearningOutcomeName,
					Nested: []entity.Outcome{
						{
							Code: c.SubProgramLearningOutcomeCode,
							Name: c.SubProgramLearningOutcomeName,
						},
					},
				}
			} else {
				if _, isSubPloAdded := addedSubPlo[c.SubProgramLearningOutcomeCode]; !isSubPloAdded {
					addedSubPlo[c.SubProgramLearningOutcomeCode] = true

					plosFromMap.Nested = append(
						plosByPloId[c.ProgramLearningOutcomeCode].Nested,
						entity.Outcome{
							Code: c.SubProgramLearningOutcomeCode,
							Name: c.SubProgramLearningOutcomeName,
						},
					)

					plosByPloId[c.ProgramLearningOutcomeCode] = plosFromMap
				}
			}

			closByCloId[c.Code] = entity.Outcome{
				Code: c.Code,
				Name: c.Description,
			}

			posByPoId[c.ProgramOutcomeName] = entity.Outcome{
				Code: c.ProgramOutcomeCode,
				Name: c.ProgramOutcomeName,
			}

		}
		for _, plo := range plosByPloId {
			plos = append(plos, plo)
		}
		for _, clo := range closByCloId {
			clos = append(clos, clo)
		}
		for _, po := range posByPoId {
			pos = append(pos, po)
		}

		return plos, clos, pos
	}

	plos, clos, pos := generateOutcome(closWithPos)

	courseResult := entity.CourseResult{
		Plos:              plos,
		Clos:              clos,
		Pos:               pos,
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

	portfolioData := entity.PortfolioData{}

	err = json.Unmarshal(course.PortfolioData, &portfolioData)
	if err != nil {
		return nil, errs.New(0, "cannot unmarshal data from db")
	}

	courseDevelopment := entity.CourseDevelopment{
		Plans:       portfolioData.Development.Plans,
		DoAndChecks: portfolioData.Development.DoAndChecks,
		Acts:        portfolioData.Development.Acts,
		SubjectComments: entity.SubjectComments{
			UpstreamSubjects:   upstreamSubject,
			DownstreamSubjects: downStreamSubject,
			Other:              portfolioData.Development.SubjectComments.Other,
		},
		OtherComment: portfolioData.Development.OtherComment,
	}

	courseSummary := entity.CourseSummary{
		TeachingMethods: portfolioData.Summary.TeachingMethods,
		Objectives:      portfolioData.Summary.Objectives,
		OnlineTools:     portfolioData.Summary.OnlineTools,
	}

	coursePortfolio := &entity.CoursePortfolio{
		CourseInfo:        courseInfo,
		CourseResult:      courseResult,
		CourseSummary:     courseSummary,
		CourseDevelopment: courseDevelopment,
		Raw:               course.PortfolioData,
	}

	return coursePortfolio, nil
}

func (u coursePortfolioUseCase) CalculateGradeDistribution(courseId string) (*entity.GradeDistribution, error) {
	course, err := u.CourseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course by id %s while calculate grade distribution", courseId, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "course id %s not found while calculate grade distribution", courseId)
	}

	assignmentGroups, err := u.AssignmentUseCase.GetGroupByCourseId(courseId, true)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get assignment group while calculate grade distribution")
	}

	// calculate student score
	weightByGroupId := make(map[string]int, 0)
	sumGroupScoreByGroupId := make(map[string]int, 0)
	sumStudentScoreByStudentIdByGroupId := make(map[string]map[string]float64, 0)

	for _, assignmentGroup := range assignmentGroups {
		weightByGroupId[assignmentGroup.Id] = assignmentGroup.Weight

		for _, assignment := range assignmentGroup.Assignments {
			sumGroupScoreByGroupId[assignmentGroup.Id] += assignment.MaxScore

			assignmentScore, _ := u.ScoreUseCase.GetByAssignmentId(assignment.Id)

			for _, score := range assignmentScore.Scores {
				_, ok := sumStudentScoreByStudentIdByGroupId[assignmentGroup.Id]
				if !ok {
					sumStudentScoreByStudentIdByGroupId[assignmentGroup.Id] = make(map[string]float64)
				}

				sumStudentScoreByStudentIdByGroupId[assignmentGroup.Id][score.StudentId] += score.Score
			}

		}
	}

	sumScoreByStudentId := make(map[string]float64, 0)
	for groupId, sumStudentScore := range sumStudentScoreByStudentIdByGroupId {
		for studentId, score := range sumStudentScore {
			studentScore := score / float64(sumGroupScoreByGroupId[groupId]) * float64(weightByGroupId[groupId])
			sumScoreByStudentId[studentId] += studentScore
		}
	}

	// score frequency
	frequencyByScore := make(map[int]int, 0)
	for _, score := range sumScoreByStudentId {
		frequencyByScore[int(math.Round(score))] += 1
	}

	scoreFrequencies := make([]entity.ScoreFrequency, 0, len(frequencyByScore))
	for score, frequency := range frequencyByScore {
		scoreFrequencies = append(scoreFrequencies, entity.ScoreFrequency{
			Score:     score,
			Frequency: frequency,
		})
	}

	// grade frequency
	frequenciesByGrade := make(map[string]int)
	for _, studentScore := range sumScoreByStudentId {
		switch {
		case studentScore >= course.A:
			frequenciesByGrade["A"] += 1
		case studentScore >= course.BP:
			frequenciesByGrade["BP"] += 1
		case studentScore >= course.B:
			frequenciesByGrade["B"] += 1
		case studentScore >= course.CP:
			frequenciesByGrade["CP"] += 1
		case studentScore >= course.C:
			frequenciesByGrade["C"] += 1
		case studentScore >= course.DP:
			frequenciesByGrade["DP"] += 1
		case studentScore >= course.D:
			frequenciesByGrade["D"] += 1
		default:
			frequenciesByGrade["F"] += 1
		}
	}

	gradeFrequencies := []entity.GradeFrequency{
		{
			Name:       "A",
			GradeScore: course.A,
			Frequency:  frequenciesByGrade["A"],
		},
		{
			Name:       "BP",
			GradeScore: course.BP,
			Frequency:  frequenciesByGrade["BP"],
		},
		{
			Name:       "B",
			GradeScore: course.B,
			Frequency:  frequenciesByGrade["B"],
		},
		{
			Name:       "CP",
			GradeScore: course.CP,
			Frequency:  frequenciesByGrade["CP"],
		},
		{
			Name:       "C",
			GradeScore: course.C,
			Frequency:  frequenciesByGrade["C"],
		},
		{
			Name:       "DP",
			GradeScore: course.DP,
			Frequency:  frequenciesByGrade["DP"],
		},
		{
			Name:       "D",
			GradeScore: course.D,
			Frequency:  frequenciesByGrade["D"],
		},
		{
			Name:       "F",
			GradeScore: course.F,
			Frequency:  frequenciesByGrade["F"],
		},
	}

	// gpa
	studentAmount := len(sumScoreByStudentId)

	totalStudentGPA := 0.0
	for grade, frequency := range frequenciesByGrade {
		totalStudentGPA += float64(frequency) * course.CriteriaGrade.GradeToGPA(grade)
	}

	gpa := 0.0
	if totalStudentGPA != 0.0 {
		gpa = totalStudentGPA / float64(studentAmount)
	}

	x := &entity.GradeDistribution{
		StudentAmount:    studentAmount,
		ScoreFrequencies: scoreFrequencies,
		GradeFrequencies: gradeFrequencies,
		GPA:              gpa,
	}

	return x, nil
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
	expectedPassingCloByPoId := make(map[string]float64, 0)
	for _, clo := range clos {
		courseOutcomeByPoId[clo.ProgramOutcomeId] = append(courseOutcomeByPoId[clo.ProgramOutcomeId], entity.CourseOutcome{
			Name:                                clo.Description,
			ExpectedPassingAssignmentPercentage: clo.ExpectedPassingAssignmentPercentage,
			Assessments:                         assessmentsByCloId[clo.Id],
		})
		expectedPassingCloByPoId[clo.ProgramOutcomeId] = clo.ExpectedPassingCloPercentage
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
		checkIsSameOutcomeName := func(foundOutcome []entity.TabeeOutcome, clo entity.CourseLearningOutcomeWithPO) bool {
			isNameSame := false

			for _, tabeeOutcome := range foundOutcome {
				if tabeeOutcome.Name == clo.ProgramOutcomeName {
					isNameSame = true
					break
				}
			}

			return isNameSame
		}

		foundOutcome, found := tabeeOutcomesByPoId[clo.ProgramOutcomeId]
		if !found {
			tabeeOutcomesByPoId[clo.ProgramOutcomeId] = append(tabeeOutcomesByPoId[clo.ProgramOutcomeId], entity.TabeeOutcome{
				Name:                  clo.ProgramOutcomeName,
				CourseOutcomes:        courseOutcomeByPoId[clo.ProgramOutcomeId],
				MinimumPercentage:     passingPoPercentageByPoId[clo.ProgramOutcomeId],
				ExpectedCloPercentage: expectedPassingCloByPoId[clo.ProgramOutcomeId],
			})
			continue
		}

		isNameSame := checkIsSameOutcomeName(foundOutcome, clo)
		if isNameSame {
			continue
		}

		tabeeOutcomesByPoId[clo.ProgramOutcomeId] = append(tabeeOutcomesByPoId[clo.ProgramOutcomeId], entity.TabeeOutcome{
			Name:                  clo.ProgramOutcomeName,
			CourseOutcomes:        courseOutcomeByPoId[clo.ProgramOutcomeId],
			MinimumPercentage:     passingPoPercentageByPoId[clo.ProgramOutcomeId],
			ExpectedCloPercentage: expectedPassingCloByPoId[clo.ProgramOutcomeId],
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

func (u coursePortfolioUseCase) GetStudentOutcomesStatusByCourseId(courseId string) ([]entity.StudentOutcomeStatus, error) {
	course, err := u.CourseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get course id %s while getting clo passing students", course, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrCourseNotFound, "course id %s not found while getting clo passing students", courseId, err)
	}

	ploRecords, err := u.CoursePortfolioRepository.EvaluatePassingPloStudents(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate passing plo student by course id %s", courseId, err)
	}

	poRecords, err := u.CoursePortfolioRepository.EvaluatePassingPoStudents(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate passing po student by course id %s", courseId, err)
	}

	studentPloMap := make(map[string][]entity.PloData)
	studentPoMap := make(map[string][]entity.PoData)

	for _, record := range ploRecords {
		studentPloMap[record.StudentId] = append(studentPloMap[record.StudentId], entity.PloData{
			Id:              record.ProgramLearningOutcomeId,
			Code:            record.Code,
			DescriptionThai: record.DescriptionThai,
			ProgramYear:     record.ProgramYear,
			Pass:            record.Pass,
		})
	}

	for _, record := range poRecords {
		studentPoMap[record.StudentId] = append(studentPoMap[record.StudentId], entity.PoData{
			Id:   record.ProgramOutcomeId,
			Code: record.Code,
			Name: record.Name,
			Pass: record.Pass,
		})
	}

	if len(studentPloMap) != len(studentPoMap) {
		return nil, errs.New(errs.SameCode, "number of students with plo is different from po by course id %s", courseId, err)
	}

	students := make([]entity.StudentOutcomeStatus, 0)

	for studentId := range studentPloMap {
		students = append(students, entity.StudentOutcomeStatus{
			StudentId:               studentId,
			ProgramLearningOutcomes: studentPloMap[studentId],
		})
	}

	for i := range students {
		students[i].ProgramOutcomes = studentPoMap[students[i].StudentId]
	}

	return students, nil
}

func (u coursePortfolioUseCase) GetAllProgramLearningOutcomeCourses() ([]entity.PloCourses, error) {
	records, err := u.CoursePortfolioRepository.EvaluateAllPloCourses()
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate plo courses %s", err)
	}

	plosMap := make(map[string][]entity.CourseData)

	for _, record := range records {
		if record.CourseId == "" {
			plosMap[record.ProgramLearningOutcomeId] = append(plosMap[record.ProgramLearningOutcomeId], entity.CourseData{})
		} else {
			plosMap[record.ProgramLearningOutcomeId] = append(plosMap[record.ProgramLearningOutcomeId], entity.CourseData{
				Id:                record.CourseId,
				Code:              record.Code,
				Name:              record.Name,
				PassingPercentage: record.PassingPercentage,
				Year:              record.Year,
				SemesterSequence:  record.SemesterSequence,
			})
		}
	}

	plos := make([]entity.PloCourses, 0)

	for ploId := range plosMap {
		plos = append(plos, entity.PloCourses{
			ProgramLearningOutcomeId: ploId,
			Courses:                  plosMap[ploId],
		})
	}

	return plos, nil
}

func (u coursePortfolioUseCase) GetAllProgramOutcomeCourses() ([]entity.PoCourses, error) {
	records, err := u.CoursePortfolioRepository.EvaluateAllPoCourses()
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate po courses %s", err)
	}

	posMap := make(map[string][]entity.CourseData)

	for _, record := range records {
		if record.CourseId == "" {
			posMap[record.ProgramOutcomeId] = append(posMap[record.ProgramOutcomeId], entity.CourseData{})
		} else {
			posMap[record.ProgramOutcomeId] = append(posMap[record.ProgramOutcomeId], entity.CourseData{
				Id:                record.CourseId,
				Code:              record.Code,
				Name:              record.Name,
				PassingPercentage: record.PassingPercentage,
				Year:              record.Year,
				SemesterSequence:  record.SemesterSequence,
			})
		}
	}

	pos := make([]entity.PoCourses, 0)

	for poId := range posMap {
		pos = append(pos, entity.PoCourses{
			ProgramOutcomeId: poId,
			Courses:          posMap[poId],
		})
	}

	return pos, nil
}

func (u coursePortfolioUseCase) UpdateCoursePortfolio(courseId string, summary entity.CourseSummary, development entity.CourseDevelopment) error {
	portfolioData := &entity.PortfolioData{
		Summary:     summary,
		Development: development,
	}

	JsonByte, err := json.Marshal(*portfolioData)
	if err != nil {
		return errs.New(errs.SameCode, "cannot marshal course summary %s", err)
	}

	err = u.CoursePortfolioRepository.UpdateCoursePortfolio(courseId, JsonByte)
	if err != nil {
		return errs.New(errs.SameCode, "cannot update course portfolio %s", err)
	}

	return nil
}

func (u coursePortfolioUseCase) GetOutcomesByStudentId(studentId string) ([]entity.StudentOutcomes, error) {
	student, err := u.StudentUseCase.GetById(studentId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get student id %s while getting student outcomes", student, err)
	} else if student == nil {
		return nil, errs.New(errs.ErrStudentNotFound, "student id %s not found while getting student outcomes", studentId, err)
	}

	ploRecords, err := u.CoursePortfolioRepository.EvaluateProgramLearningOutcomesByStudentId(studentId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate student plos by student id %s", studentId, err)
	}

	poRecords, err := u.CoursePortfolioRepository.EvaluateProgramOutcomesByStudentId(studentId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot evaluate student pos by student id %s", studentId, err)
	}

	studentPloMap := make(map[string][]entity.StudentPloData)
	studentPoMap := make(map[string][]entity.StudentPoData)

	PloCourseMap := make(map[string][]entity.StudentCourseData)
	PoCourseMap := make(map[string][]entity.StudentCourseData)

	for _, record := range ploRecords {
		studentPloData, ok := studentPloMap[record.StudentId]
		if !ok {

			studentPloMap[record.StudentId] = append(studentPloMap[record.StudentId], entity.StudentPloData{
				ProgramLearningOutcomeId: record.ProgramLearningOutcomeId,
				Code:                     record.ProgramLearningOutcomeCode,
				DescriptionThai:          record.DescriptionThai,
				ProgramYear:              record.ProgramYear,
			})
		} else {
			isExist := false
			for i := range studentPloData {
				if studentPloData[i].ProgramLearningOutcomeId == record.ProgramLearningOutcomeId {
					isExist = true
					break
				}
			}
			if !isExist {
				studentPloMap[record.StudentId] = append(studentPloMap[record.StudentId], entity.StudentPloData{
					ProgramLearningOutcomeId: record.ProgramLearningOutcomeId,
					Code:                     record.ProgramLearningOutcomeCode,
					DescriptionThai:          record.DescriptionThai,
					ProgramYear:              record.ProgramYear,
				})
			}
		}

		ploData, ok := PloCourseMap[record.ProgramLearningOutcomeId]

		if !ok {

			PloCourseMap[record.ProgramLearningOutcomeId] = append(PloCourseMap[record.ProgramLearningOutcomeId], entity.StudentCourseData{
				Id:               record.CourseId,
				Code:             record.CourseCode,
				Name:             record.CourseName,
				Pass:             record.Pass,
				Year:             record.Year,
				SemesterSequence: record.SemesterSequence,
			})
		} else {
			isExist := false
			for i := range ploData {
				if ploData[i].Id == record.CourseId {
					isExist = true
					break
				}
			}
			if !isExist {
				PloCourseMap[record.ProgramLearningOutcomeId] = append(PloCourseMap[record.ProgramLearningOutcomeId], entity.StudentCourseData{
					Id:               record.CourseId,
					Code:             record.CourseCode,
					Name:             record.CourseName,
					Pass:             record.Pass,
					Year:             record.Year,
					SemesterSequence: record.SemesterSequence,
				})
			}
		}
	}

	for _, record := range poRecords {
		studentData, found := studentPoMap[record.StudentId]
		if !found {
			studentPoMap[record.StudentId] = append(studentPoMap[record.StudentId], entity.StudentPoData{
				ProgramOutcomeId: record.ProgramOutcomeId,
				Code:             record.ProgramOutcomeCode,
				Name:             record.ProgramOutcomeName,
			})

		} else {
			isExist := false
			for i := range studentData {
				if studentData[i].ProgramOutcomeId == record.ProgramOutcomeId {
					isExist = true
					break
				}
			}
			if !isExist {
				studentPoMap[record.StudentId] = append(studentPoMap[record.StudentId], entity.StudentPoData{
					ProgramOutcomeId: record.ProgramOutcomeId,
					Code:             record.ProgramOutcomeCode,
					Name:             record.ProgramOutcomeName,
				})
			}
		}

		poData, found := PoCourseMap[record.ProgramOutcomeId]

		if !found {
			PoCourseMap[record.ProgramOutcomeId] = append(PoCourseMap[record.ProgramOutcomeId], entity.StudentCourseData{
				Id:               record.CourseId,
				Code:             record.CourseCode,
				Name:             record.CourseName,
				Pass:             record.Pass,
				Year:             record.Year,
				SemesterSequence: record.SemesterSequence,
			})
		} else {
			isExist := false
			for i := range poData {
				if poData[i].Id == record.CourseId {
					isExist = true
					break
				}
			}
			if !isExist {
				PoCourseMap[record.ProgramOutcomeId] = append(PoCourseMap[record.ProgramOutcomeId], entity.StudentCourseData{
					Id:               record.CourseId,
					Code:             record.CourseCode,
					Name:             record.CourseName,
					Pass:             record.Pass,
					Year:             record.Year,
					SemesterSequence: record.SemesterSequence,
				})
			}
		}
	}

	students := make([]entity.StudentOutcomes, 0)

	for studentId := range studentPloMap {
		for ploIndex := range studentPloMap[studentId] {
			studentPloMap[studentId][ploIndex].Courses = PloCourseMap[studentPloMap[studentId][ploIndex].ProgramLearningOutcomeId]
		}
		for poIndex := range studentPoMap[studentId] {
			studentPoMap[studentId][poIndex].Courses = PoCourseMap[studentPoMap[studentId][poIndex].ProgramOutcomeId]
		}

		students = append(students, entity.StudentOutcomes{
			StudentId:               studentId,
			ProgramLearningOutcomes: studentPloMap[studentId],
			ProgramOutcomes:         studentPoMap[studentId],
		})
	}

	return students, nil
}
