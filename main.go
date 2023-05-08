package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/csrf"
	"github.com/gsom95/go-web-dev/controllers"
	"github.com/gsom95/go-web-dev/migrations"
	"github.com/gsom95/go-web-dev/models"
	"github.com/gsom95/go-web-dev/templates"
	"github.com/gsom95/go-web-dev/view"
)

func main() {
	// Setup a db connection
	cfg := models.DefaultPostgresConfig()
	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Change the following line of code
	err = models.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	r := chi.NewRouter()

	csrfKey := "gFvi45R4fy5xNBlnEeZtQbfAVCYEIAUX"
	csrfMw := csrf.Protect(
		[]byte(csrfKey),
		// TODO: Fix this before deploying
		csrf.Secure(false),
	)
	r.Use(middleware.Logger, middleware.Recoverer, csrfMw)

	sessionService := &models.SessionService{
		DB: db,
	}
	userMw := controllers.UserMiddleware{
		SessionService: sessionService,
	}
	r.Use(userMw.SetUser)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	tailwind := "tailwind.gohtml"

	r.Get("/", controllers.StaticHandler(
		view.Must(view.ParseFS(templates.FS, "home.gohtml", tailwind)),
	))
	r.Get("/contact", controllers.StaticHandler(
		view.Must(view.ParseFS(templates.FS, "contact.gohtml", tailwind)),
	))

	r.Get("/faq", controllers.FAQ(
		view.Must(view.ParseFS(templates.FS, "faq.gohtml", tailwind)),
	))

	var usersCtrl controllers.Users
	usersCtrl.UserService = &models.UserService{
		DB: db,
	}

	usersCtrl.Templates.New = view.Must(view.ParseFS(
		templates.FS, "signup.gohtml", tailwind,
	))
	usersCtrl.Templates.SignIn = view.Must(view.ParseFS(
		templates.FS, "signin.gohtml", tailwind,
	))

	usersCtrl.SessionService = sessionService

	r.Get("/signup", usersCtrl.New)
	r.Post("/signup", usersCtrl.Create)

	r.Route("/users/me", func(r chi.Router) {
		r.Use(userMw.RequireUser)
		r.Get("/", usersCtrl.CurrentUser)
	})

	r.Get("/signin", usersCtrl.SignIn)
	r.Post("/signin", usersCtrl.ProcessSignIn)
	r.Post("/signout", usersCtrl.ProcessSignOut)

	log.Println("Starting the server on :3000...")
	log.Println(http.ListenAndServe(":3000", r))
}
