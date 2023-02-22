package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gsom95/go-web-dev/controllers"
	"github.com/gsom95/go-web-dev/view"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	tpl, err := view.Parse(filepath)
	if err != nil {
		log.Printf("processing template: %v", err)
		http.Error(w, "There was an error processing the template.", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "faq.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func getGallery(w http.ResponseWriter, r *http.Request) {
	galleryID := chi.URLParam(r, "galleryID")
	fmt.Fprintf(w, "Got id=%s", galleryID)
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger, middleware.Recoverer)

	homeTpl, err := view.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		log.Fatalln(err)
	}
	r.Method(http.MethodGet, "/", controllers.Static{Template: homeTpl})
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
