package unsplash

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (unsplashClient *UnsplashClient) Request(method string, endpoint string) (*http.Response, error) {
	req, err := http.NewRequest(method, unsplashClient.baseUrl+endpoint, nil)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", os.Getenv("UNSPLASH_ACCESS_KEY")))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf(err.Error())
		return nil, err
	}
	return res, nil
}
