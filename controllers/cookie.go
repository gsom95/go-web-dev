package controllers

import (
	"fmt"
	"net/http"
)

const (
	// CookieSession stores the name of our session cookie.
	CookieSession string = "session"
)

// newCookie creates a new cookie.
// This will help ensure that we are creating cookies with the same secure settings every time.
func newCookie(name, value string) *http.Cookie {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	return &cookie
}

// In most cases, we will want to set a cookie immediately after creating it.
func setCookie(w http.ResponseWriter, name, value string) {
	cookie := newCookie(name, value)
	http.SetCookie(w, cookie)
}

func readCookie(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("cookie %s: %w", name, err)
	}
	return c.Value, nil
}
