package routes

import (
	"example/go1/internal/handlers"

	"github.com/gin-gonic/gin"
)

// SetupBookRoutes sets up book-related routes
func SetupBookRoutes(router *gin.Engine) {
	bookHandler := handlers.NewBookHandler()

	// Group routes with /api prefix
	api := router.Group("/api")
	{
		books := api.Group("/books")
		{
			books.POST("", bookHandler.CreateBook)
			books.GET("", bookHandler.GetBooks)
			books.GET("/:id", bookHandler.GetBookByID)
			books.PUT("/:id", bookHandler.UpdateBook)
			books.DELETE("/:id", bookHandler.DeleteBook)
		}
	}
}
