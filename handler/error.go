package handler

import (
	"html/template"
	"log"
	"net/http"
	"os"
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

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request, err error) {
	ErrorHandler(w, http.StatusMethodNotAllowed, "Method "+r.Method+" is not allowed", tmplError)
	WriteLogs(err)
}

func BadRequestHandler(w http.ResponseWriter, r *http.Request, err error) {
	ErrorHandler(w, http.StatusBadRequest, "Bad Request", tmplError)
	WriteLogs(err)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	ErrorHandler(w, http.StatusInternalServerError, "Internal Server Error", tmplError)
	WriteLogs(err)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request, err error) {
	ErrorHandler(w, http.StatusNotFound, "Page not found", tmplError)
	WriteLogs(err)
}

func WriteLogs(err error) error {
	// If the file doesn't exist, create it, or append to the file
	f, fileerr := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if fileerr != nil {
		log.Println(fileerr)
	}
	if _, err := f.Write([]byte(err.Error() + "\n")); err != nil {
		log.Println(err)
	}
	if err := f.Close(); err != nil {
		log.Println(err)
	}
	return err
}
