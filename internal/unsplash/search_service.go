package unsplash

import "encoding/json"

type SearchService struct {
	*service
}

type SearchOptions struct {
	Page    int    `url:"page"`
	PerPage int    `url:"per_page"`
	Query   string `url:"query"`
}

func (opt *SearchOptions) Valid() bool {
	if opt.Query == "" {
		return false
	}

	if opt.Page == 0 {
		opt.Page = 1
	}

	if opt.PerPage == 0 {
		opt.PerPage = 10
	} else if opt.PerPage > 30 {
		opt.PerPage = 30
	}

	return true
}

func (ss *SearchService) Photos(opt *SearchOptions) (*SearchPhotoResult, *Response, error) {
	if opt == nil {
		return nil, nil, &IllegalArgumentError{"opt argument can't be null"}
	}

	if !opt.Valid() {
		return nil, nil, &IllegalArgumentError{"search query can't be empty"}
	}

	req, err := newRequest("GET", getEndpoint(searchPhotos), opt, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := ss.do(req)
	if err != nil {
		return nil, nil, err
	}

	var photos SearchPhotoResult
	err = json.Unmarshal(*resp.Body, &photos)
	if err != nil {
		return nil, nil, err
	}

	return &photos, resp, nil
}
