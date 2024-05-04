package entity

type UserRole string

const (
	UserRoleLecturer         UserRole = "LECTURER"
	UserRoleModerator        UserRole = "MODERATOR"
	UserRoleHeadOfCurriculum UserRole = "HEAD_OF_CURRICULUM"
	UserRoleTABEEManager     UserRole = "TABEE_MANAGER"
)

type User struct {
	Id        string   `json:"id" gorm:"primaryKey;type:char(255)"`
	Email     string   `json:"email" gorm:"unique"`
	Password  string   `json:"password"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Role      UserRole `json:"role" gorm:"default:'LECTURER'"`
}

func (u User) IsRoles(expectedRoles []UserRole) bool {
	for _, expectedRole := range expectedRoles {
		if u.Role == expectedRole {
			return true
		}
	}

	return false
}

type UserRepository interface {
	GetAll() ([]User, error)
	GetById(id string) (*User, error)
	GetByParams(params *User, limit int, offset int) ([]User, error)
	GetByEmail(email string) (*User, error)
	GetBySessionId(sessionId string) (*User, error)
	Create(user *User) error
	CreateMany(users []User) error
	Update(id string, user *User) error
	Delete(id string) error
}

type UserUseCase interface {
	GetAll() ([]User, error)
	GetById(id string) (*User, error)
	GetByParams(params *User, limit int, offset int) ([]User, error)
	GetByEmail(email string) (*User, error)
	Create(name string, firstName string, lastName string, password string, role UserRole) error
	CreateMany(users []User) error
	Update(id string, user *User) error
	Delete(id string) error
	GetBySessionId(sessionId string) (*User, error)
}
