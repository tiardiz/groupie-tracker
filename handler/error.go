package handler

import (
	"html/template"
	"net/http"
)

var tmplError *template.Template

func InitTemplates() error {
	var err error
	tmplError, err = template.ParseFiles("templates/error.html")
	return err
}

func ErrorHandler(w http.ResponseWriter, code int, message string, tmpl *template.Template) {
	w.WriteHeader(code)

	data := struct {
		Code    int
		Message string
	}{
		Code:    code,
		Message: message,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, message, code)
	}
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, http.StatusMethodNotAllowed, "Method is not allowed", tmplError)
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, http.StatusBadRequest, "Bad Request", tmplError)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", tmplError)
}
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	ErrorHandler(w, http.StatusNotFound, "Page not found", tmplError)
}
