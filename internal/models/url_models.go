package models

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	db "example.com/url-shorter/config/db"
	"example.com/url-shorter/internal/utils"
)

type URL struct {
	ID        int64  `gorm:"primaryKey"`
	LongURL   string `gorm:"index"`       // Index hai, unique nahi
	ShortURL  string `gorm:"uniqueIndex"` // ShortURL unique hona chahiye
	CreatedAt time.Time
}

func (u *URL) GenerateURL() (string, error) {
	query := `
		INSERT INTO urls (long_url, short_url, created_at)
		SELECT $1, $2::VARCHAR, $3
		WHERE NOT EXISTS (
			SELECT 1 FROM urls WHERE short_url = $2::VARCHAR
		)
		RETURNING id;
	`

	if !utils.IsValidURL(u.LongURL) {
		return "", errors.New("invalid URL")
	}

	for {
		hash := utils.GenerateHash()
		err := db.DB.QueryRow(query, u.LongURL, hash, u.CreatedAt).Scan(&u.ID)

		if err == nil {
			return hash, nil
		}

		if errors.Is(err, sql.ErrNoRows) {
			continue
		}

		return "", fmt.Errorf("DB insert error: %w", err)
	}
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
