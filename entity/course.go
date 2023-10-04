package entity

type Course struct {
	ID         string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	SemesterID string `db:"semester_id" json:"semester_id"`
	LecturerID string `db:"lecturer_id" json:"lecturer_id"`

	Semester Semester
	Lecturer Lecturer
}
