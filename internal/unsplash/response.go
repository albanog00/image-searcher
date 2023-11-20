package unsplash

import (
	"io"
	"net/http"
	"strconv"
)

type Response struct {
	httpResponse       *http.Response
	Body               *[]byte
	RateLimit          int
	RateLimitRemaining int
}

func newResponse(r *http.Response) (*Response, error) {
	if r == nil {
		return nil,
			&IllegalArgumentError{"http.Response can't be null"}
	}

	resp := &Response{httpResponse: r}
	resp.GetRateLimitValues()

	buff, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	resp.Body = &buff

	return resp, nil
}

func (r *Response) GetRateLimitValues() {
	maxLimit, ok := r.httpResponse.Header["X-Ratelimit-Limit"]
	if ok && len(maxLimit) == 1 {
		r.RateLimit, _ = strconv.Atoi(maxLimit[0])
	}

	remaining, ok := r.httpResponse.Header["X-Ratelimit-Remaining"]
	if ok && len(remaining) == 1 {
		r.RateLimitRemaining, _ = strconv.Atoi(remaining[0])
	}
}
