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

	homeTpl := views.MustParse(filepath.Join("templates", "home.gohtml"))
	r.Get("/", controllers.StaticHandler(homeTpl))

	contactTpl := views.MustParse(filepath.Join("templates", "contact.gohtml"))
	r.Get("/contact", controllers.StaticHandler(contactTpl))

	faqTpl := views.MustParse(filepath.Join("templates", "faq.gohtml"))
	r.Get("/faq", controllers.StaticHandler(faqTpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	slog.Info("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", r)
}
