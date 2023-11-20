package unsplash

import (
	"net/http"
	"os"
)

type Unsplash struct {
	client_id string
	client    *http.Client
	Photo     *PhotoService
	Search    *SearchService
	// User      *UserService
	// Collection *CollectionService
}

func NewUnsplash(client *http.Client) *Unsplash {
	if client == nil {
		client = http.DefaultClient
	}

	service := newService(client)

	unsplash := new(Unsplash)
	unsplash.client_id = os.Getenv("UNSPLASH_ACCESS_KEY")
	unsplash.client = client
	unsplash.Photo = &PhotoService{service}
	unsplash.Search = &SearchService{service}

	return unsplash
}
