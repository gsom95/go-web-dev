package controllers

import (
	"net/http"

	"github.com/gsom95/go-web-dev/views"
)

type Static struct {
	Template views.Template
}

func (s Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Template.Execute(w, nil)
}
