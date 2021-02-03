package models

type User struct {
	ID    int64  `json:"id"`
	Phone string `json:"phone"`
	Email string `json:"email"`
}

var Users = []User{
	{ID: 1, Phone: "998996096060", Email: "fprotimaru@gmail.com"},
}
