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

	ErrCLONotFound = 20300
	ErrCreateCLO   = 20301
	ErrUpdateCLO   = 20302
	ErrDeleteCLO   = 20303
	ErrQueryCLO    = 20304

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

	ErrLecturerNotFound = 20900
	ErrCreateLecturer   = 20901
	ErrUpdateLecturer   = 20902
	ErrDeleteLecturer   = 20903
	ErrQueryLecturer    = 20904

	ErrAssessmentNotFound = 21000
	ErrCreateAssessment   = 21001
	ErrUpdateAssessment   = 21002
	ErrDeleteAssessment   = 21003
	ErrQueryAssessment    = 21004
)
