package server

import (
	"fmt"
	"groupie-tracker/frontend/static"
	"groupie-tracker/internal/handlers"
	"groupie-tracker/internal/helpers"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func Init(link, port string) {
	setupErr := prepareEnv()
	if setupErr != nil {
		log.Fatal(setupErr)
	}
	fs := http.FileServer(http.FS(static.Files))
	http.HandleFunc("/", handlers.HandleHome)
	http.HandleFunc("/artist/", handlers.HandleArtist)
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	addr := link + port
	log.Println("Starting server at", addr)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func prepareEnv() error {
	if err := helpers.ChangeDirProjectRoot(); err != nil {
		return err
	}
	if err := checkStaticFiles(); err != nil {
		return err
	}
	return nil
}

var staticFilesLocation = filepath.Join("frontend", "static")

var staticFilenames = [...]string{
	"background.webp",
	"error.gif",
	"header-bg.jpg",
	"internalerror.gif",
	"style.css",
}

func checkStaticFiles() error {
	for _, filename := range staticFilenames {
		if _, err := os.Stat(filepath.Join(staticFilesLocation, filename)); err != nil {
			return fmt.Errorf("cannot access \"%s\": %w", filename, err)
		}
	}
	return nil
}
