package unsplash

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/google/go-querystring/query"
)

type request struct {
	Request *http.Request
}

func newRequest(method string, e string, qs interface{}, body interface{}) (*request, error) {
	if e == "" {
		return nil, &IllegalArgumentError{"endpoint argument cannot be empty"}
	}

	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	httpRequest, err := http.NewRequest(method, getEndpoint(base)+e, bytes.NewBuffer(buf))
	if err != nil {
		return nil, err
	}

	if qs != nil {
		values, err := query.Values(qs)
		if err != nil {
			return nil, err
		}
		httpRequest.URL.RawQuery = values.Encode()
	}

	req := new(request)
	req.Request = httpRequest
	req.Request.Header.Add("Content-Type", "application/json")

	return req, nil
}
