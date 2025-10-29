package main

import (
	"log"
	"os"

	"github.com/0xatanda/country-currency/internal/database"
	"github.com/0xatanda/country-currency/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️  No .env file found, using system env")
	}

	db, err := database.ConnectProgres()
	if err != nil {
		log.Fatalf("❌  Failed to connect to database: %v", err)
	}
	log.Println("✅  Connected to database")

	r := gin.Default()

	handlers.RegisterCountryRoutes(r, db)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀  Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("√❌  Failed to run server: %v", err)
	}
}
