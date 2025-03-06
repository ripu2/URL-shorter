package utils

import (
	"crypto/rand"
	"math/big"
	"net/url"
)

func GenerateHash() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 6)
	for i := range result {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		result[i] = charset[num.Int64()]
	}
	return string(result)
}

func IsValidURL(testURL string) bool {
	parsedURL, err := url.ParseRequestURI(testURL)
	return err == nil && parsedURL.Scheme != "" && parsedURL.Host != ""
}
