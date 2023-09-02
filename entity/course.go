package entity

type Course struct {
	ID         string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Year       int    `json:"year"`
	LecturerID string `db:"lecturer_id" json:"lecturer_id"`

	Lecturer Lecturer
}
