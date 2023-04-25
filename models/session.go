package models

import (
	"database/sql"
	"fmt"

	"github.com/gsom95/go-web-dev/rand"
)

// Session is a one-to-one mapping of sessions table.
type Session struct {
	ID     int
	UserID int
	// Token is only set when creating a new session. When looking up a session
	// this will be left empty, as we only store the hash of a session token
	// in our database and we cannot reverse it into a raw token.
	Token     string
	TokenHash string
}

// SessionService is a service for storing and fetching session info.
type SessionService struct {
	DB *sql.DB
}

// Create will create a new session for the user provided. The session token
// will be returned as the Token field on the Session type, but only the hashed
// session token is stored in the database.
func (ss *SessionService) Create(userID int) (*Session, error) {
	token, err := rand.SessionToken()
	if err != nil {
		return nil, fmt.Errorf("SessionService.Create: %w", err)
	}
	// TODO: hash session token
	session := Session{
		UserID: userID,
		Token:  token,
		// TODO: Set the TokenHash
	}
	// TODO: Store the session in our DB
	return &session, nil
}

// User queries a user via raw token using the SessionService.
func (ss *SessionService) User(token string) (*User, error) {
	// TODO: Implement SessionService.User
	return nil, nil
}
