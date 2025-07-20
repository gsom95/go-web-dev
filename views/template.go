package views

import (
	"fmt"
	"html/template"
	"log/slog"
	"net/http"
)

type Template struct {
	htmlTpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err := t.htmlTpl.Execute(w, data)
	if err != nil {
		slog.Error("cannot execute template", slog.String("error", err.Error()))
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func Parse(filepath string) (Template, error) {
	htmlTpl, err := template.ParseFiles(filepath)
	if err != nil {
		return Template{}, fmt.Errorf("cannot parse template: %w", err)
	}

	return Template{
		htmlTpl: htmlTpl,
	}, nil
}
