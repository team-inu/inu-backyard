package fiber

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/database"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/controller"
	"github.com/team-inu/inu-backyard/internal/logger"
	"github.com/team-inu/inu-backyard/repository"
	"github.com/team-inu/inu-backyard/usecase"
	"gorm.io/gorm"
)

type FiberServerConfig struct {
	Database database.GormConfig
}

type fiberServer struct {
	config FiberServerConfig

	gorm *gorm.DB

	studentRepository                   entity.StudentRepository
	courseRepository                    entity.CourseRepository
	courseLearningOutcomeRepository     entity.CourseLearningOutcomeRepository
	programLearningOutcomeRepository    entity.ProgramLearningOutcomeRepository
	subProgramLearningOutcomeRepository entity.SubProgramLearningOutcomeRepository
	programOutcomeRepository            entity.ProgramOutcomeRepository
	facultyRepository                   entity.FacultyRepository
	departmentRepository                entity.DepartmentRepository
	scoreRepository                     entity.ScoreRepository
	lecturerRepository                  entity.LecturerRepository
	assessmentRepository                entity.AssessmentRepository
	programmeRepository                 entity.ProgrammeRepository
	semesterRepository                  entity.SemesterRepository
	enrollmentRepository                entity.EnrollmentRepository
	gradeRepository                     entity.GradeRepository

	studentUseCase                   entity.StudentUseCase
	courseUseCase                    entity.CourseUsecase
	courseLearningOutcomeUsecase     entity.CourseLearningOutcomeUsecase
	programLearningOutcomeUsecase    entity.ProgramLearningOutcomeUsecase
	subProgramLearningOutcomeUsecase entity.SubProgramLearningOutcomeUsecase
	programOutcomeUsecase            entity.ProgramOutcomeUsecase
	facultyUsecase                   entity.FacultyUseCase
	departmentUsecase                entity.DepartmentUseCase
	scoreUsecase                     entity.ScoreUsecase
	lecturerUsecase                  entity.LecturerUseCase
	assessmentUsecase                entity.AssessmentUseCase
	programmeUsecase                 entity.ProgrammeUseCase
	semesterUsecase                  entity.SemesterUseCase
	enrollmentUsecase                entity.EnrollmentUseCase
	gradeUsecase                     entity.GradeUseCase
}

func NewFiberServer() *fiberServer {
	return &fiberServer{}
}

func (f *fiberServer) Run(config FiberServerConfig) {
	f.config = config
	err := f.initRepository()
	if err != nil {
		panic(err)
	}
	err = f.gorm.AutoMigrate(
		&entity.Student{},
		&entity.Enrollment{},
	)

	if err != nil {
		panic(err)
	}

	f.initUseCase()

	f.initController()
}

func (f *fiberServer) initRepository() (err error) {
	gormDB, err := database.NewGorm(&f.config.Database)
	if err != nil {
		panic(err)
	}

	f.gorm = gormDB

	f.studentRepository = repository.NewStudentRepositoryGorm(f.gorm)
	f.courseRepository = repository.NewCourseRepositoryGorm(f.gorm)
	f.courseLearningOutcomeRepository = repository.NewCourseLearningOutcomeRepositoryGorm(f.gorm)
	f.programLearningOutcomeRepository = repository.NewProgramLearningOutcomeRepositoryGorm(f.gorm)
	f.subProgramLearningOutcomeRepository = repository.NewSubProgramLearningOutcomeRepositoryGorm(f.gorm)
	f.programOutcomeRepository = repository.NewProgramOutcomeRepositoryGorm(f.gorm)
	f.facultyRepository = repository.NewFacultyRepositoryGorm(f.gorm)
	f.departmentRepository = repository.NewDepartmentRepositoryGorm(f.gorm)
	f.scoreRepository = repository.NewScoreRepositoryGorm(f.gorm)
	f.lecturerRepository = repository.NewLecturerRepositoryGorm(f.gorm)
	f.assessmentRepository = repository.NewAssessmentRepositoryGorm(f.gorm)
	f.programmeRepository = repository.NewProgrammeRepositoryGorm(f.gorm)
	f.semesterRepository = repository.NewSemesterRepositoryGorm(f.gorm)

	f.enrollmentRepository = repository.NewEnrollmentRepositoryGorm(f.gorm)

	f.gradeRepository = repository.NewGradeRepositoryGorm(f.gorm)

	return nil
}

func (f *fiberServer) initUseCase() {
	f.studentUseCase = usecase.NewStudentUseCase(f.studentRepository)
	f.courseUseCase = usecase.NewCourseUsecase(f.courseRepository)
	f.courseLearningOutcomeUsecase = usecase.NewCourseLearningOutcomeUsecase(f.courseLearningOutcomeRepository)
	f.programLearningOutcomeUsecase = usecase.NewProgramLearningOutcomeUsecase(f.programLearningOutcomeRepository)
	f.subProgramLearningOutcomeUsecase = usecase.NewSubProgramLearningOutcomeUsecase(f.subProgramLearningOutcomeRepository)
	f.programOutcomeUsecase = usecase.NewProgramOutcomeUsecase(f.programOutcomeRepository)
	f.facultyUsecase = usecase.NewFacultyUseCase(f.facultyRepository)
	f.departmentUsecase = usecase.NewDepartmentUseCase(f.departmentRepository)
	f.scoreUsecase = usecase.NewScoreUseCase(f.scoreRepository)
	f.lecturerUsecase = usecase.NewLecturerUseCase(f.lecturerRepository)
	f.assessmentUsecase = usecase.NewAssessmentUseCase(f.assessmentRepository)
	f.programmeUsecase = usecase.NewProgrammeUseCase(f.programmeRepository)
	f.enrollmentUsecase = usecase.NewEnrollmentUseCase(f.enrollmentRepository)
	f.semesterUsecase = usecase.NewSemesterUseCase(f.semesterRepository)
	f.gradeUsecase = usecase.NewGradeUseCase(f.gradeRepository)
}

func (f *fiberServer) initController() {
	fiberConfig := fiber.Config{
		AppName:      "inu-backyard",
		ErrorHandler: errorHandler(logger.NewZapLogger()),
		// EnablePrintRoutes: true,
	}

	app := fiber.New(fiberConfig)

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin",
		AllowOrigins:     "*",
		AllowCredentials: true,
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))

	studentController := controller.NewStudentController(f.studentUseCase)
	courseController := controller.NewCourseController(f.courseUseCase)
	courseLearningOutcomeController := controller.NewCourseLearningOutcomeController(f.courseLearningOutcomeUsecase)
	programLearningOutcomeController := controller.NewProgramLearningOutcomeController(f.programLearningOutcomeUsecase)
	subProgramLearningOutcomeController := controller.NewSubProgramLearningOutcomeController(f.subProgramLearningOutcomeUsecase)
	programOutcomeController := controller.NewProgramOutcomeController(f.programOutcomeUsecase)
	facultyController := controller.NewFacultyController(f.facultyUsecase)
	departmentController := controller.NewDepartmentController(f.departmentUsecase)
	scoreController := controller.NewScoreController(f.scoreUsecase)

	lecturerController := controller.NewLecturerController(f.lecturerUsecase)

	assessmentController := controller.NewAssessmentController(f.assessmentUsecase)
	programmeController := controller.NewProgrammeController(f.programmeUsecase)
	semesterController := controller.NewSemesterController(f.semesterUsecase)

	enrollmentController := controller.NewEnrollmentController(f.enrollmentUsecase)

	gradeController := controller.NewGradeController(f.gradeUsecase)

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.NewZapLogger(),
	}))

	app.Get("/students/:studentId", studentController.GetByID)
	app.Get("/students", studentController.GetStudents)
	app.Post("/students", studentController.Create)
	app.Post("/students/bulk", studentController.CreateMany)
	app.Patch("/students/:studentId", studentController.Update)
	app.Delete("/students/:studentId", studentController.Delete)

	app.Get("/courses", courseController.GetAll)
	app.Get("/courses/:courseId", courseController.GetByID)
	app.Post("/courses", courseController.Create)
	app.Patch("/courses/:courseId", courseController.Update)
	app.Delete("/courses/:courseId", courseController.Delete)

	app.Get("/clos", courseLearningOutcomeController.GetAll)
	app.Get("/clos/:cloId", courseLearningOutcomeController.GetByID)
	app.Get("/courses/:courseId/clos", courseLearningOutcomeController.GetByCourseID)
	app.Post("/clos", courseLearningOutcomeController.Create)
	app.Patch("/clos/:cloId", courseLearningOutcomeController.Update)
	app.Delete("/clos/:cloId", courseLearningOutcomeController.Delete)

	app.Get("/plos", programLearningOutcomeController.GetAll)
	app.Get("/plos/:ploId", programLearningOutcomeController.GetByID)
	app.Post("/plos", programLearningOutcomeController.Create)
	app.Patch("/plos/:ploId", programLearningOutcomeController.Update)
	app.Delete("/plos/:ploId", programLearningOutcomeController.Delete)

	app.Get("/splos", subProgramLearningOutcomeController.GetAll)
	app.Get("/splos/:sploId", subProgramLearningOutcomeController.GetByID)
	app.Post("/splos", subProgramLearningOutcomeController.Create)
	app.Patch("/splos/:sploId", subProgramLearningOutcomeController.Update)
	app.Delete("/splos/:sploId", subProgramLearningOutcomeController.Delete)

	app.Get("/pos", programOutcomeController.GetAll)
	app.Get("/pos/:poId", programOutcomeController.GetByID)
	app.Post("/pos", programOutcomeController.Create)
	app.Patch("/pos/:poId", programOutcomeController.Update)
	app.Delete("/pos/:poId", programOutcomeController.Delete)

	app.Get("/faculties", facultyController.GetAll)
	app.Get("/faculties/:facultyName", facultyController.GetByID)
	app.Post("/faculties", facultyController.Create)
	app.Patch("/faculties/:facultyName", facultyController.Update)
	app.Delete("/faculties/:facultyName", facultyController.Delete)

	app.Get("/departments", departmentController.GetAll)
	app.Get("/departments/:departmentName", departmentController.GetByName)
	app.Post("/departments", departmentController.Create)
	app.Patch("/departments/:departmentName", departmentController.Update)
	app.Delete("/departments/:departmentName", departmentController.Delete)

	app.Get("/scores", scoreController.GetAll)
	app.Get("/scores/:scoreID", scoreController.GetByID)
	app.Post("/scores", scoreController.Create)
	app.Patch("/scores/:scoreID", scoreController.Update)
	app.Delete("/scores/:scoreID", scoreController.Delete)

	app.Get("/lecturers", lecturerController.GetAll)
	app.Get("/lecturers/:lecturerID", lecturerController.GetByID)
	app.Post("/lecturers", lecturerController.Create)
	app.Patch("/lecturers/:lecturerID", lecturerController.Update)
	app.Delete("/lecturers/:lecturerID", lecturerController.Delete)

	app.Get("/assessments", assessmentController.GetAssessments)
	app.Get("/assessments/:assessmentID", assessmentController.GetByID)
	app.Post("/assessments", assessmentController.Create)
	app.Post("/assessments/bulk", assessmentController.CreateMany)
	app.Patch("/assessments/:assessmentID", assessmentController.Update)
	app.Delete("/assessments/:assessmentID", assessmentController.Delete)

	app.Get("/programmes", programmeController.GetAll)
	app.Get("/programmes/:programmeName", programmeController.GetByName)
	app.Post("/programmes", programmeController.Create)
	app.Patch("/programmes/:programmeName", programmeController.Update)
	app.Delete("/programmes/:programmeName", programmeController.Delete)

	app.Get("/enrollments", enrollmentController.GetAll)
	app.Get("/enrollments/:enrollmentID", enrollmentController.GetByID)
	app.Post("/enrollments", enrollmentController.Create)
	app.Patch("/enrollments/:enrollmentID", enrollmentController.Update)
	app.Delete("/enrollments/:enrollmentID", enrollmentController.Delete)

	app.Get("/semesters", semesterController.GetAll)
	app.Get("/semesters/:semesterID", semesterController.GetByID)
	app.Post("/semesters", semesterController.Create)
	app.Patch("/semesters/:semesterID", semesterController.Update)
	app.Delete("/semesters/:semesterID", semesterController.Delete)

	app.Get("/grades", gradeController.GetAll)
	app.Get("/grades/:gradeID", gradeController.GetByID)
	app.Post("/grades", gradeController.Create)
	app.Patch("/grades/:gradeID", gradeController.Update)
	app.Delete("/grades/:gradeID", gradeController.Delete)

	app.Get("/metrics", monitor.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3001")
}
