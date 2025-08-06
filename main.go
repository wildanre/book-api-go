// @title Books API
// @version 1.0
// @description This is a simple CRUD API for managing books
// @contact.name API Support
// @contact.email support@example.com
// @host book-api-go.zeabur.app
// @BasePath /api
// @schemes https http
package main

import (
	"log"
	"os"
	"runtime"

	_ "example/go1/docs" // Import for swagger docs
	"example/go1/internal/database"
	"example/go1/internal/middleware"
	"example/go1/internal/models"
	"example/go1/internal/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Load environment variables dari file .env (optional untuk development)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate tables
	db := database.GetDB()
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Set Gin mode based on environment
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	// Initialize Gin router
	r := gin.New()

	// Add middleware
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// Health check endpoint (root level for basic monitoring)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":     "ok",
			"message":    "Books API is running",
			"version":    "1.0.0",
			"go_version": runtime.Version(),
		})
	})

	// API group
	api := r.Group("/api")
	{
		// Health check endpoint under API
		api.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":     "ok",
				"message":    "Books API is running",
				"version":    "1.0.0",
				"go_version": runtime.Version(),
			})
		})
	}

	// Swagger documentation
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Setup routes
	routes.SetupBookRoutes(r)

	// Get port from environment or default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server starting on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
