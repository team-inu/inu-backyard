package entity

type Programme struct {
	Name string `json:"name" gorm:"primaryKey;type:char(255)"`
}
