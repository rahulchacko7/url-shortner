package main

import (
	"url-shorter/handler"
	"url-shorter/storage"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize storage
	urlStorage := storage.NewURLShortener()

	// Set up routes
	r.POST("/shorten", handler.ShortenURL(urlStorage))
	r.GET("/:short_url", handler.RedirectURL(urlStorage))

	// Start server
	r.Run(":8080")
}
