package fiber

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
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

	courseLearningOutcomeRepository entity.CourseLearningOutcomeRepository

	courseLearningOutcomeUsecase entity.CourseLearningOutcomeUsecase

	programLearningOutcomeRepository entity.ProgramLearningOutcomeRepository

	programLearningOutcomeUsecase entity.ProgramLearningOutcomeUsecase

	subProgramLearningOutcomeRepository entity.SubProgramLearningOutcomeRepository

	subProgramLearningOutcomeUsecase entity.SubProgramLearningOutcomeUsecase

	programOutcomeRepository entity.ProgramOutcomeRepository

	programOutcomeUsecase entity.ProgramOutcomeUsecase
}

func NewFiberServer() *fiberServer {
	return &fiberServer{}
}

func (f *fiberServer) Run(config FiberServerConfig) {
	f.config = config

	f.initRepository()
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

	f.courseLearningOutcomeRepository = repository.NewCourseLearningOutcomeRepositoryGorm(f.gorm)

	f.programLearningOutcomeRepository = repository.NewProgramLearningOutcomeRepositoryGorm(f.gorm)

	f.subProgramLearningOutcomeRepository = repository.NewSubProgramLearningOutcomeRepositoryGorm(f.gorm)

	f.programOutcomeRepository = repository.NewProgramOutcomeRepositoryGorm(f.gorm)

	return nil
}

func (f *fiberServer) initUseCase() {
	f.studentUseCase = usecase.NewStudentUseCase(f.studentRepository)
	f.courseLearningOutcomeUsecase = usecase.NewCourseLearningOutcomeUsecase(f.courseLearningOutcomeRepository)
	f.programLearningOutcomeUsecase = usecase.NewProgramLearningOutcomeUsecase(f.programLearningOutcomeRepository)
	f.subProgramLearningOutcomeUsecase = usecase.NewSubProgramLearningOutcomeUsecase(f.subProgramLearningOutcomeRepository)
	f.programOutcomeUsecase = usecase.NewProgramOutcomeUsecase(f.programOutcomeRepository)
}

func (f *fiberServer) initController() {
	fiberConfig := fiber.Config{
		AppName: "inu-backyard",
		// EnablePrintRoutes: true,
	}

	app := fiber.New(fiberConfig)

	studentController := controller.NewStudentController(f.studentUseCase)

	courseLearningOutcomeController := controller.NewCourseLearningOutcomeController(f.courseLearningOutcomeUsecase)

	programLearningOutcomeController := controller.NewProgramLearningOutcomeController(f.programLearningOutcomeUsecase)

	subProgramLearningOutcomeController := controller.NewSubProgramLearningOutcomeController(f.subProgramLearningOutcomeUsecase)

	programOutcomeController := controller.NewProgramOutcomeController(f.programOutcomeUsecase)

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.NewZapLogger(),
	}))

	app.Get("/students", studentController.GetAll)
	app.Get("/students/:studentId", studentController.GetByID)
	app.Post("/students", studentController.Create)

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

	app.Get("/metrics", monitor.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
