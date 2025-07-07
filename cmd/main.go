package main

import (
	"groupie-tracker/internal/helpers"
	"groupie-tracker/internal/server"
	"log"
)

func main() {
	if err := helpers.ChangeDirProjectRoot(); err != nil {
		log.Fatal(err)
	}
	server.Init(link, port)
}

const port = ":8080"
const link = "http://localhost"
