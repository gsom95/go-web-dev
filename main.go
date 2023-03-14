package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gsom95/go-web-dev/controllers"
	"github.com/gsom95/go-web-dev/models"
	"github.com/gsom95/go-web-dev/templates"
	"github.com/gsom95/go-web-dev/view"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	tailwind := "tailwind.gohtml"

	r.Get("/", controllers.StaticHandler(
		view.Must(view.ParseFS(templates.FS, "home.gohtml", tailwind)),
	))
	r.Get("/contact", controllers.StaticHandler(
		view.Must(view.ParseFS(templates.FS, "contact.gohtml", tailwind)),
	))

	r.Get("/faq", controllers.FAQ(
		view.Must(view.ParseFS(templates.FS, "faq.gohtml", tailwind)),
	))

	// Setup a db connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var usersCtrl controllers.Users
	usersCtrl.Service = &models.UserService{
		DB: db,
	}

	usersCtrl.Templates.New = view.Must(view.ParseFS(
		templates.FS, "signup.gohtml", tailwind,
	))
	usersCtrl.Templates.SignIn = view.Must(view.ParseFS(
		templates.FS, "signin.gohtml", tailwind,
	))

	r.Get("/users/new", usersCtrl.New)
	r.Post("/signup", usersCtrl.Create)

	r.Get("/signin", usersCtrl.SignIn)
	r.Post("/signin", usersCtrl.ProcessSignIn)

	log.Println("Starting the server on :3000...")
	log.Println(http.ListenAndServe(":3000", r))
}
