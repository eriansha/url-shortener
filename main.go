package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Shorten(c *gin.Context) {
	// TODO
}

func GetOriginURL(c *gin.Context) {
	// TODO
}

func main() {
	r := gin.Default()

	r.POST("/shorten", Shorten)
	r.GET("/:shortKey", GetOriginURL)

	log.Println("Server running on :8080")
	r.Run(":8080")
}
