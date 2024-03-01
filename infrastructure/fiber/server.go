package fiber

import (
	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/team-inu/inu-backyard/entity"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/controller"
	"github.com/team-inu/inu-backyard/infrastructure/fiber/middleware"
	"github.com/team-inu/inu-backyard/internal/config"
	"github.com/team-inu/inu-backyard/internal/validator"
	"github.com/team-inu/inu-backyard/repository"
	"github.com/team-inu/inu-backyard/usecase"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type fiberServer struct {
	config config.FiberServerConfig
	gorm   *gorm.DB
	logger *zap.Logger

	studentRepository                entity.StudentRepository
	courseRepository                 entity.CourseRepository
	courseLearningOutcomeRepository  entity.CourseLearningOutcomeRepository
	programLearningOutcomeRepository entity.ProgramLearningOutcomeRepository
	programOutcomeRepository         entity.ProgramOutcomeRepository
	facultyRepository                entity.FacultyRepository
	departmentRepository             entity.DepartmentRepository
	scoreRepository                  entity.ScoreRepository
	userRepository                   entity.UserRepository
	assignmentRepository             entity.AssignmentRepository
	programmeRepository              entity.ProgrammeRepository
	semesterRepository               entity.SemesterRepository
	enrollmentRepository             entity.EnrollmentRepository
	gradeRepository                  entity.GradeRepository
	sessionRepository                entity.SessionRepository

	studentUseCase                entity.StudentUseCase
	courseUseCase                 entity.CourseUseCase
	courseLearningOutcomeUseCase  entity.CourseLearningOutcomeUseCase
	programLearningOutcomeUseCase entity.ProgramLearningOutcomeUseCase
	programOutcomeUseCase         entity.ProgramOutcomeUseCase
	facultyUseCase                entity.FacultyUseCase
	departmentUseCase             entity.DepartmentUseCase
	scoreUseCase                  entity.ScoreUseCase
	userUseCase                   entity.UserUseCase
	assignmentUseCase             entity.AssignmentUseCase
	programmeUseCase              entity.ProgrammeUseCase
	semesterUseCase               entity.SemesterUseCase
	enrollmentUseCase             entity.EnrollmentUseCase
	gradeUseCase                  entity.GradeUseCase
	sessionUseCase                entity.SessionUseCase
	authUseCase                   entity.AuthUseCase
}

func NewFiberServer(
	config config.FiberServerConfig,
	gorm *gorm.DB,
	logger *zap.Logger,
) *fiberServer {
	return &fiberServer{
		config: config,
		gorm:   gorm,
		logger: logger,
	}
}

func (f *fiberServer) Run() {
	err := f.initRepository()
	if err != nil {
		panic(err)
	}

	f.initUseCase()

	err = f.initController()
	if err != nil {
		panic(err)
	}
}

func (f *fiberServer) initRepository() (err error) {
	f.studentRepository = repository.NewStudentRepositoryGorm(f.gorm)
	f.courseRepository = repository.NewCourseRepositoryGorm(f.gorm)
	f.courseLearningOutcomeRepository = repository.NewCourseLearningOutcomeRepositoryGorm(f.gorm)
	f.programLearningOutcomeRepository = repository.NewProgramLearningOutcomeRepositoryGorm(f.gorm)
	f.programOutcomeRepository = repository.NewProgramOutcomeRepositoryGorm(f.gorm)
	f.facultyRepository = repository.NewFacultyRepositoryGorm(f.gorm)
	f.departmentRepository = repository.NewDepartmentRepositoryGorm(f.gorm)
	f.scoreRepository = repository.NewScoreRepositoryGorm(f.gorm)
	f.userRepository = repository.NewUserRepositoryGorm(f.gorm)
	f.assignmentRepository = repository.NewAssignmentRepositoryGorm(f.gorm)
	f.programmeRepository = repository.NewProgrammeRepositoryGorm(f.gorm)
	f.semesterRepository = repository.NewSemesterRepositoryGorm(f.gorm)
	f.enrollmentRepository = repository.NewEnrollmentRepositoryGorm(f.gorm)
	f.gradeRepository = repository.NewGradeRepositoryGorm(f.gorm)
	f.sessionRepository = repository.NewSessionRepository(f.gorm)

	return nil
}

func (f *fiberServer) initUseCase() {
	programmeUseCase := usecase.NewProgrammeUseCase(f.programmeRepository)
	facultyUseCase := usecase.NewFacultyUseCase(f.facultyRepository)
	departmentUseCase := usecase.NewDepartmentUseCase(f.departmentRepository)
	studentUseCase := usecase.NewStudentUseCase(f.studentRepository, departmentUseCase, programmeUseCase)
	programLearningOutcomeUseCase := usecase.NewProgramLearningOutcomeUseCase(f.programLearningOutcomeRepository, programmeUseCase)
	userUseCase := usecase.NewUserUseCase(f.userRepository)
	semesterUseCase := usecase.NewSemesterUseCase(f.semesterRepository)
	courseUseCase := usecase.NewCourseUseCase(f.courseRepository, semesterUseCase, userUseCase)
	enrollmentUseCase := usecase.NewEnrollmentUseCase(f.enrollmentRepository, studentUseCase, courseUseCase)
	gradeUseCase := usecase.NewGradeUseCase(f.gradeRepository)
	sessionUseCase := usecase.NewSessionUseCase(f.sessionRepository, f.config.Client.Auth)
	authUseCase := usecase.NewAuthUseCase(sessionUseCase, userUseCase)
	programOutcomeUseCase := usecase.NewProgramOutcomeUseCase(f.programOutcomeRepository, semesterUseCase)
	courseLearningOutcomeUseCase := usecase.NewCourseLearningOutcomeUseCase(f.courseLearningOutcomeRepository, courseUseCase, programOutcomeUseCase, programLearningOutcomeUseCase)
	assignmentUseCase := usecase.NewAssignmentUseCase(f.assignmentRepository, courseLearningOutcomeUseCase, courseUseCase)
	scoreUseCase := usecase.NewScoreUseCase(f.scoreRepository, enrollmentUseCase, assignmentUseCase, userUseCase)

	f.assignmentUseCase = assignmentUseCase
	f.authUseCase = authUseCase
	f.courseLearningOutcomeUseCase = courseLearningOutcomeUseCase
	f.courseUseCase = courseUseCase
	f.departmentUseCase = departmentUseCase
	f.enrollmentUseCase = enrollmentUseCase
	f.facultyUseCase = facultyUseCase
	f.gradeUseCase = gradeUseCase
	f.userUseCase = userUseCase
	f.programLearningOutcomeUseCase = programLearningOutcomeUseCase
	f.programOutcomeUseCase = programOutcomeUseCase
	f.programmeUseCase = programmeUseCase
	f.scoreUseCase = scoreUseCase
	f.semesterUseCase = semesterUseCase
	f.sessionUseCase = sessionUseCase
	f.studentUseCase = studentUseCase
}

func (f *fiberServer) initController() error {

	app := fiber.New(fiber.Config{
		AppName:      "inu-backyard",
		ErrorHandler: errorHandler(f.logger),
	})

	app.Use(middleware.NewCorsMiddleware(f.config.Client.Cors.AllowOrigins))
	app.Use(middleware.NewLogger(fiberzap.Config{
		Logger: f.logger,
	}))

	validator := validator.NewPayloadValidator(&f.config.Client.Auth)

	authMiddleware := middleware.NewAuthMiddleware(validator, f.authUseCase)

	studentController := controller.NewStudentController(validator, f.studentUseCase)
	courseController := controller.NewCourseController(validator, f.courseUseCase)
	courseLearningOutcomeController := controller.NewCourseLearningOutcomeController(validator, f.courseLearningOutcomeUseCase)
	programLearningOutcomeController := controller.NewProgramLearningOutcomeController(validator, f.programLearningOutcomeUseCase)
	subProgramLearningOutcomeController := controller.NewSubProgramLearningOutcomeController(validator, f.programLearningOutcomeUseCase)
	programOutcomeController := controller.NewProgramOutcomeController(validator, f.programOutcomeUseCase)
	facultyController := controller.NewFacultyController(validator, f.facultyUseCase)
	departmentController := controller.NewDepartmentController(validator, f.departmentUseCase)
	scoreController := controller.NewScoreController(validator, f.scoreUseCase)
	userController := controller.NewUserController(validator, f.userUseCase)
	assignmentController := controller.NewAssignmentController(validator, f.assignmentUseCase)
	programmeController := controller.NewProgrammeController(validator, f.programmeUseCase)
	semesterController := controller.NewSemesterController(validator, f.semesterUseCase)
	enrollmentController := controller.NewEnrollmentController(validator, f.enrollmentUseCase)
	gradeController := controller.NewGradeController(validator, f.gradeUseCase)
	authController := controller.NewAuthController(validator, f.config.Client.Auth, f.authUseCase, f.userUseCase)

	app.Get("/students/:studentId", studentController.GetById)
	app.Get("/students", studentController.GetStudents)
	app.Post("/students", studentController.Create)
	app.Post("/students/bulk", studentController.CreateMany)
	app.Patch("/students/:studentId", studentController.Update)
	app.Delete("/students/:studentId", studentController.Delete)

	app.Get("/courses", courseController.GetAll)
	app.Get("/courses/:courseId", courseController.GetById)
	app.Post("/courses", courseController.Create)
	app.Patch("/courses/:courseId", courseController.Update)
	app.Delete("/courses/:courseId", courseController.Delete)

	app.Get("/clos", courseLearningOutcomeController.GetAll)
	app.Get("/clos/:cloId", courseLearningOutcomeController.GetById)
	app.Post("/clos", courseLearningOutcomeController.Create)
	app.Patch("/clos/:cloId", courseLearningOutcomeController.Update)
	app.Delete("/clos/:cloId", courseLearningOutcomeController.Delete)

	app.Post("/clos/:cloId/subplos", courseLearningOutcomeController.CreateLinkSubProgramLearningOutcome)
	app.Delete("/clos/:cloId/subplos/:subploId", courseLearningOutcomeController.DeleteLinkSubProgramLearningOutcome)

	app.Get("/courses/:courseId/clos", courseLearningOutcomeController.GetByCourseId)
	app.Get("/courses/:courseId/enrollments", enrollmentController.GetByCourseId)
	app.Get("/courses/:courseId/assignments", assignmentController.GetByCourseId)

	app.Get("/plos", programLearningOutcomeController.GetAll)
	app.Get("/plos/:ploId", programLearningOutcomeController.GetById)
	app.Post("/plos", programLearningOutcomeController.Create)
	app.Patch("/plos/:ploId", programLearningOutcomeController.Update)
	app.Delete("/plos/:ploId", programLearningOutcomeController.Delete)

	app.Get("/splos", subProgramLearningOutcomeController.GetAll)
	app.Get("/splos/:sploId", subProgramLearningOutcomeController.GetById)
	app.Post("/splos", subProgramLearningOutcomeController.Create)
	app.Patch("/splos/:sploId", subProgramLearningOutcomeController.Update)
	app.Delete("/splos/:sploId", subProgramLearningOutcomeController.Delete)

	app.Get("/pos", programOutcomeController.GetAll)
	app.Get("/pos/:poId", programOutcomeController.GetById)
	app.Post("/pos", programOutcomeController.Create)
	app.Patch("/pos/:poId", programOutcomeController.Update)
	app.Delete("/pos/:poId", programOutcomeController.Delete)

	app.Get("/faculties", facultyController.GetAll)
	app.Get("/faculties/:facultyName", facultyController.GetById)
	app.Post("/faculties", facultyController.Create)
	app.Patch("/faculties/:facultyName", facultyController.Update)
	app.Delete("/faculties/:facultyName", facultyController.Delete)

	app.Get("/departments", departmentController.GetAll)
	app.Get("/departments/:departmentName", departmentController.GetByName)
	app.Post("/departments", departmentController.Create)
	app.Patch("/departments/:departmentName", departmentController.Update)
	app.Delete("/departments/:departmentName", departmentController.Delete)

	app.Get("/scores", scoreController.GetAll)
	app.Get("/scores/:scoreId", scoreController.GetById)
	app.Post("/scores", scoreController.CreateMany)
	app.Patch("/scores/:scoreId", scoreController.Update)
	app.Delete("/scores/:scoreId", scoreController.Delete)

	app.Get("/users", userController.GetAll)
	app.Get("/users/:userId", userController.GetById)
	app.Post("/users", userController.Create)
	app.Post("/users/bulk", userController.CreateMany)
	app.Patch("/users/:userId", userController.Update)
	app.Delete("/users/:userId", userController.Delete)

	app.Get("/assignments", assignmentController.GetAssignments)
	app.Get("/assignments/:assignmentId", assignmentController.GetById)
	app.Post("/assignments", assignmentController.Create)
	app.Patch("/assignments/:assignmentId", assignmentController.Update)
	app.Delete("/assignments/:assignmentId", assignmentController.Delete)
	app.Get("/assignments/:assignmentId/scores", scoreController.GetByAssignmentId)

	app.Get("/programmes", programmeController.GetAll)
	app.Get("/programmes/:programmeName", programmeController.GetByName)
	app.Post("/programmes", programmeController.Create)
	app.Patch("/programmes/:programmeName", programmeController.Update)
	app.Delete("/programmes/:programmeName", programmeController.Delete)

	app.Get("/enrollments", enrollmentController.GetAll)
	app.Get("/enrollments/:enrollmentId", enrollmentController.GetById)
	app.Post("/enrollments", enrollmentController.Create)
	app.Patch("/enrollments/:enrollmentId", enrollmentController.Update)
	app.Delete("/enrollments/:enrollmentId", enrollmentController.Delete)

	app.Get("/semesters", semesterController.GetAll)
	app.Get("/semesters/:semesterId", semesterController.GetById)
	app.Post("/semesters", semesterController.Create)
	app.Patch("/semesters/:semesterId", semesterController.Update)
	app.Delete("/semesters/:semesterId", semesterController.Delete)

	app.Get("/grades", gradeController.GetAll)
	app.Get("/grades/:gradeId", gradeController.GetById)
	app.Post("/grades", gradeController.Create)
	app.Patch("/grades/:gradeId", gradeController.Update)
	app.Delete("/grades/:gradeId", gradeController.Delete)

	app.Post("/auth/login", authController.SignIn)
	app.Get("/auth/logout", authController.SignOut)
	app.Get("/auth/me", authMiddleware, authController.Me)

	app.Get("/metrics", monitor.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	err := app.Listen(":3001")

	return err
}
