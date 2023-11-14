package entity

type Lecturer struct {
	ID        string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name      string `json:"name"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role"`
}

type LecturerRepository interface {
	GetAll() ([]Lecturer, error)
	GetByID(id string) (*Lecturer, error)
	Create(lecturer *Lecturer) error
	Update(lecturer *Lecturer) error
	Delete(id string) error
}

type LecturerUseCase interface {
	GetAll() ([]Lecturer, error)
	GetByID(id string) (*Lecturer, error)
	Create(name string, firstName string, lastName string) (*Lecturer, error)
	Update(lecturer *Lecturer) error
	Delete(id string) error
}
