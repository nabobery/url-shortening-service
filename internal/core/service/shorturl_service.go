package service

import (
	"context"
	"errors"
	"log"
	"time"
	"url-shortener/internal/core/model"
	"url-shortener/internal/core/repository"
	"url-shortener/internal/platform/shortener"

	"github.com/google/uuid"
)

var (
	ErrURLNotFound = errors.New("URL not found")
	ErrInvalidURL  = errors.New("invalid URL")
)

// ShortURLService defines the interface for URL shortening business operations
type ShortURLService interface {
	CreateShortURL(ctx context.Context, url string) (*model.CreateShortURLResponse, error)
	GetOriginalURL(ctx context.Context, shortCode string) (string, error)
	GetURLStats(ctx context.Context, shortCode string) (*model.URLStatsResponse, error)
	UpdateURL(ctx context.Context, shortCode, newURL string) (*model.ShortURL, error)
	DeleteURL(ctx context.Context, shortCode string) error
}

type shortURLService struct {
	repo      repository.ShortURLRepository
	generator *shortener.Generator
}

// NewShortURLService creates a new instance of ShortURLService
func NewShortURLService(repo repository.ShortURLRepository, generator *shortener.Generator) ShortURLService {
	return &shortURLService{
		repo:      repo,
		generator: generator,
	}
}

func (s *shortURLService) CreateShortURL(ctx context.Context, url string) (*model.CreateShortURLResponse, error) {
	if url == "" {
		return nil, ErrInvalidURL
	}

	shortCode := s.generator.Generate()
	now := time.Now()

	shortURL := &model.ShortURL{
		ID:          uuid.New().String(),
		URL:         url,
		ShortCode:   shortCode,
		CreatedAt:   now,
		UpdatedAt:   now,
		AccessCount: 0,
	}

	if err := s.repo.Create(ctx, shortURL); err != nil {
		return nil, err
	}

	return &model.CreateShortURLResponse{
		ShortCode: shortCode,
		URL:       url,
	}, nil
}

func (s *shortURLService) GetOriginalURL(ctx context.Context, shortCode string) (string, error) {
	shortURL, err := s.repo.FindByShortCode(ctx, shortCode)
	if err != nil {
		return "", ErrURLNotFound
	}

	if err := s.repo.IncrementAccessCount(ctx, shortCode); err != nil {
		// Log the error but don't fail the request
		log.Printf("Failed to increment access count: %v", err)
	}

	return shortURL.URL, nil
}

func (s *shortURLService) GetURLStats(ctx context.Context, shortCode string) (*model.URLStatsResponse, error) {
	return s.repo.GetStats(ctx, shortCode)
}

func (s *shortURLService) UpdateURL(ctx context.Context, shortCode, newURL string) (*model.ShortURL, error) {
	if newURL == "" {
		return nil, ErrInvalidURL
	}
	return s.repo.Update(ctx, shortCode, newURL)
}

func (s *shortURLService) DeleteURL(ctx context.Context, shortCode string) error {
	return s.repo.Delete(ctx, shortCode)
}
