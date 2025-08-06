package handlers

import (
	"net/http"

	"example/go1/internal/models"
	"example/go1/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BookHandler struct {
	bookService *services.BookService
	validator   *validator.Validate
}

// NewBookHandler creates a new book handler
func NewBookHandler() *BookHandler {
	return &BookHandler{
		bookService: services.NewBookService(),
		validator:   validator.New(),
	}
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error string `json:"error"`
}

// SuccessResponse represents a success response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// CreateBook creates a new book
// @Summary Create a new book
// @Description Create a new book with title and author
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.CreateBookRequest true "Book data"
// @Success 201 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [post]
func (h *BookHandler) CreateBook(c *gin.Context) {
	var req models.CreateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request payload: " + err.Error(),
		})
		return
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Validation failed: " + err.Error(),
		})
		return
	}

	book, err := h.bookService.CreateBook(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to create book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, SuccessResponse{
		Message: "Book created successfully",
		Data:    book,
	})
}

// GetBooks retrieves all books
// @Summary Get all books
// @Description Get a list of all books with optional pagination
// @Tags books
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Items per page" default(10)
// @Success 200 {object} SuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [get]
func (h *BookHandler) GetBooks(c *gin.Context) {
	books, err := h.bookService.GetAllBooks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to fetch books: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Books retrieved successfully",
		Data:    books,
	})
}

// GetBookByID retrieves a book by ID
// @Summary Get book by ID
// @Description Get a book by its ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [get]
func (h *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	book, err := h.bookService.GetBookByID(id)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Error: "Book not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to fetch book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Book retrieved successfully",
		Data:    book,
	})
}

// UpdateBook updates an existing book
// @Summary Update book
// @Description Update an existing book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Param book body models.UpdateBookRequest true "Updated book data"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [put]
func (h *BookHandler) UpdateBook(c *gin.Context) {
	id := c.Param("id")
	var req models.UpdateBookRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Invalid request payload: " + err.Error(),
		})
		return
	}

	// Validate request
	if err := h.validator.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Error: "Validation failed: " + err.Error(),
		})
		return
	}

	book, err := h.bookService.UpdateBook(id, req)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Error: "Book not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to update book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Book updated successfully",
		Data:    book,
	})
}

// DeleteBook deletes a book by ID
// @Summary Delete book
// @Description Delete a book by ID
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} SuccessResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [delete]
func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := h.bookService.DeleteBook(id)
	if err != nil {
		if err.Error() == "book not found" {
			c.JSON(http.StatusNotFound, ErrorResponse{
				Error: "Book not found",
			})
			return
		}

		c.JSON(http.StatusInternalServerError, ErrorResponse{
			Error: "Failed to delete book: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, SuccessResponse{
		Message: "Book deleted successfully",
	})
}
