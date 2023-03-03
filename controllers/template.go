package controllers

import "net/http"

// Templates defines what a template should do. We're doing it for decoupling our code.
type Template interface {
	Execute(w http.ResponseWriter, data any)
}
