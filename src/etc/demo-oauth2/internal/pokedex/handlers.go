package pokedex

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PokedexHandler struct {
	client *Client
}

func NewPokedexhandler(client *Client) *PokedexHandler {
	return &PokedexHandler{
		client: client,
	}
}

func (a *PokedexHandler) ListPokemons(c *gin.Context) {
	authCodeURL := a.client.OAuth2.AuthCodeURL(fmt.Sprintf("%d", time.Now().Unix()))
	c.HTML(http.StatusOK, "pokemons.list", map[string]interface{}{
		"AuthCodeURL": authCodeURL,
	})
}

func (a *PokedexHandler) OAuth2Callbacks(c *gin.Context) {
	// TODO
}
