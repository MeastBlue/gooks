package models

type User struct {
	ID        uint   `json:"id" db:"id"`
	Lastname  string `json:"lastname" db:"lastname"`
	Firstname string `json:"firstname" db:"firstname"`
	IsAdmin   bool   `json:"is_admin" db:"is_admin"`
}

type Users []User
