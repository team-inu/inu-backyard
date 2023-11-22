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

	studentRepository entity.StudentRepository

	studentUseCase entity.StudentUseCase

	courseRepository entity.CourseRepository

	courseUseCase entity.CourseUsecase

	courseLearningOutcomeRepository entity.CourseLearningOutcomeRepository

	courseLearningOutcomeUsecase entity.CourseLearningOutcomeUsecase

	programLearningOutcomeRepository entity.ProgramLearningOutcomeRepository

	programLearningOutcomeUsecase entity.ProgramLearningOutcomeUsecase

	subProgramLearningOutcomeRepository entity.SubProgramLearningOutcomeRepository

	subProgramLearningOutcomeUsecase entity.SubProgramLearningOutcomeUsecase

	programOutcomeRepository entity.ProgramOutcomeRepository

	programOutcomeUsecase entity.ProgramOutcomeUsecase

	facultyRepository entity.FacultyRepository

	facultyUsecase entity.FacultyUseCase

	departmentRepository entity.DepartmentRepository

	departmentUsecase entity.DepartmentUseCase

	assessmentRepository entity.AssessmentRepository

	assessmentUsecase entity.AssessmentUseCase

	programmeRepository entity.ProgrammeRepository

	programmeUsecase entity.ProgrammeUseCase
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

	f.assessmentRepository = repository.NewAssessmentRepositoryGorm(f.gorm)

	f.programmeRepository = repository.NewProgrammeRepositoryGorm(f.gorm)

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
	f.assessmentUsecase = usecase.NewAssessmentUseCase(f.assessmentRepository)
	f.programmeUsecase = usecase.NewProgrammeUseCase(f.programmeRepository)
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

	assessmentController := controller.NewAssessmentController(f.assessmentUsecase)

	programmeController := controller.NewProgrammeController(f.programmeUsecase)

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.NewZapLogger(),
	}))

	app.Get("/students/:studentId", studentController.GetByID)
	app.Get("/students", studentController.GetStudents)
	app.Post("/students", studentController.Create)
	app.Post("/students/bulk", studentController.CreateMany)

	app.Get("/courses", courseController.GetAll)
	app.Get("/courses/:courseId", courseController.GetByID)
	app.Post("/courses", courseController.Create)
	app.Delete("/courses/:courseId", courseController.Delete)

	app.Get("/clos", courseLearningOutcomeController.GetAll)
	app.Get("/clos/:cloId", courseLearningOutcomeController.GetByID)
	app.Get("/courses/:courseId/clos", courseLearningOutcomeController.GetByCourseID)
	app.Post("/clos", courseLearningOutcomeController.Create)

	app.Get("/plos", programLearningOutcomeController.GetAll)
	app.Get("/plos/:ploId", programLearningOutcomeController.GetByID)
	app.Post("/plos", programLearningOutcomeController.Create)

	app.Get("/splos", subProgramLearningOutcomeController.GetAll)
	app.Get("/splos/:sploId", subProgramLearningOutcomeController.GetByID)
	app.Post("/splos", subProgramLearningOutcomeController.Create)

	app.Get("/pos", programOutcomeController.GetAll)
	app.Get("/pos/:poId", programOutcomeController.GetByID)
	app.Post("/pos", programOutcomeController.Create)

	app.Get("/faculties", facultyController.GetAll)
	app.Get("/faculties/:facultyID", facultyController.GetByID)
	app.Post("/faculties", facultyController.Create)
	app.Patch("/faculties", facultyController.Update)
	app.Delete("/faculties", facultyController.Delete)

	app.Get("/departments", departmentController.GetAll)
	app.Get("/departments/:departmentID", departmentController.GetByName)
	app.Post("/departments", departmentController.Create)
	app.Patch("/departments", departmentController.Update)
	app.Delete("/departments", departmentController.Delete)

	app.Get("/assessments", assessmentController.GetAssessments)
	app.Get("/assessments/:assessmentID", assessmentController.GetByID)
	app.Post("/assessments", assessmentController.Create)
	app.Post("/assessments/bulk", assessmentController.CreateMany)
	app.Patch("/assessments", assessmentController.Update)
	app.Delete("/assessments", assessmentController.Delete)

	app.Get("/programmes", programmeController.GetAll)
	app.Get("/programmes/:programmeName", programmeController.GetByName)
	app.Post("/programmes", programmeController.Create)
	app.Patch("/programmes/:programmeName", programmeController.Update)
	app.Delete("/programmes/:programmeName", programmeController.Delete)

	app.Get("/metrics", monitor.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3001")
}
