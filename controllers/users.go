package controllers

import (
	"net/http"

	"github.com/gsom95/go-web-dev/view"
)

// Users will be used by HTTP handlers.
type Users struct {
	Templates struct {
		New view.Template
	}
}

// New shows signup page.
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	u.Templates.New.Execute(w, nil)
}
