package config

import (
	"fmt"
	"os"
	"strconv" 
)

type Config struct {
	DatabaseURL    string
	JWTSecret      string
	Port           string
	GINMode        string 
	DBMaxOpenConns int  
	DBMaxIdleConns int  
}

func Load() *Config {
	// Construct DATABASE_URL from individual components
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnv("DB_PORT", "5432") // Added DB_PORT if not explicitly in .env
	dbName := getEnv("DB_NAME", "pubtrans-eticketing") // Default to transportation_db or your preferred default
	dbUser := getEnv("DB_USERNAME", "postgres")
	dbPass := getEnv("DB_PASS", "password")
	sslMode := getEnv("DB_SSLMODE", "disable") // Allow configuring SSL mode

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslMode)

	return &Config{
		DatabaseURL:    databaseURL, // Now constructed from individual DB_* vars
		JWTSecret:      getEnv("JWT_SECRET", "your-256-bit-secret-that-you-must-change"), // CHANGE THIS DEFAULT
		Port:           getEnv("PORT", "8080"),
		GINMode:        getEnv("GIN_MODE", "debug"), // Read GIN_MODE
		DBMaxOpenConns: getEnvAsInt("DB_MAX_OPEN_CONNS", 25), 
		DBMaxIdleConns: getEnvAsInt("DB_MAX_IDLE_CONNS", 25), 
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	strValue := getEnv(key, strconv.Itoa(defaultValue))
	intValue, err := strconv.Atoi(strValue)
	if err != nil {
		return defaultValue // Fallback to default if parsing fails
	}
	return intValue
}