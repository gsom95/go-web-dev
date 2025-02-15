package main

import (
	"fmt"
	"net/http"

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
	fmt.Fprint(w, "Welcome to the contact page!")
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the home page!")
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
