package routes

import (
	"net/http"
)

func Setup() {
	http.HandleFunc("/", nil)
}
