package web

import (
	"net/http"
	"url-shortener/internal/platform/web/handler"
	"url-shortener/internal/platform/web/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter configures the Gin router with all routes and middleware
func SetupRouter(urlHandler *handler.ShortURLHandler) *gin.Engine {
	router := gin.New()

	// Use the custom logger middleware
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	// Load HTML templates
	router.LoadHTMLGlob("web/templates/*")

	// Serve static files
	router.Static("/static", "web/static")

	// Serve index page
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	// API routes
	api := router.Group("/api")
	{
		// URL shortening endpoints
		api.POST("/shorten", urlHandler.CreateShortURL)
		api.GET("/urls/:shortCode", urlHandler.GetURLStats)
		api.PUT("/urls/:shortCode", urlHandler.UpdateURL)
		api.DELETE("/urls/:shortCode", urlHandler.DeleteURL)
	}

	// Redirect endpoint (not under /api as it's user-facing)
	router.GET("/:shortCode", urlHandler.RedirectToOriginalURL)

	return router
}
