package entity

type Lecturer struct {
	ID        string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
