package entity

type Faculty struct {
	Name string `json:"name" gorm:"primaryKey;type:char(255)"`
}
