package errs

const (
	SameCode = 0

	ErrInternal   = 1
	ErrRoute      = 2
	ErrFileSystem = 3

	ErrAuthHeader       = 10000
	ErrPayloadValidator = 10001
	ErrBodyParser       = 10002
	ErrQueryParser      = 10003
	ErrParamsParser     = 10004

	// TODO: the error code is temporary, we will rearrange the error code later
	ErrStudentNotFound = 20100
	ErrCreateStudent   = 20101
	ErrUpdateStudent   = 20102
	ErrDeleteStudent   = 20103
	ErrQueryStudent    = 20104

	ErrCourseNotFound = 20200
	ErrCreateCourse   = 20201
	ErrUpdateCourse   = 20202
	ErrDeleteCourse   = 20203
	ErrQueryCourse    = 20204

	ErrCLONotFound  = 20300
	ErrCreateCLO    = 20301
	ErrUpdateCLO    = 20302
	ErrDeleteCLO    = 20303
	ErrQueryCLO     = 20304
	ErrUnLinkSubPLO = 20305

	ErrPLONotFound = 20400
	ErrCreatePLO   = 20401
	ErrUpdatePLO   = 20402
	ErrDeletePLO   = 20403
	ErrQueryPLO    = 20404

	ErrPONotFound = 20500
	ErrCreatePO   = 20501
	ErrUpdatePO   = 20502
	ErrDeletePO   = 20503
	ErrQueryPO    = 20504

	ErrSubPLONotFound = 20600
	ErrCreateSubPLO   = 20601
	ErrUpdateSubPLO   = 20602
	ErrDeleteSubPLO   = 20603
	ErrQuerySubPLO    = 20604

	ErrFacultyNotFound = 20700
	ErrCreateFaculty   = 20701
	ErrUpdateFaculty   = 20702
	ErrDeleteFaculty   = 20703
	ErrQueryFaculty    = 20704

	ErrDepartmentNotFound = 20800
	ErrCreateDepartment   = 20801
	ErrUpdateDepartment   = 20802
	ErrDeleteDepartment   = 20803
	ErrQueryDepartment    = 20804

	ErrAssignmentNotFound = 20900
	ErrCreateAssignment   = 20901
	ErrUpdateAssignment   = 20902
	ErrDeleteAssignment   = 20903
	ErrQueryAssignment    = 20904

	ErrUserNotFound = 21000
	ErrCreateUser   = 21001
	ErrUpdateUser   = 21002
	ErrDeleteUser   = 21003
	ErrQueryUser    = 21004
	ErrUserPassword = 21005

	ErrProgrammeNotFound = 21100
	ErrCreateProgramme   = 21101
	ErrDupName           = 21102
	ErrUpdateProgramme   = 21103
	ErrDeleteProgramme   = 21103
	ErrQueryProgramme    = 21104

	ErrScoreNotFound = 21200
	ErrCreateScore   = 21201
	ErrUpdateScore   = 21202
	ErrDeleteScore   = 21203
	ErrQueryScore    = 21204

	ErrEnrollmentNotFound = 21300
	ErrCreateEnrollment   = 21301
	ErrUpdateEnrollment   = 21302
	ErrDeleteEnrollment   = 21303
	ErrQueryEnrollment    = 21304

	ErrSemesterNotFound = 21400
	ErrCreateSemester   = 21401
	ErrUpdateSemester   = 21402
	ErrDeleteSemester   = 21403
	ErrQuerySemester    = 21404

	ErrGradeNotFound = 21500
	ErrCreateGrade   = 21501
	ErrUpdateGrade   = 21502
	ErrDeleteGrade   = 21503
	ErrQueryGrade    = 21504

	ErrSessionNotFound   = 21600
	ErrCreateSession     = 21601
	ErrUpdateSession     = 21602
	ErrDeleteSession     = 21603
	ErrQuerySession      = 21604
	ErrInvalidSession    = 21605
	ErrSessionExpired    = 21606
	ErrGetSession        = 21607
	ErrSignatureMismatch = 21608
	ErrSessionPrefix     = 21609
	ErrDupSession        = 21610

	ErrorPredictionNotFound = 21700
	ErrorCreatePrediction   = 21701
	ErrorUpdatePrediction   = 21702
)
