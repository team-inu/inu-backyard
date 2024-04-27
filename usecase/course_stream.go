package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
	errs "github.com/team-inu/inu-backyard/entity/error"
)

type courseStreamUseCase struct {
	courseStreamRepository entity.CourseStreamRepository
	courseUseCase          entity.CourseUseCase
}

func NewCourseStreamUseCase(
	courseStreamRepository entity.CourseStreamRepository,
	courseUseCase entity.CourseUseCase,
) entity.CourseStreamsUseCase {
	return &courseStreamUseCase{
		courseStreamRepository: courseStreamRepository,
		courseUseCase:          courseUseCase,
	}
}

func (u *courseStreamUseCase) Create(fromCourseId string, targetCourseId string, streamType entity.CourseStreamType, comment string) error {
	fromCourse, err := u.courseUseCase.GetById(fromCourseId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot validate from course id %s while creating stream course", targetCourseId, err)
	} else if fromCourse == nil {
		return errs.New(errs.ErrCreateCourseStream, "from course id: %s not found", fromCourseId)
	}

	targetCourse, err := u.courseUseCase.GetById(targetCourseId)
	if err != nil {
		return errs.New(errs.SameCode, "cannot validate target course id %s while creating stream course", targetCourseId, err)
	} else if targetCourse == nil {
		return errs.New(errs.ErrCreateCourseStream, "target course id: %s not found", targetCourseId)
	}

	courseStream := entity.CourseStream{
		Id:             ulid.Make().String(),
		FromCourseId:   fromCourseId,
		TargetCourseId: targetCourseId,
		StreamType:     streamType,
		Comment:        comment,
	}

	err = u.courseStreamRepository.Create(&courseStream)
	if err != nil {
		return errs.New(errs.ErrCreateCourseStream, "cannot create stream course", err)
	}

	return nil
}

func (u *courseStreamUseCase) Delete(id string) error {
	err := u.courseStreamRepository.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCourseStream, "cannot delete stream course", err)
	}

	return nil
}

func (u *courseStreamUseCase) GetByTargetCourseId(courseId string) ([]entity.CourseStream, error) {
	course, err := u.courseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot validate course id %s while getting stream course", courseId, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrQueryCourseStream, "course id %s not found while getting stream course", courseId)
	}

	courseStream, err := u.courseStreamRepository.GetByQuery(entity.CourseStream{TargetCourseId: courseId})
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get stream course of course id: %s", courseStream, err)
	}

	return courseStream, nil
}

func (u *courseStreamUseCase) GetByFromCourseId(courseId string) ([]entity.CourseStream, error) {
	course, err := u.courseUseCase.GetById(courseId)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot validate course id %s while getting stream course", courseId, err)
	} else if course == nil {
		return nil, errs.New(errs.ErrQueryCourseStream, "course id %s not found while getting stream course", courseId)
	}

	courseStream, err := u.courseStreamRepository.GetByQuery(entity.CourseStream{FromCourseId: courseId})
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get stream course of course id: %s", courseStream, err)
	}

	return courseStream, nil
}

func (u *courseStreamUseCase) Get(id string) (*entity.CourseStream, error) {
	courseStream, err := u.courseStreamRepository.Get(id)
	if err != nil {
		return nil, errs.New(errs.SameCode, "cannot get stream course id %s", courseStream, err)
	}

	return courseStream, nil
}

func (u *courseStreamUseCase) Update(id string, comment string) error {
	courseStream, err := u.Get(id)
	if err != nil {
		return errs.New(errs.SameCode, "cannot validate stream course id %s to delete", id, err)
	} else if courseStream == nil {
		return errs.New(errs.ErrDeleteCourseStream, "stream course id %s to delete not found", id)
	}

	err = u.Delete(id)
	if err != nil {
		return errs.New(errs.ErrDeleteCourseStream, "cannot delete stream course id %s", id)
	}

	return nil
}
