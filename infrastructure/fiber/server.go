package fiber

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/database"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/controller"
	"github.com/team-inu/inu-backyard/infrastructure/logger"
	"github.com/team-inu/inu-backyard/repository"
	"github.com/team-inu/inu-backyard/usecase"
	"gorm.io/gorm"
)

type fiberServer struct {
	gorm *gorm.DB

	studentRepository entity.StudentRepository

	studentUseCase entity.StudentUseCase
}

func NewFiberServer() *fiberServer {
	return &fiberServer{}
}

func (f *fiberServer) Run() {
	f.initRepository()
	f.initUseCase()
	f.initController()
}

func (f *fiberServer) initRepository() (err error) {
	gormDB, err := database.NewGorm(&database.GormConfig{
		User:     "root",
		Password: "root",
		Host:     "mysql",
		Port:     "3306",
		Database: "inu_backyard",
	})
	if err != nil {
		panic(err)
	}

	f.gorm = gormDB

	f.studentRepository = repository.NewStudentRepositoryGorm(f.gorm)

	return nil
}

func (f *fiberServer) initUseCase() {
	f.studentUseCase = usecase.NewStudentUseCase(f.studentRepository)
}

func (f *fiberServer) initController() {
	fiberConfig := fiber.Config{
		AppName: "inu-backyard",
		// EnablePrintRoutes: true,
	}

	app := fiber.New(fiberConfig)

	studentController := controller.NewStudentController(f.studentUseCase)

	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.NewZapLogger(),
	}))

	app.Get("/students", studentController.GetAll)
	app.Get("/students/:studentId", studentController.GetByID)
	app.Post("/students", studentController.Create)

	app.Get("/metrics", monitor.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	app.Listen(":3000")
}
