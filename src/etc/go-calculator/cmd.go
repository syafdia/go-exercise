package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    "pong",
	})
}

func AddHandler(c *gin.Context) {
	urlQuery := c.Request.URL.Query()
	rawA := urlQuery.Get("a")
	rawB := urlQuery.Get("b")

	a, _ := strconv.Atoi(rawA)
	b, _ := strconv.Atoi(rawB)

	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
		"data":    a + b,
	})
}

func main() {
	httpServerPort := os.Getenv("HTTP_SERVER_PORT")
	if httpServerPort == "" {
		httpServerPort = "8080"
	}

	r := gin.Default()

	r.GET("/ping", PingHandler)
	r.GET("/add", AddHandler)

	r.Run(fmt.Sprintf(":%s", httpServerPort))
}
