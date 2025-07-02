package handler

import (
	"errors"
	"html/template"
	"net/http"
)

var tmpl *template.Template

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		err := errors.New(r.Method + "method not allowed")
		MethodNotAllowedHandler(w, r, err)
		return
	}
	pageData, err := GetPageData()
	if err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, pageData); err != nil {
		InternalServerErrorHandler(w, r, err)
		return
	}
}
