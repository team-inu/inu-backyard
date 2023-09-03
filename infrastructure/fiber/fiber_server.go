package fiber_handler

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/db_connector"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/fiber_controller"
	"github.com/team-inu/inu-backyard/infrastructure/logger"
	"github.com/team-inu/inu-backyard/repository/repository_gorm"
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
	gormDB, err := db_connector.NewGorm(&db_connector.GormConfig{
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

	f.studentRepository = repository_gorm.NewStudentRepository(f.gorm)

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

	studentController := fiber_controller.NewStudentController(f.studentUseCase)

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
