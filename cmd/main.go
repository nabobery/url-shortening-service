package main

import (
	"fmt"
	"log"
	"url-shortener/internal/config"
	"url-shortener/internal/core/service"
	"url-shortener/internal/platform/database"
	"url-shortener/internal/platform/shortener"
	"url-shortener/internal/platform/web"
	"url-shortener/internal/platform/web/handler"
)

func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	// Connect to database
	db, err := database.NewPostgresDB(cfg.PostgresDSN())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create database tables
	if err := database.CreateTables(db); err != nil {
		log.Fatalf("Failed to create database tables: %v", err)
	}

	// Initialize components
	urlRepo := database.NewPostgresRepository(db)
	urlGenerator := shortener.NewGenerator(6) // 6-character short codes
	urlService := service.NewShortURLService(urlRepo, urlGenerator)
	urlHandler := handler.NewShortURLHandler(urlService)

	// Setup router
	router := web.SetupRouter(urlHandler)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	log.Printf("Server starting on %s", addr)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
