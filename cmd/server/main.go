package main

import (
	"log"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	
	"pubtrans-eticketing/internal/config"
	"pubtrans-eticketing/internal/db"
	"pubtrans-eticketing/internal/handlers"
	"pubtrans-eticketing/internal/middleware"
	"pubtrans-eticketing/internal/repositories"
	"pubtrans-eticketing/internal/services"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, relying on system environment variables or defaults.")
	}

	cfg := config.Load()
	log.Printf("Loaded configuration: DatabaseURL=%s, JWTSecret=<hidden>, Port=%s, GINMode=%s, DBMaxOpenConns=%d, DBMaxIdleConns=%d",
		cfg.DatabaseURL, cfg.Port, cfg.GINMode, cfg.DBMaxOpenConns, cfg.DBMaxIdleConns)

	// Set Gin mode based on config
	gin.SetMode(cfg.GINMode)

	// Pass connection pool settings from config to NewConnection
	gormDB, err := db.NewConnection(cfg.DatabaseURL, cfg.DBMaxOpenConns, cfg.DBMaxIdleConns)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	
	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatalf("Failed to get underlying SQL DB instance: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Printf("Error closing database connection: %v", err)
		} else {
			log.Println("Database connection closed gracefully.")
		}
	}()

	log.Println("Database connection established and migrations run successfully.")

	userRepo := repositories.NewUserRepository(gormDB)
	terminalRepo := repositories.NewTerminalRepository(gormDB)
	log.Println("Repositories initialized.")

	authService := services.NewAuthService(userRepo, cfg.JWTSecret)
	terminalService := services.NewTerminalService(terminalRepo)
	log.Println("Services initialized.")

	authHandler := handlers.NewAuthHandler(authService)
	terminalHandler := handlers.NewTerminalHandler(terminalService)
	log.Println("Handlers initialized.")

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authHandler.Login)
		}

		protected := v1.Group("/")
		protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
		{
			protected.POST("/terminals", terminalHandler.CreateTerminal)
		}
	}

	log.Printf("Server starting on port %s...", cfg.Port)
	log.Fatal(router.Run(fmt.Sprintf(":%s", cfg.Port)))
}