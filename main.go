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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `<h1>Welcome to the FAQ page!</h1>
	<p>Here you will find all the answers to your questions.</p>
	<p>But you won't find them here.</p>
	<p>Because this is just a demo.</p>
	<p>But you can still <a href="/contact">contact us</a> if you want.</p>

	<p>Or you can go back <a href="/">home</a>.</p>
`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "contact.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		slog.Error("cannot parse the contact page template", slog.String("error", err.Error()))
		http.Error(w, "cannot parse the contact page template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		slog.Error("cannot execute the contact page template", slog.String("error", err.Error()))
		http.Error(w, "cannot create a response", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	tplPath := filepath.Join("templates", "home.gohtml")
	tpl, err := template.ParseFiles(tplPath)
	if err != nil {
		slog.Error("cannot parse home page template", slog.String("error", err.Error()))
		http.Error(w, "cannot parse home page template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		slog.Error("cannot execute a home page template", slog.String("error", err.Error()))
		http.Error(w, "cannot create a response", http.StatusInternalServerError)
		return
	}
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
