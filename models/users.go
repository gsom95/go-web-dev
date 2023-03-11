package models

import "database/sql"

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

// UserService will interact with database.
type UserService struct {
	DB *sql.DB
}
