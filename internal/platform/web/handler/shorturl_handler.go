package handler

import (
	"net/http"
	"url-shortener/internal/core/model"
	"url-shortener/internal/core/service"

	"github.com/gin-gonic/gin"
)

// ShortURLHandler handles HTTP requests for URL shortening operations
type ShortURLHandler struct {
	service service.ShortURLService
}

// NewShortURLHandler creates a new ShortURLHandler
func NewShortURLHandler(service service.ShortURLService) *ShortURLHandler {
	return &ShortURLHandler{
		service: service,
	}
}

// CreateShortURL handles the creation of a new shortened URL
func (h *ShortURLHandler) CreateShortURL(c *gin.Context) {
	var req model.CreateShortURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	response, err := h.service.CreateShortURL(c.Request.Context(), req.URL)
	if err != nil {
		if err == service.ErrInvalidURL {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create short URL"})
		return
	}

	c.JSON(http.StatusCreated, response)
}

// RedirectToOriginalURL handles the redirection to the original URL
func (h *ShortURLHandler) RedirectToOriginalURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short code is required"})
		return
	}

	originalURL, err := h.service.GetOriginalURL(c.Request.Context(), shortCode)
	if err != nil {
		if err == service.ErrURLNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve URL"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, originalURL)
}

// GetURLStats handles retrieving statistics for a shortened URL
func (h *ShortURLHandler) GetURLStats(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short code is required"})
		return
	}

	stats, err := h.service.GetURLStats(c.Request.Context(), shortCode)
	if err != nil {
		if err == service.ErrURLNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve stats"})
		return
	}

	c.JSON(http.StatusOK, stats)
}

// UpdateURL handles updating an existing shortened URL
func (h *ShortURLHandler) UpdateURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short code is required"})
		return
	}

	var req model.CreateShortURLRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedURL, err := h.service.UpdateURL(c.Request.Context(), shortCode, req.URL)
	if err != nil {
		if err == service.ErrURLNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		if err == service.ErrInvalidURL {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update URL"})
		return
	}

	c.JSON(http.StatusOK, updatedURL)
}

// DeleteURL handles deleting a shortened URL
func (h *ShortURLHandler) DeleteURL(c *gin.Context) {
	shortCode := c.Param("shortCode")
	if shortCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Short code is required"})
		return
	}

	err := h.service.DeleteURL(c.Request.Context(), shortCode)
	if err != nil {
		if err == service.ErrURLNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete URL"})
		return
	}

	c.Status(http.StatusNoContent)
}
