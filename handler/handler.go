package handler

import (
	"net/http"
	"url-shorter/models"
	"url-shorter/storage"
	"url-shorter/utils"

	"github.com/gin-gonic/gin"
)

// ShortenURL handles the POST /shorten request
func ShortenURL(store *storage.URLShortener) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req models.ShortenRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
			return
		}

		// Normalize the URL
		normalizedURL, err := utils.NormalizeURL(req.URL)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}

		// Get or create the shortened URL
		shortURL := store.GetOrCreateShortURL(normalizedURL)

		c.JSON(http.StatusOK, gin.H{"short_url": shortURL})
	}
}

// RedirectURL handles the GET /{short_url} request
func RedirectURL(store *storage.URLShortener) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortURL := c.Param("short_url")

		// Find original URL
		originalURL, exists := store.GetOriginalURL(shortURL)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}

		c.Redirect(http.StatusMovedPermanently, originalURL)
	}
}
