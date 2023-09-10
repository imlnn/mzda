package mzda

type role int

const (
	USER role = iota
	ADMIN
)

type User struct {
	ID       int
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
	Email    string `json:"email"`
	Role     role   `json:"role"`
}

type UserDTO struct {
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
	Email    string `json:"email"`
}

type UserStorage interface {
	AddUser(usr *UserDTO) error
	UserByName(username string) (*User, error)
	UserByEmail(email string) (*User, error)
	UserByID(userID int) (*User, error)
	DeleteUser(usr *User) error
	UpdateUser(usr *User) error
}
