package controllers

import (
	"net/http"

	"github.com/gsom95/go-web-dev/views"
)

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		tpl.Execute(w, nil)
	}
}
