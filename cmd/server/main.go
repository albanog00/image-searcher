package main

import (
	"log"
	"net/http"

	"giuseppealbano.dev/img-searcher/internal/routes"
)

func main() {
	routes.Setup()

	log.Println("Listening on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Listen: %v", err)
	}
}
