package storage

import (
	"sync"
	"url-shorter/utils"
)

// URLShortener stores the mapping of short URLs to original URLs
type URLShortener struct {
	store map[string]string
	mu    sync.RWMutex
}

// NewURLShortener initializes a new URLShortener instance
func NewURLShortener() *URLShortener {
	return &URLShortener{
		store: make(map[string]string),
	}
}

// GetOrCreateShortURL checks if the URL exists and returns it, otherwise creates a new short URL
func (u *URLShortener) GetOrCreateShortURL(originalURL string) string {
	u.mu.RLock()
	for shortURL, storedURL := range u.store {
		if storedURL == originalURL {
			u.mu.RUnlock()
			return shortURL
		}
	}
	u.mu.RUnlock()

	// If not found, generate a new short URL
	shortURL := utils.GenerateShortURL()

	// Store the new mapping
	u.mu.Lock()
	u.store[shortURL] = originalURL
	u.mu.Unlock()

	return shortURL
}

// GetOriginalURL retrieves the original URL for a given short URL
func (u *URLShortener) GetOriginalURL(shortURL string) (string, bool) {
	u.mu.RLock()
	originalURL, exists := u.store[shortURL]
	u.mu.RUnlock()

	return originalURL, exists
}
