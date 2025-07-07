package server

import (
	"fmt"
	"groupie-tracker/frontend/static"
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/helpers"
	"log"
	"net/http"
	"os"
)

func Init(link, port string) {
	if err := helpers.ChangeDirProjectRoot(); err != nil {
		log.Fatal(err)
	}
	if err := checkStaticFiles(); err != nil {
		log.Fatal(err)
	}
	fs := http.FileServer(http.FS(static.Files))
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artist/", handlers.HandleArtist)
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	log.Println("Starting server at", link+port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

const staticFilesLocation = "frontend/static/"

var staticFilenames = [...]string{
	"background.webp",
	"error.gif",
	"header-bg.jpg",
	"internalerror.gif",
	"style.css",
}

func checkStaticFiles() error {
	for _, filename := range staticFilenames {
		if _, err := os.Stat(staticFilesLocation + filename); err != nil {
			return fmt.Errorf("cannot access \"%s\": %w", filename, err)
		}
	}
	return nil
}
