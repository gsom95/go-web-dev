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

	tpl, err := view.Parse(filepath.Join("templates", "home.gohtml"))
	if err != nil {
		log.Fatalln(err)
	}
	r.Get("/", controllers.StaticHandler(tpl))

	tpl, err = view.Parse(filepath.Join("templates", "contact.gohtml"))
	if err != nil {
		log.Fatalln(err)
	}
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl, err = view.Parse(filepath.Join("templates", "faq.gohtml"))
	if err != nil {
		log.Fatalln(err)
	}
	r.Get("/faq", controllers.StaticHandler(tpl))

	log.Println("Starting the server on :3000...")
	_ = http.ListenAndServe(":3000", r)
}
