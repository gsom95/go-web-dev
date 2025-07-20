package main

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
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

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tpl, err := template.ParseFiles(templatePath)
	if err != nil {
		logger.Error("cannot parse a template file",
			slog.String("error", err.Error()),
		)
		http.Error(w, "cannot parse a template file", http.StatusInternalServerError)
		return
	}
	if err = tpl.Execute(w, nil); err != nil {
		logger.Error("cannot execute a template",
			slog.String("error", err.Error()),
		)
		http.Error(w, "cannot execute a template file", http.StatusInternalServerError)
		return
	}

	logger.Debug("request completed successfully")
}

func main() {
	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", r)
}
