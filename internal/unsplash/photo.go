package unsplash

import "time"

type PhotoLinks struct {
	Self     string `json:"self"`
	Html     string `json:"html"`
	Download string `json:"download"`
}

type Photo struct {
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
	Links                  PhotoLinks   `json:"links"`
}
