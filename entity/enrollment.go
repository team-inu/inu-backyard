package entity

import "github.com/oklog/ulid/v2"

type Enrollment struct {
	ID        ulid.ULID `json:"id" gorm:"primaryKey;type:char(255)"`
	CourseID  ulid.ULID `json:"course_id"`
	StudentID ulid.ULID `json:"student_id"`

	Course  Course
	Student Student
}
