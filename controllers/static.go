package controllers

import (
	"net/http"

	"github.com/gsom95/go-web-dev/view"
)

// StaticHandler returns a HandlerFunc for serving static web pages.
func StaticHandler(tpl view.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}
