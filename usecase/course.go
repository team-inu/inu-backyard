package usecase

import (
	"github.com/oklog/ulid/v2"
	"github.com/team-inu/inu-backyard/entity"
)

type courseUsecase struct {
	courseRepo entity.CourseRepository
}

func NewCourseUsecase(courseRepo entity.CourseRepository) entity.CourseUsecase {
	return &courseUsecase{courseRepo: courseRepo}
}

func (c courseUsecase) GetAll() ([]entity.Course, error) {
	return c.courseRepo.GetAll()
}

func (c courseUsecase) GetByID(id string) (*entity.Course, error) {
	return c.courseRepo.GetByID(id)
}

func (c courseUsecase) Create(name string, code string, year int, lecturerId string) (*entity.Course, error) {
	course := entity.Course{
		ID:         ulid.Make().String(),
		Name:       name,
		Code:       code,
		LecturerID: lecturerId,
	}

	err := c.courseRepo.Create(&course)

	if err != nil {
		return nil, err
	}

	return &course, nil
}

func (c courseUsecase) Delete(id string) error {
	return c.courseRepo.Delete(id)
}
