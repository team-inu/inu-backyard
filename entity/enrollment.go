package entity

type Enrollment struct {
	ID        string `json:"id" gorm:"primaryKey;type:char(255)"`
	CourseID  string `json:"course_id"`
	StudentID string `json:"student_id"`

	Course  Course
	Student Student
}
