package controllers

import (
	"fmt"
	"net/http"
)

// Users will be used by HTTP handlers.
type Users struct {
	Templates struct {
		New Template
	}
}

// New shows signup page.
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}

// Create handles creation of a user
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	// for now r.FormValue would be enough because it calls r.ParseForm()
	fmt.Fprintf(w, "<p>Email: %s</p>\n", r.FormValue("email"))
	fmt.Fprintf(w, "<p>Password: %s</p>\n", r.FormValue("password"))
}
