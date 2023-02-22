package controllers

import (
	"net/http"

	"github.com/gsom95/go-web-dev/view"
)

// Static controller serves static web pages from a template.
type Static struct {
	Template view.Template
}

// ServeHttp implements http.Handler interface.
func (s Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Template.Execute(w, nil)
}
