package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syafdia/demo-oauth2/internal/auth"
)

type withHeader struct {
	http.Header
	rt http.RoundTripper
}

func WithHeader(rt http.RoundTripper) withHeader {
	if rt == nil {
		rt = http.DefaultTransport
	}

	return withHeader{Header: make(http.Header), rt: rt}
}

func (h withHeader) RoundTrip(req *http.Request) (*http.Response, error) {
	for k, v := range h.Header {
		req.Header[k] = v
	}

	return h.rt.RoundTrip(req)
}

func setupHTMLTemplate(r *gin.Engine) {
	r.SetHTMLTemplate(template.Must(template.New("session.new").Parse(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<meta name="viewport" content="width=device-width, initial-scale=1.0">
			<title>Sign In</title>
		</head>
		<body>
			<h1>Please Sign In</h1>
			<br/>
			<form method="POST" action="/auth/session?login_challenge={{ .LoginChallenge  }}">
				<div style="margin:8px 16px">
					<label> Username </label>
					<input type="text" name="username" />
					<br/>
				</div>
				<div style="margin:8px 16px">
					<label> Password </label>
					<input type="password" name="password" />
					<br/>
				</div>

				<input type="submit" value="submit" style="margin:8px 16px">
			</form>
		</body>
		</html>
	`)))
}

func main() {
	rp := auth.NewRepo()
	cl := auth.NewClient()

	h := auth.NewAuthHandler(rp, cl)

	r := gin.Default()
	setupHTMLTemplate(r)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("auth/session/new", h.NewSession)
	r.POST("auth/session", h.CreateSession)
	r.GET("auth/consent/new", h.NewConsent)
	r.POST("auth/consent", h.CreateConsent)

	r.Run(":9001")
}
