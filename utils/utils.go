package utils

import (
	"math/rand"
	"net/url"
	"strings"
	"time"
)

// NormalizeURL ensures URLs have the correct scheme and removes www. prefix
func NormalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "http"
	}

	// Remove "www." if present
	parsedURL.Host = strings.TrimPrefix(parsedURL.Host, "www.")

	return parsedURL.String(), nil
}

// GenerateShortURL creates a random short URL
func GenerateShortURL() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 6
	rand.Seed(time.Now().UnixNano())
	shortURL := make([]byte, length)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}
	return string(shortURL)
}
