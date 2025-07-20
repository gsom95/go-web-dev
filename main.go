package main

import (
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/gsom95/go-web-dev/controllers"
	"github.com/gsom95/go-web-dev/views"
)

func main() {
	r := chi.NewRouter()

	homeTpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(homeTpl))

	contactTpl, err := views.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	faqTpl, err := views.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(faqTpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	slog.Info("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", r)
}
