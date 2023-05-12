package view

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/gsom95/go-web-dev/context"
	"github.com/gsom95/go-web-dev/models"
)

// Template is used as "view" object in our MVC-like app architecture.
// It contains methods that will be used by controllers.
type Template struct {
	htmlTpl *template.Template
}

// Execute executes the template and writes result to the http.ResponseWriter.
func (t Template) Execute(w http.ResponseWriter, r *http.Request, data any) {
	// Use the Clone method to clone our template before making any changes to it.
	tpl, err := t.htmlTpl.Clone()
	if err != nil {
		log.Printf("cloning template: %v", err)
		http.Error(w, "There was an error rendering the page.", http.StatusInternalServerError)
		return
	}
	tpl = tpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
			"currentUser": func() *models.User {
				return context.User(r.Context())
			},
		},
	)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// We are using the bytes.Buffer type to store the final results from the template execution,
	// then once we know the entire template has been executed without any errors we then
	// call io.Copy to copy the data from our buffer into the http.ResponseWriter.
	// This way we avoid "http: superfluous response.WriteHeader call" error.
	var buf bytes.Buffer
	err = tpl.Execute(&buf, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
	io.Copy(w, &buf)
}

// ParseFS tries to parse template from embedded FS.
func ParseFS(fs fs.FS, patterns ...string) (Template, error) {
	tpl := template.New(patterns[0])
	tpl = tpl.Funcs(
		template.FuncMap{
			// error will indicate that function is not implemented
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("csrfField not implemented")
			},
			"currentUser": func() (*models.User, error) {
				return nil, fmt.Errorf("currentUser not implemented")
			},
		},
	)
	tpl, err := tpl.ParseFS(fs, patterns...)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: tpl,
	}, nil
}

// Must is only used when starting up our application and parsing templates for the first time,
// and we know that almost all errors that occur during a call to Parse are developer errors.
func Must(t Template, err error) Template {
	if err != nil {
		panic(err)
	}
	return t
}
