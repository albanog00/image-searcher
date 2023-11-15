package routes

import (
	"net/http"

	"giuseppealbano.dev/img-searcher/internal/handlers"
)

func Setup() {
	http.HandleFunc("/search", handlers.SearchHandler)
}
