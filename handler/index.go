package handler

import (
	"html/template"
	"net/http"
)

var tmpl *template.Template

func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		MethodNotAllowedHandler(w, r)
		return
	}
	pageData, err := GetPageData()
	// fmt.Println(pageData)
	if err != nil {
		// fmt.Println("error here 2", err)
		InternalServerErrorHandler(w, r)
		return
	}
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	if err := tmpl.Execute(w, pageData); err != nil {
		// fmt.Println("error here!")
		InternalServerErrorHandler(w, r)
		return
	}
}
