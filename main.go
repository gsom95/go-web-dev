package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gsom95/go-web-dev/controllers"
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

	var usersCtrl controllers.Users
	usersCtrl.Templates.New = view.Must(view.ParseFS(
		templates.FS, "signup.gohtml", tailwind,
	))
	r.Get("/signup", usersCtrl.New)

	log.Println("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", r)
}
