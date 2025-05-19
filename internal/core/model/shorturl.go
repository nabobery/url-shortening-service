package model

import "time"

// ShortURL represents a shortened URL entry in the system
type ShortURL struct {
	ID          string    `json:"id" db:"id"`
	URL         string    `json:"url" db:"url"`
	ShortCode   string    `json:"shortCode" db:"short_code"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
	AccessCount int64     `json:"accessCount,omitempty" db:"access_count"`
}

// CreateShortURLRequest represents the request body for creating a short URL
type CreateShortURLRequest struct {
	URL string `json:"url" binding:"required,url"`
}

// CreateShortURLResponse represents the response body for creating a short URL
type CreateShortURLResponse struct {
	ShortCode string `json:"shortCode"`
	URL       string `json:"url"`
}

// URLStatsResponse represents the response body for URL statistics
type URLStatsResponse struct {
	ShortCode   string    `json:"shortCode"`
	URL         string    `json:"url"`
	AccessCount int64     `json:"accessCount"`
	CreatedAt   time.Time `json:"createdAt"`
}
