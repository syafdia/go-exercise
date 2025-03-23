package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syafdia/demo-oauth2/internal/pokedex"
	"golang.org/x/oauth2"
)

func newOauth2Client() *oauth2.Config {
	return &oauth2.Config{
		RedirectURL: "http://localhost:9002/callbacks", // Pokedex App callbacks URL

		// For testing purpose, we hardcoded the client_id & client_secret
		ClientID:     "client-id-my-pokedex",
		ClientSecret: "client-secret-my-pokedex",

		Scopes: []string{"offline", "pokemon:all"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "http://localhost:4444/oauth2/auth",
			TokenURL: "http://localhost:4444/oauth2/token",
		},
	}
}

func setupHTMLTemplate(r *gin.Engine) {
	r.SetHTMLTemplate(template.Must(template.New("pokemons.list").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>List Pokemon</title>
		</head>
		<body>
			<h1>Please Sign In First</h1>
			<br/>
			Go to this URL to Sign In:
			<br/>
			<a href="{{ .AuthCodeURL }}">{{ .AuthCodeURL }}</a>
			
		</body>
		</html>
	`)))
}

func main() {
	c := pokedex.NewClient(newOauth2Client())

	h := pokedex.NewPokedexhandler(c)

	r := gin.Default()
	setupHTMLTemplate(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/callbacks", h.OAuth2Callbacks)
	r.GET("/pokemons", h.ListPokemons)

	r.Run(":9002")
}
