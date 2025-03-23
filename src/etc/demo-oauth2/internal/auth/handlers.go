package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	repo   *Repo
	client *Client
}

func NewAuthHandler(repo *Repo, client *Client) *AuthHandler {
	return &AuthHandler{
		repo:   repo,
		client: client,
	}
}

func (a *AuthHandler) NewSession(c *gin.Context) {
	c.HTML(http.StatusOK, "session.new", map[string]interface{}{
		"LoginChallenge": c.Query("login_challenge"),
	})
}

func (a *AuthHandler) CreateSession(c *gin.Context) {
	username := c.Request.FormValue("username")
	password := c.Request.FormValue("password")
	// loginChallenge := c.Request.FormValue("login_challenge")

	user, err := a.repo.User.FindByUsernameAndPassword(c, username, password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		// TODO: render token
		"username": user.Username,
	})
}

func (a *AuthHandler) NewConsent(c *gin.Context) {

}

func (a *AuthHandler) CreateConsent(c *gin.Context) {

}
