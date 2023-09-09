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
	UserByName(username string) (usr *User, err error)
	UserByEmail(email string) (usr *User, err error)
	UserByID(userID int) (usr *User, err error)
	DeleteUser(usr *User) error
	UpdateUser(usr *User) error
}
