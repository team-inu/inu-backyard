package entity

type CourseStreamType string

const (
	UpCourseStreamType   CourseStreamType = "UPSTREAM"
	DownCourseStreamType CourseStreamType = "DOWNSTREAM"
)

type CourseStream struct {
	Id             string           `json:"id"`
	StreamType     CourseStreamType `json:"streamType"`
	Comment        string           `json:"comment"`
	FromCourseId   string           `json:"fromCourseId"`
	TargetCourseId string           `json:"targetCourseId"`

	FromCourse   Course `json:"fromCourse" gorm:"foreignKey:FromCourseId"`
	TargetCourse Course `json:"targetCourse" gorm:"foreignKey:TargetCourseId"`
}

type CourseStreamRepository interface {
	Get(id string) (*CourseStream, error)
	GetByQuery(query CourseStream) ([]CourseStream, error)
	Create(courseStream *CourseStream) error
	Update(id string, courseStream *CourseStream) error
	Delete(id string) error
}

type CourseStreamsUseCase interface {
	Get(id string) (*CourseStream, error)
	GetByFromCourseId(courseId string) ([]CourseStream, error)
	GetByTargetCourseId(courseId string) ([]CourseStream, error)
	Create(fromCourseId string, targetCourseId string, streamType CourseStreamType, comment string) error
	Update(id string, comment string) error
	Delete(id string) error
}
