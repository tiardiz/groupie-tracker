package main

import (
	"errors"
	"groupietracker/handler"
	"log"
	"net/http"
	"os"
)

func main() {
	err := handler.InitTemplates()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	// tmpl, err = template.ParseFiles("templates/index.html")
	// if err != nil {
	// 	log.Fatalf("Error loading: %v", err)
	// }
	// tmplError, err = template.ParseFiles("templates/error.html")
	// if err != nil {
	// 	log.Fatalf("Error loading error template: %v", err)
	// }

	http.HandleFunc("/", handler.MainHandler)
	http.HandleFunc("/info/", handler.InfoHandler)

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/static/"):]
		_, err := os.Stat(r.URL.Path[1:])
		if path == "" || errors.Is(err, os.ErrNotExist) {
			handler.NotFoundHandler(w, r)
			return
		}
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	http.HandleFunc("/test500", func(w http.ResponseWriter, r *http.Request) {
		handler.InternalServerErrorHandler(w, r)
	})

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
