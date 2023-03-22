package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
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
		Email     string
		CSRFField template.HTML
	}
	data.Email = r.FormValue("email")

	// give us the HTML for a hidden <input> tag that has the CSRF token
	// for the incoming request. We assign this to the CSRFField field of
	// the data struct so that it will be available inside of our template.
	data.CSRFField = csrf.TemplateField(r)
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

// ProcessSignIn processes data from sign in form.
func (u Users) ProcessSignIn(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Email    string
		Password string
	}
	data.Email = r.FormValue("email")
	data.Password = r.FormValue("password")
	user, err := u.Service.Authenticate(data.Email, data.Password)
	if err != nil {
		log.Println(err)
		http.Error(w, "Something went wrong.", http.StatusInternalServerError)
		return
	}
	cookie := http.Cookie{
		Name:  "email",
		Value: user.Email,

		// If we want a cookie to be accessible from any page on our website,
		// we provide a Path value of '/' which maps to any path on our website.
		Path: "/",

		// To disable JavaScript access to cookies.
		HttpOnly: true,
	}
	http.SetCookie(w, &cookie)

	fmt.Fprintf(w, "User authenticated: %+v", user)
}
