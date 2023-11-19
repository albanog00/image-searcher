package unsplash

type UnsplashClient struct {
	baseUrl string
}

func New() *UnsplashClient {
	return &UnsplashClient{
		baseUrl: "https://api.unsplash.com",
	}
}
