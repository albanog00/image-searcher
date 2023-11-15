package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func SearchHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("/search handler called")
	// Retrieves query from url string

	q := r.URL.Query().Get("q")
	if q == "" {
		http.Error(w, "query parameter is required", 400)
		return
	}

	// TODO: Search images

	var urls []string
	num := rand.Intn(100)
	for i := 0; i < 12; i++ {
		// generate random numbers
		urls = append(urls, fmt.Sprintf("https://source.unsplash.com/random/?sig=%d", (i%5)*num))
	}

	// return urls as json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(urls)
	if err != nil {
		log.Println(err)
		http.Error(w, "failed to encode response", 500)
		return
	}
}
