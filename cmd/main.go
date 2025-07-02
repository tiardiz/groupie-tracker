package main

import (
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/helpers"
	"log"
	"net/http"
)

func main() {
	if err := helpers.ChangeDirProjectRoot(); err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("GET /", handlers.HandleHome)
	http.HandleFunc("GET /artist/", handlers.HandleArtist)
	log.Println("Starting server at", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

const port = ":8080"
