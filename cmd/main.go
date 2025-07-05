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
	fs := http.FileServer(http.Dir("frontend/static"))
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artist/", handlers.HandleArtist)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Starting server at", link+port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

const port = ":8080"
const link = "https://localhost"
