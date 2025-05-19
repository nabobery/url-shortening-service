package database

import (
	"context"
	"database/sql"
	"errors"
	"time"
	"url-shortener/internal/core/model"
	"url-shortener/internal/core/repository"
)

var (
	ErrNotFound = errors.New("URL not found")
)

type postgresRepository struct {
	db *sql.DB
}

// NewPostgresRepository creates a new PostgreSQL repository
func NewPostgresRepository(db *sql.DB) repository.ShortURLRepository {
	return &postgresRepository{db: db}
}

func (r *postgresRepository) Create(ctx context.Context, shortURL *model.ShortURL) error {
	query := `
		INSERT INTO short_urls (id, url, short_code, created_at, updated_at, access_count)
		VALUES ($1, $2, $3, $4, $5, $6)
	`
	_, err := r.db.ExecContext(ctx, query,
		shortURL.ID,
		shortURL.URL,
		shortURL.ShortCode,
		shortURL.CreatedAt,
		shortURL.UpdatedAt,
		shortURL.AccessCount,
	)
	return err
}

func (r *postgresRepository) FindByShortCode(ctx context.Context, shortCode string) (*model.ShortURL, error) {
	query := `
		SELECT id, url, short_code, created_at, updated_at, access_count
		FROM short_urls
		WHERE short_code = $1
	`
	shortURL := &model.ShortURL{}
	err := r.db.QueryRowContext(ctx, query, shortCode).Scan(
		&shortURL.ID,
		&shortURL.URL,
		&shortURL.ShortCode,
		&shortURL.CreatedAt,
		&shortURL.UpdatedAt,
		&shortURL.AccessCount,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (r *postgresRepository) Update(ctx context.Context, shortCode string, newURL string) (*model.ShortURL, error) {
	query := `
		UPDATE short_urls
		SET url = $1, updated_at = $2
		WHERE short_code = $3
		RETURNING id, url, short_code, created_at, updated_at, access_count
	`
	shortURL := &model.ShortURL{}
	err := r.db.QueryRowContext(ctx, query, newURL, time.Now(), shortCode).Scan(
		&shortURL.ID,
		&shortURL.URL,
		&shortURL.ShortCode,
		&shortURL.CreatedAt,
		&shortURL.UpdatedAt,
		&shortURL.AccessCount,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return shortURL, nil
}

func (r *postgresRepository) Delete(ctx context.Context, shortCode string) error {
	query := "DELETE FROM short_urls WHERE short_code = $1"
	result, err := r.db.ExecContext(ctx, query, shortCode)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *postgresRepository) IncrementAccessCount(ctx context.Context, shortCode string) error {
	query := `
		UPDATE short_urls
		SET access_count = access_count + 1
		WHERE short_code = $1
	`
	result, err := r.db.ExecContext(ctx, query, shortCode)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *postgresRepository) GetStats(ctx context.Context, shortCode string) (*model.URLStatsResponse, error) {
	query := `
		SELECT url, short_code, access_count, created_at
		FROM short_urls
		WHERE short_code = $1
	`
	stats := &model.URLStatsResponse{}
	err := r.db.QueryRowContext(ctx, query, shortCode).Scan(
		&stats.URL,
		&stats.ShortCode,
		&stats.AccessCount,
		&stats.CreatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, ErrNotFound
	}
	if err != nil {
		return nil, err
	}
	return stats, nil
}
