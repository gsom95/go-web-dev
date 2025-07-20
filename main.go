package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/gsom95/go-web-dev/controllers"
	"github.com/gsom95/go-web-dev/views"
)

func faqHandler(w http.ResponseWriter, r *http.Request) {
	executeTemplate(w, filepath.Join("templates", "faq.gohtml"))
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func executeTemplate(w http.ResponseWriter, templatePath string) {
	logger := slog.With(slog.String("filepath", templatePath))

	tpl, err := views.Parse(templatePath)
	if err != nil {
		logger.Error("cannot parse a template file", slog.String("error", err.Error()))
		http.Error(w, "cannot parse a template file", http.StatusInternalServerError)

		return
	}

	tpl.Execute(w, nil)
}

func main() {
	r := chi.NewRouter()

	homeTpl, err := views.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		panic(err)
	}

	r.Get("/", controllers.Static{Template: homeTpl}.ServeHTTP)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", r)
}
