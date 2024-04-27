package entity

type CourseStreamType string

const (
	UpCourseStreamType   CourseStreamType = "UpStream"
	DownCourseStreamType CourseStreamType = "DownStream"
)

type CourseStream struct {
	Id             string
	FromCourseId   string
	TargetCourseId string
	StreamType     CourseStreamType
	Comment        string
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
