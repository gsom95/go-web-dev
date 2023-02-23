package view

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
)

// Template is used as "view" object in our MVC-like app architecture.
// It contains methods that will be used by controllers.
type Template struct {
	htmlTpl *template.Template
}

// Execute executes the template and writes result to the http.ResponseWriter.
func (t *Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template.", http.StatusInternalServerError)
		return
	}
}

// Parse gets a template from filepath and tries to parse it.
func Parse(filepath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
	}, nil

}

// ParseFS tries to parse template from embedded FS.
func ParseFS(fs fs.FS, pattern string) (Template, error) {
	htmlTpl, err := template.ParseFS(fs, pattern)
	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}
	return Template{
		htmlTpl: htmlTpl,
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
