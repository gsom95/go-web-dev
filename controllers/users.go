package controllers

import (
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
