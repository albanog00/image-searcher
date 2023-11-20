package unsplash

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
