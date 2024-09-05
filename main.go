package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type URLShortener struct {
	store map[string]string
	mu    sync.RWMutex
}

func NewURLShortener() *URLShortener {
	return &URLShortener{
		store: make(map[string]string),
	}
}

func generateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	rand.Seed(time.Now().UnixNano())
	shortURL := make([]byte, length)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortURL)
}

func (u *URLShortener) shorten(c *gin.Context) {
	var req struct {
		URL string `json:"url" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	_, err := url.ParseRequestURI(req.URL)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	shortURL := generateShortURL()

	u.mu.Lock()
	u.store[shortURL] = req.URL
	u.mu.Unlock()

	shortenedURL := fmt.Sprintf("http://localhost:8080/%s", shortURL)
	c.JSON(http.StatusOK, gin.H{"short_url": shortenedURL})
}

func main() {
	r := gin.Default()

	urlShortener := NewURLShortener()

	r.POST("/shorten", urlShortener.shorten)

	r.Run(":8080")
}
