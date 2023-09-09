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

type UserStorage interface {
	AddUser(usr *User) error
	UserByName(username string) (usr *User, err error)
	UserByID(userID int) (usr *User, err error)
	DeleteUser(usr *User) error
	UpdateUser(usr *User) error
}
