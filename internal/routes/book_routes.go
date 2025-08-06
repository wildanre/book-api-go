package routes

import (
	"example/go1/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupBookRoutes sets up book-related routes
func SetupBookRoutes(router *gin.Engine) {
	bookHandler := handlers.NewBookHandler()

	// Group routes with /api/v1 prefix
	v1 := router.Group("/api/v1")
	{
		books := v1.Group("/books")
		{
			books.POST("", bookHandler.CreateBook)
			books.GET("", bookHandler.GetBooks)
			books.GET("/:id", bookHandler.GetBookByID)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}
	}

	// Keep old routes for backward compatibility
	router.POST("/books", bookHandler.CreateBook)
	router.GET("/books", bookHandler.GetBooks)
	router.GET("/books/:id", bookHandler.GetBookByID)
	router.PUT("/books/:id", bookHandler.UpdateBook)
	router.DELETE("/books/:id", bookHandler.DeleteBook)
}
