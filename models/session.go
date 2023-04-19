package models

// Session is a one-to-one mapping of sessions table.
type Session struct {
	ID        int
	UserID    int
	TokenHash string
}
