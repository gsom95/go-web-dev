package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"errors"
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

const (
	// The minimum number of bytes to be used for each session token.
	MinBytesPerToken = 32
)

// SessionService is a service for storing and fetching session info.
type SessionService struct {
	DB *sql.DB
	// BytesPerToken is used to determine how many bytes to use when generating
	// each session token. If this value is not set or is less than the
	// MinBytesPerToken const it will be ignored and MinBytesPerToken will be
	// used.
	BytesPerToken int
}

// Create will create a new session for the user provided. The session token
// will be returned as the Token field on the Session type, but only the hashed
// session token is stored in the database.
func (ss *SessionService) Create(userID int) (*Session, error) {
	bytesPerToken := ss.BytesPerToken
	if bytesPerToken < MinBytesPerToken {
		bytesPerToken = MinBytesPerToken
	}
	token, err := rand.String(bytesPerToken)
	if err != nil {
		return nil, fmt.Errorf("SessionService.Create: %w", err)
	}

	session := Session{
		UserID:    userID,
		Token:     token,
		TokenHash: ss.hash(token),
	}

	row := ss.DB.QueryRow(`
		UPDATE sessions
		SET token_hash = $2
		WHERE user_id = $1
    	RETURNING id;`,
		session.UserID, session.TokenHash)
	err = row.Scan(&session.ID)
	if errors.Is(err, sql.ErrNoRows) {
		// If no session exists, we will get ErrNoRows. That means we need to
		// create a session object for that user.

		row = ss.DB.QueryRow(`
		INSERT INTO sessions (user_id, token_hash)
		VALUES ($1, $2)
		RETURNING id;`,
			session.UserID, session.TokenHash)

		err = row.Scan(&session.ID)
	}
	if err != nil {
		return nil, fmt.Errorf("create: %w", err)
	}

	return &session, nil
}

// User queries a user via raw token using the SessionService.
func (ss *SessionService) User(token string) (*User, error) {
	tokenHash := ss.hash(token)

	var user User
	row := ss.DB.QueryRow(`
	SELECT u.id, u.email, u.password_hash
	FROM users u
	JOIN sessions s on u.id = s.user_id
	WHERE token_hash = $1;`, tokenHash)
	err := row.Scan(&user.ID, &user.Email, &user.PasswordHash)
	if err != nil {
		return nil, fmt.Errorf("user: %w", err)
	}

	return &user, nil
}

// Delete handles deleting a session.
func (ss *SessionService) Delete(token string) error {
	tokenHash := ss.hash(token)
	_, err := ss.DB.Exec(`
		DELETE FROM sessions
		WHERE token_hash = $1;`, tokenHash)
	if err != nil {
		return fmt.Errorf("delete: %w", err)
	}
	return nil
}

// hash handles taking in a session token as a string and
// returning the hash of that session token.
func (ss *SessionService) hash(token string) string {
	tokenHash := sha256.Sum256([]byte(token))
	// base64 encode the data into a string
	return base64.URLEncoding.EncodeToString(tokenHash[:])
}
