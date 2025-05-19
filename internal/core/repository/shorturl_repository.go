package repository

import (
	"context"
	"url-shortener/internal/core/model"
)

// ShortURLRepository defines the interface for URL shortening data operations
type ShortURLRepository interface {
	// Create stores a new shortened URL
	Create(ctx context.Context, shortURL *model.ShortURL) error

	// FindByShortCode retrieves a URL by its short code
	FindByShortCode(ctx context.Context, shortCode string) (*model.ShortURL, error)

	// Update modifies an existing shortened URL
	Update(ctx context.Context, shortCode string, newURL string) (*model.ShortURL, error)

	// Delete removes a shortened URL
	Delete(ctx context.Context, shortCode string) error

	// IncrementAccessCount increases the access count for a URL
	IncrementAccessCount(ctx context.Context, shortCode string) error

	// GetStats retrieves the statistics for a shortened URL
	GetStats(ctx context.Context, shortCode string) (*model.URLStatsResponse, error)
}
