package main

import (
	"log"
	"net/http"
	"text/template"
)

var tmpl *template.Template

func main() {
	var err error
	tmpl, err = template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatalf("Error loading: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
			log.Println("Template execute error:", err)
		}
	})

	log.Println("Сервер запущен на : http://localhost:8080/")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
