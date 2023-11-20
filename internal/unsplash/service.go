package unsplash

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type service struct {
	client *http.Client
}

func newService(client *http.Client) *service {
	return &service{
		client: client,
	}
}

func (s *service) do(req *request) (*Response, error) {
	if req == nil {
		return nil, &IllegalArgumentError{"req argument can't be null"}
	}

	req.Request.Header.Set("Accept-Version", "1")
	req.Request.Header.Set("Authorization", fmt.Sprintf("Client-ID %s", os.Getenv("UNSPLASH_ACCESS_KEY")))

	client := s.client
	rawResp, err := client.Do(req.Request)
	if err != nil {
		return nil, err
	}
	if rawResp != nil {
		defer rawResp.Body.Close()
	}

	resp, err := newResponse(rawResp)
	if err != nil {
		return nil, err
	}

	_, err = io.Copy(io.Discard, rawResp.Body)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
