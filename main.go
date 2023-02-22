package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gsom95/go-web-dev/controllers"
	"github.com/gsom95/go-web-dev/view"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	templatesFolder := "templates"

	r.Get("/", controllers.StaticHandler(
		view.Must(view.Parse(filepath.Join(templatesFolder, "home.gohtml"))),
	))
	r.Get("/contact", controllers.StaticHandler(
		view.Must(view.Parse(filepath.Join(templatesFolder, "contact.gohtml"))),
	))

	r.Get("/faq", controllers.StaticHandler(
		view.Must(view.Parse(filepath.Join(templatesFolder, "faq.gohtml"))),
	))
	))

	log.Println("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", r)
}
