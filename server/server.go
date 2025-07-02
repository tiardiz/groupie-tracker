package server

import (
	"errors"
	"log"
	"net/http"
	"os"

	"groupietracker/handler"
)

func RouteHandler(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		handler.MainHandler(w, r)
	case "/test500":
		handler.InternalServerErrorHandler(w, r, errors.New(r.Method+" testing error 500"))
	default:
		handler.NotFoundHandler(w, r, errors.New("page not found"))
	}
}

func Start() {
	err := handler.InitTemplates()
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	http.HandleFunc("/info/", handler.InfoHandler)

	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path[len("/static/"):]
		_, err := os.Stat(r.URL.Path[1:])
		if path == "" || errors.Is(err, os.ErrNotExist) {
			handler.NotFoundHandler(w, r, err)
			return
		}
		http.StripPrefix("/static/", http.FileServer(http.Dir("static"))).ServeHTTP(w, r)
	})

	http.HandleFunc("/", RouteHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
