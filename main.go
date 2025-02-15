package main

import (
	"fmt"
	"net/http"
)

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

func pathHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found", http.StatusNotFound)
	}
}

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
	var _ Router
	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", http.HandlerFunc(pathHandler))
}
