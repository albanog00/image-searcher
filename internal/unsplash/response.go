package unsplash

import "time"

type ProfileImage struct {
	Small  string `json:"small"`
	Medium string `json:"medium"`
	Large  string `json:"large"`
}

type UserLinks struct {
	Self   string `json:"self"`
	Html   string `json:"html"`
	Photos string `json:"photos"`
	Likes  string `json:"likes"`
}

type User struct {
	Id                string       `json:"id"`
	Username          string       `json:"username"`
	Name              string       `json:"name"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	InstagramUsername string       `json:"instagram_username"`
	TwitterUsername   string       `json:"twitter_username"`
	PortfolioUrl      string       `json:"portfolio_url"`
	ProfileImage      ProfileImage `json:"profile_image"`
	Links             UserLinks    `json:"links"`
}

type Collection struct {
}

type Urls struct {
	Raw     string `json:"raw"`
	Full    string `json:"full"`
	Regular string `json:"regular"`
	Small   string `json:"small"`
	Thumb   string `json:"thumb"`
}

type PhotosLinks struct {
	Self     string `json:"self"`
	Html     string `json:"html"`
	Download string `json:"download"`
}

type PhotosSearchResult struct {
	Id                     string       `json:"id"`
	CreatedAt              time.Time    `json:"created_at"`
	UpdatedAt              time.Time    `json:"updated_at"`
	Width                  int          `json:"width"`
	Height                 int          `json:"height"`
	Color                  string       `json:"color"`
	BlurHash               string       `json:"blur_hash"`
	Likes                  int          `json:"likes"`
	LikedByUser            bool         `json:"liked_by_user"`
	Description            string       `json:"description"`
	User                   User         `json:"user"`
	CurrentUserCollections []Collection `json:"current_user_collections"`
	Urls                   Urls         `json:"urls"`
	Links                  PhotosLinks  `json:"links"`
}

type PhotosSearch struct {
	Total      int                  `json:"total"`
	TotalPages int                  `json:"total_pages"`
	Results    []PhotosSearchResult `json:"results"`
}
