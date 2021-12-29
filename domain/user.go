package domain

type Role uint

const (
	Admin Role = iota
	Regular
)

type User struct {
	ID       string
	Username string
	Password string
	Role     Role
}
