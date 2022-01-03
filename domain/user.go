package domain

type Role uint

const (
	Admin Role = iota
	Regular
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}
