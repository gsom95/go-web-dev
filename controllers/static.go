package controllers

import (
	"html/template"
	"net/http"
)

// StaticHandler returns a HandlerFunc for serving static web pages.
func StaticHandler(tpl Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, nil)
	}
}

// FAQ renders FAQ page.
func FAQ(tpl Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "What are you doing?",
			Answer:   "I'm trying to learn how to write backend web applications using Go!",
		},
		{
			Question: "Yeah, but why exactly you decided to do it?",
			Answer:   `Well, because I want to be sure that I know all the basic and required stuff. There's lots of nuances in writing backend server apps.`,
		},
		{
			Question: "Is there a free version?",
			Answer:   "Yes! We offer a free trial for 30 days on any paid plans.",
		},
		{
			Question: "What are your support hours?",
			Answer:   "We have support staff answering emails 24/7, though response times may be a bit slower on weekends.",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, r, questions)
	}
}
