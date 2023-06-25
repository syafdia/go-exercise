package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Module struct {
	redisClient *redis.Client
}

func PingHandler(m *Module) gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := m.redisClient.Ping(c).Result()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": fmt.Sprintf("Redis is not OK, err: %v", err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    "pong",
		})
	}
}

func AddHandler(m *Module) gin.HandlerFunc {
	return func(c *gin.Context) {
		urlQuery := c.Request.URL.Query()
		rawA := urlQuery.Get("a")
		rawB := urlQuery.Get("b")

		maxNum, _ := strconv.Atoi(envMaxNumber)
		a, _ := strconv.Atoi(rawA)
		b, _ := strconv.Atoi(rawB)

		if a > maxNum || b > maxNum {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": fmt.Sprintf("max allowed a or b value is %d", maxNum),
			})
			return
		}

		cacheKey := fmt.Sprintf("add:%d+%d", a, b)
		var result int

		_, err := m.redisClient.Get(c, cacheKey).Result()
		if err != nil {
			if !errors.Is(err, redis.Nil) {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Redis is not OK when GET, err: %v", err),
				})
				return
			}

			result = doHeavyCalculation(a, b)

			_, err := m.redisClient.Set(c, cacheKey, result, 30*time.Second).Result()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"message": fmt.Sprintf("Redis is not OK when SET, err: %v", err),
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "OK",
			"data":    result,
		})
	}
}

func doHeavyCalculation(a int, b int) int {
	time.Sleep(3 * time.Second)
	return a + b
}

var (
	envHttpServerPort = os.Getenv("HTTP_SERVER_PORT")
	envMaxNumber      = os.Getenv("MAX_NUMBER")
	envRedisURL       = os.Getenv("REDIS_URL")
)

func main() {
	httpServerPort := "8080"
	if envHttpServerPort != "" {
		httpServerPort = envHttpServerPort
	}

	redisOpt, err := redis.ParseURL(envRedisURL)
	if err != nil {
		log.Panicf("[main] Failed parsing Redis URL, err: %s", err)
	}

	redisClient := redis.NewClient(redisOpt)

	module := &Module{
		redisClient: redisClient,
	}

	r := gin.Default()

	r.GET("/ping", PingHandler(module))
	r.GET("/add", AddHandler(module))

	r.Run(fmt.Sprintf(":%s", httpServerPort))
}
