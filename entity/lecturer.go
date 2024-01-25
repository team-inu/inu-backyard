package entity

type Lecturer struct {
	ID        string `json:"id" gorm:"primaryKey;type:char(255)"`
	Name      string `json:"name"`
	Email     string `json:"email" gorm:"unique"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Role      string `json:"role" gorm:"default:'lecturer'"`
}

type LecturerRepository interface {
	GetAll() ([]Lecturer, error)
	GetByID(id string) (*Lecturer, error)
	GetByParams(params *Lecturer, limit int, offset int) ([]Lecturer, error)
	GetByEmail(email string) (*Lecturer, error)
	GetBySessionId(sessionId string) (*Lecturer, error)
	Create(lecturer *Lecturer) error
	Update(id string, lecturer *Lecturer) error
	Delete(id string) error
}

type LecturerUseCase interface {
	GetAll() ([]Lecturer, error)
	GetByID(id string) (*Lecturer, error)
	GetByParams(params *Lecturer, limit int, offset int) ([]Lecturer, error)
	GetByEmail(email string) (*Lecturer, error)
	Create(name string, firstName string, lastName string) error
	Update(id string, lecturer *Lecturer) error
	Delete(id string) error
	GetBySessionId(sessionId string) (*Lecturer, error)
}
