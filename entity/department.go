package entity

type Department struct {
	Name        string `json:"name" gorm:"type:char(255);unique;not null;primaryKey"`
	FacultyName string `json:"faculty_name"`

	Faculty Faculty `gorm:"foreignKey:FacultyName"`
}
