package main

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func generateShortKey() string {
	// make a byte slice b of length 6
	b := make([]byte, 6)

	// generate random bytes and write them into b
	_, _ = rand.Read(b)

	// converts the random bytes into a URL-safe base64 string, return the first 6 characters
	return base64.URLEncoding.EncodeToString(b)[:6]
}

type ShortenRequest struct {
	LongURL string `json:"long_url"`
}

func Shorten(c *gin.Context) {
	var req ShortenRequest

	if err := c.ShouldBindJSON(&req); err != nil || req.LongURL == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	shortKey := generateShortKey()

	err := rdb.Set(ctx, shortKey, req.LongURL, 0).Err()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store URL"})
		return
	}

	shortURL := fmt.Sprintf("http://localhost:8080/%s", shortKey)
	c.JSON(http.StatusOK, gin.H{
		"short_url": shortURL})
}

func GetOriginURL(c *gin.Context) {
	shortKey := c.Param("shortKey")

	longURL, err := rdb.Get(ctx, shortKey).Result()

	if err == redis.Nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
		return
	}

	c.Redirect(http.StatusFound, longURL)
}

func main() {
	r := gin.Default()

	r.POST("/shorten", Shorten)
	r.GET("/:shortKey", GetOriginURL)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
