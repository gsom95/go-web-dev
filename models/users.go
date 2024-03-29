package models

import (
	"database/sql"
	"fmt"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// User struct represents data stored in the database.
type User struct {
	ID           int
	Email        string
	PasswordHash string
}

// UserService will interact with database.
type UserService struct {
	DB *sql.DB
}

// Create handles creation and adding of a new user to the database.
func (us *UserService) Create(email, password string) (*User, error) {
	email = strings.ToLower(email)

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	passwordHash := string(hashedBytes)
	user := User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	row := us.DB.QueryRow(`
	INSERT INTO users (email, password_hash)
	VALUES ($1, $2) RETURNING id;`, user.Email, user.PasswordHash)
	err = row.Scan(&user.ID)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}

	return &user, nil
}

// Authenticate is used for user authentication.
func (us UserService) Authenticate(email, password string) (*User, error) {
	user := User{
		Email: strings.ToLower(email),
	}
	row := us.DB.QueryRow(`
	SELECT id, password_hash
	FROM users WHERE email = $1`, email)

	// probably, we should also check if no user was found with email
	// TODO: add check for "no user found" case
	err := row.Scan(&user.ID, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("authenticate: %w", err)
	}
	return &user, nil
}
