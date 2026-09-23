package dto

type User struct {
	Name     string
	Email    string
	Password string
}

type Login struct {
	Email    string
	Password string
}
