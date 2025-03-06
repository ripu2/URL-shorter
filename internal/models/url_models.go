package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"example.com/url-shorter/db"
	"example.com/url-shorter/internal/utils"
)

type URL struct {
	ID        int64  `gorm:"primaryKey"`
	LongURL   string `gorm:"index"`       // Index hai, unique nahi
	ShortURL  string `gorm:"uniqueIndex"` // ShortURL unique hona chahiye
	CreatedAt time.Time
}

func isStringInDB(value string) bool {
	var exists bool
	query := "SELECT EXISTS(SELECT 1 FROM your_table WHERE short_url = $1)"
	err := db.DB.QueryRow(query, value).Scan(&exists)
	if err != nil {
		errors.New("DB query error:")
	}
	return exists
}

func getUniqueString() string {
	for {
		randomStr := utils.GenerateHash()
		if !isStringInDB(randomStr) {
			return randomStr
		}
	}
}

func (u *URL) GenerateURL() (string, error) {
	query := `
		INSERT INTO urls (long_url, short_url, created_at)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	if !utils.IsValidURL(u.LongURL) {
		return "", errors.New("invalid URL")
	}

	hash := getUniqueString()

	fmt.Println("Hash: ", hash)
	err := db.DB.QueryRow(query, u.LongURL, hash, u.CreatedAt).Scan(&u.ID)
	if err != nil {
		return "", errors.New(err.Error())
	}

	return hash, nil
}

func GetLongURL(shortURL string) (string, error) {
	var longURL string
	query := "SELECT long_url FROM urls WHERE short_url = $1"
	err := db.DB.QueryRow(query, shortURL).Scan(&longURL)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("link: %s is either expired or invalid", shortURL)
		}
		return "", fmt.Errorf("DB query error: %w", err)
	}

	return longURL, nil
}
