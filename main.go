package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// this can determine how a browser would render response body
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, `<h1>FAQ Page</h1>
<ul>
	<li>
	<b>What are you doing?</b>
	I'm trying to learn how to write backend web applications using Go!
	</li>
	<li>
	<b>Yeah, but why exactly you decided to do it?</b>
	Well, because I want to be sure that I know all the basic and required stuff. There's lots of nuances in writing backend server apps.
	</li>
</ul>
	`)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:jon@calhoun.io\">jon@calhoun.io</a>.</p>")
}

func getGallery(w http.ResponseWriter, r *http.Request) {
	galleryID := chi.URLParam(r, "galleryID")
	fmt.Fprintf(w, "Got id=%s", galleryID)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.Get("/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("at the disco!")
	})
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})
	r.Route("/galleries", func(r chi.Router) {
		r.Get("/{galleryID}", getGallery)
	})

	fmt.Println("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", r)
}
