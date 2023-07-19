package entity

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func NewUser() *User {
	return &User{}
}
