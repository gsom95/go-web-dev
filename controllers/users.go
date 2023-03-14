package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gsom95/go-web-dev/models"
)

// Users will be used by HTTP handlers.
type Users struct {
	Templates struct {
		New    Template
		SignIn Template
	}
	Service *models.UserService
}

// New renders signup page.
func (u Users) New(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.New.Execute(w, data)
}

// Create handles creation of a user
func (u Users) Create(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	password := r.FormValue("password")
	user, err := u.Service.Create(email, password)
	if err != nil {
		log.Println("error creaing new user:", err.Error())
		http.Error(w, "Something went wrong while creating a user", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "User created: %+v", user)
}

// Signin renders sign in page.
func (u Users) SignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email string
	}
	data.Email = r.FormValue("email")
	u.Templates.SignIn.Execute(w, data)
}
