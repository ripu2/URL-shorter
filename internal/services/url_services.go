package services

import (
	"errors"
	"time"

	"example.com/url-shorter/internal/models"
)

func SortUrl(url *models.URL) (string, error) {
	url.CreatedAt = time.Now()
	val, err := url.GenerateURL()
	if err != nil {
		return "", errors.New(err.Error())
	}
	return val, nil
}

func GetRedirectURL(url string) (string, error) {
	longURL, err := models.GetLongURL(url)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return longURL, nil
}
