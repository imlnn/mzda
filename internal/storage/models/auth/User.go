package auth

type role int

const (
	USER role = iota
	ADMIN
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Pwd      string `json:"pwd"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Role     role   `json:"role"`
}
