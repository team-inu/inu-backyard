package entity

type Lecturer struct {
	Id        string `json:"id" gorm:"primaryKey;type:char(255)"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Role      string `json:"role" gorm:"default:'lecturer'"`
}

type LecturerRepository interface {
	GetAll() ([]Lecturer, error)
	GetById(id string) (*Lecturer, error)
	GetByParams(params *Lecturer, limit int, offset int) ([]Lecturer, error)
	GetByEmail(email string) (*Lecturer, error)
	GetBySessionId(sessionId string) (*Lecturer, error)
	Create(lecturer *Lecturer) error
	Update(id string, lecturer *Lecturer) error
	Delete(id string) error
}

type LecturerUseCase interface {
	GetAll() ([]Lecturer, error)
	GetById(id string) (*Lecturer, error)
	GetByParams(params *Lecturer, limit int, offset int) ([]Lecturer, error)
	GetByEmail(email string) (*Lecturer, error)
	Create(name string, firstName string, lastName string) error
	Update(id string, lecturer *Lecturer) error
	Delete(id string) error
	GetBySessionId(sessionId string) (*Lecturer, error)
}
