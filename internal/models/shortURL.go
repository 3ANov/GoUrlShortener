package models

import "time"

type ShortURL struct {
	Id          int
	OriginalUrl string
	ShortCode   string
	CreatedAt   time.Time
	ExpiresAt   time.Time
	UsageCount  int
}

type ShortURLRequest struct {
	OriginalUrl string    `json:"original_url" binding:"required"`
	ExpiresAt   time.Time `json:"expires_at" binding:"required"`
}

type ShortURLResponse struct {
	ShortCode string    `json:"short_code"`
	ExpiresAt time.Time `json:"expires_at"`
}
