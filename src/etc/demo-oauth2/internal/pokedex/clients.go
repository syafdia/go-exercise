package pokedex

import "golang.org/x/oauth2"

type Client struct {
	OAuth2  *oauth2.Config
	Pokemon interface{}
}

func NewClient(oauth2 *oauth2.Config) *Client {
	return &Client{
		OAuth2: oauth2,
	}
}
