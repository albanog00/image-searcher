package unsplash

type endpoint int

const (
	base endpoint = iota
	search
	photos
	searchPhotos
)

const (
	searchEndpoint       = "search"
	photosEndpoint       = "photos"
	searchPhotosEndpoint = searchEndpoint + "/" + photosEndpoint
)

const apiUrl = "https://api.unsplash.com/"

var mapUrl map[endpoint]string

func init() {
	mapUrl = make(map[endpoint]string)
	mapUrl[base] = apiUrl
	mapUrl[search] = searchEndpoint
	mapUrl[photos] = photosEndpoint
	mapUrl[searchPhotos] = searchPhotosEndpoint
}

func getEndpoint(e endpoint) string {
	return mapUrl[e]
}
