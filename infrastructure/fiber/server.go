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
	coursePortfolioRepository        entity.CoursePortfolioRepository
	predictionRepository             entity.PredictionRepository
	courseStreamRepository           entity.CourseStreamRepository

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
	coursePortfolioUseCase        entity.CoursePortfolioUseCase
	predictionUseCase             entity.PredictionUseCase
	courseStreamUseCase           entity.CourseStreamsUseCase
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
	f.coursePortfolioRepository = repository.NewCoursePortfolioRepositoryGorm(f.gorm)
	f.predictionRepository = repository.NewPredictionRepositoryGorm(f.gorm)
	f.courseStreamRepository = repository.NewCourseStreamRepository(f.gorm)

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
	gradeUseCase := usecase.NewGradeUseCase(f.gradeRepository, studentUseCase, semesterUseCase)
	sessionUseCase := usecase.NewSessionUseCase(f.sessionRepository, f.config.Client.Auth)
	authUseCase := usecase.NewAuthUseCase(sessionUseCase, userUseCase)
	programOutcomeUseCase := usecase.NewProgramOutcomeUseCase(f.programOutcomeRepository, semesterUseCase)
	courseLearningOutcomeUseCase := usecase.NewCourseLearningOutcomeUseCase(f.courseLearningOutcomeRepository, courseUseCase, programOutcomeUseCase, programLearningOutcomeUseCase)
	assignmentUseCase := usecase.NewAssignmentUseCase(f.assignmentRepository, courseLearningOutcomeUseCase, courseUseCase)
	scoreUseCase := usecase.NewScoreUseCase(f.scoreRepository, enrollmentUseCase, assignmentUseCase, courseUseCase, userUseCase, studentUseCase)
	courseStreamUseCase := usecase.NewCourseStreamUseCase(f.courseStreamRepository, courseUseCase)
	coursePortfolioUseCase := usecase.NewCoursePortfolioUseCase(f.coursePortfolioRepository, courseUseCase, userUseCase, enrollmentUseCase, assignmentUseCase, scoreUseCase, courseLearningOutcomeUseCase, courseStreamUseCase)
	predictionUseCase := usecase.NewPredictionUseCase(f.predictionRepository, f.config)

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
	f.coursePortfolioUseCase = coursePortfolioUseCase
	f.predictionUseCase = predictionUseCase
	f.courseStreamUseCase = courseStreamUseCase
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
	predictionController := controller.NewPredictionController(validator, f.predictionUseCase)
	coursePortfolioController := controller.NewCoursePortfolioController(validator, f.coursePortfolioUseCase)
	courseStreamController := controller.NewCourseStreamController(validator, f.courseStreamUseCase)
	authController := controller.NewAuthController(validator, f.config.Client.Auth, f.authUseCase, f.userUseCase)

	api := app.Group("/")

	// student route
	student := api.Group("/students", authMiddleware)

	student.Get("/", studentController.GetStudents)
	student.Post("/", studentController.Create)
	student.Post("/bulk", studentController.CreateMany)
	student.Get("/:studentId", studentController.GetById)
	student.Patch("/:studentId", studentController.Update)
	student.Delete("/:studentId", studentController.Delete)

	// course route
	course := api.Group("/courses", authMiddleware)

	course.Get("/", courseController.GetAll)
	course.Post("/", courseController.Create)
	course.Get("/:courseId", courseController.GetById)
	course.Patch("/:courseId", courseController.Update)
	course.Delete("/:courseId", courseController.Delete)

	course.Get("/:courseId/clos", courseLearningOutcomeController.GetByCourseId)
	course.Get("/:courseId/clos/students", coursePortfolioController.GetCloPassingStudentsByCourseId)
	course.Get("/:courseId/enrollments", enrollmentController.GetByCourseId)
	course.Get("/:courseId/assignments", assignmentController.GetByCourseId)
	course.Get("/:courseId/portfolio", coursePortfolioController.Generate)

	// course learning outcome route
	clo := api.Group("/clos", authMiddleware)

	clo.Get("/", courseLearningOutcomeController.GetAll)
	clo.Post("/", courseLearningOutcomeController.Create)
	clo.Get("/:cloId", courseLearningOutcomeController.GetById)
	clo.Patch("/:cloId", courseLearningOutcomeController.Update)
	clo.Delete("/:cloId", courseLearningOutcomeController.Delete)

	// sub program learning outcome by course learning outcome route
	subPloByClo := clo.Group("/:cloId/subplos", authMiddleware)

	subPloByClo.Post("/", courseLearningOutcomeController.CreateLinkSubProgramLearningOutcome)
	subPloByClo.Delete("/:subploId", courseLearningOutcomeController.DeleteLinkSubProgramLearningOutcome)

	// program learning outcome route
	plo := api.Group("/plos", authMiddleware)

	plo.Get("/", programLearningOutcomeController.GetAll)
	plo.Post("/", programLearningOutcomeController.Create)
	plo.Get("/:ploId", programLearningOutcomeController.GetById)
	plo.Patch("/:ploId", programLearningOutcomeController.Update)
	plo.Delete("/:ploId", programLearningOutcomeController.Delete)

	// sub program learning outcome route
	subPlo := api.Group("/splos", authMiddleware)

	subPlo.Get("/", subProgramLearningOutcomeController.GetAll)
	subPlo.Post("/", subProgramLearningOutcomeController.Create)
	subPlo.Get("/:sploId", subProgramLearningOutcomeController.GetById)
	subPlo.Patch("/:sploId", subProgramLearningOutcomeController.Update)
	subPlo.Delete("/:sploId", subProgramLearningOutcomeController.Delete)

	// program outcome route
	pos := api.Group("/pos", authMiddleware)

	pos.Get("/", programOutcomeController.GetAll)
	pos.Post("/", programOutcomeController.Create)
	pos.Get("/:poId", programOutcomeController.GetById)
	pos.Patch("/:poId", programOutcomeController.Update)
	pos.Delete("/:poId", programOutcomeController.Delete)

	// faculty route
	faculty := api.Group("/faculties", authMiddleware)

	faculty.Get("/", facultyController.GetAll)
	faculty.Post("/", facultyController.Create)
	faculty.Get("/:facultyName", facultyController.GetById)
	faculty.Patch("/:facultyName", facultyController.Update)
	faculty.Delete("/:facultyName", facultyController.Delete)

	// department route
	department := api.Group("/departments", authMiddleware)

	department.Get("/", departmentController.GetAll)
	department.Post("/", departmentController.Create)
	department.Get("/:departmentName", departmentController.GetByName)
	department.Patch("/:departmentName", departmentController.Update)
	department.Delete("/:departmentName", departmentController.Delete)

	// score route
	score := api.Group("/scores", authMiddleware)

	score.Get("/", scoreController.GetAll)
	score.Post("/", scoreController.CreateMany)
	score.Get("/:scoreId", scoreController.GetById)
	score.Patch("/:scoreId", scoreController.Update)
	score.Delete("/:scoreId", scoreController.Delete)

	// user route
	user := api.Group("/users", authMiddleware)

	user.Get("/", userController.GetAll)
	user.Post("/", userController.Create)
	user.Get("/:userId", userController.GetById)
	user.Patch("/:userId", userController.Update)
	user.Delete("/:userId", userController.Delete)
	user.Post("/bulk", userController.CreateMany)

	user.Get("/:userId/course", courseController.GetByUserId)

	// assignment route
	assignment := api.Group("/assignments", authMiddleware)

	assignment.Get("/", assignmentController.GetAssignments)
	assignment.Post("/", assignmentController.Create)
	assignment.Get("/:assignmentId", assignmentController.GetById)
	assignment.Patch("/:assignmentId", assignmentController.Update)
	assignment.Delete("/:assignmentId", assignmentController.Delete)
	assignment.Get("/:assignmentId/scores", scoreController.GetByAssignmentId)

	assignmentGroup := api.Group("/assignment-groups", authMiddleware)
	assignmentGroup.Post("/", assignmentController.CreateGroup)
	assignmentGroup.Patch("/:assignmentGroupId", assignmentController.UpdateGroup)
	assignmentGroup.Delete("/:assignmentGroupId", assignmentController.DeleteGroup)

	// clo by assignment route
	cloByAssignment := assignment.Group("/:assignmentId/clos/", authMiddleware)
	cloByAssignment.Post("/", assignmentController.CreateLinkCourseLearningOutcome)
	cloByAssignment.Delete("/:cloId", assignmentController.DeleteLinkCourseLearningOutcome)

	// programme route
	programme := api.Group("/programmes", authMiddleware)

	programme.Get("/", programmeController.GetAll)
	programme.Post("/", programmeController.Create)
	programme.Get("/:programmeName", programmeController.GetByName)
	programme.Patch("/:programmeName", programmeController.Update)
	programme.Delete("/:programmeName", programmeController.Delete)

	// enrollment route
	enrollment := api.Group("/enrollments", authMiddleware)

	enrollment.Get("/", enrollmentController.GetAll)
	enrollment.Post("/", enrollmentController.Create)
	enrollment.Get("/:enrollmentId", enrollmentController.GetById)
	enrollment.Patch("/:enrollmentId", enrollmentController.Update)
	enrollment.Delete("/:enrollmentId", enrollmentController.Delete)

	// semester route
	semester := api.Group("/semesters", authMiddleware)

	semester.Get("/", semesterController.GetAll)
	semester.Get("/:semesterId", semesterController.GetById)
	semester.Post("/", semesterController.Create)
	semester.Patch("/:semesterId", semesterController.Update)
	semester.Delete("/:semesterId", semesterController.Delete)

	// grade route
	grade := api.Group("/grades")

	grade.Get("/", gradeController.GetAll)
	grade.Post("/", gradeController.CreateMany)
	grade.Get("/:gradeId", gradeController.GetById)
	grade.Patch("/:gradeId", gradeController.Update)
	grade.Delete("/:gradeId", gradeController.Delete)

	// course stream route
	courseStream := api.Group("/course-streams")
	courseStream.Get("/", courseStreamController.Get)
	courseStream.Post("/", courseStreamController.Create)
	courseStream.Delete("/:courseStreamId", courseStreamController.Delete)

	// prediction
	prediction := api.Group("/prediction")

	prediction.Get("/Train", predictionController.Train)

	// authentication route
	auth := app.Group("/auth")

	auth.Post("/login", authController.SignIn)
	auth.Get("/logout", authMiddleware, authController.SignOut)
	auth.Get("/me", authMiddleware, authController.Me)

	app.Get("/metrics", monitor.New())

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	err := app.Listen(":3001")

	return err
}
