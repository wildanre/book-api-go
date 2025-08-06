package services

import (
	"errors"
	"strconv"

	"example/go1/internal/database"
	"example/go1/internal/models"
	"example/go1/internal/utils"

	"gorm.io/gorm"
)

type BookService struct {
	db *gorm.DB
}

// NewBookService creates a new book service
func NewBookService() *BookService {
	return &BookService{
		db: database.GetDB(),
	}
}

// CreateBook creates a new book
func (s *BookService) CreateBook(req models.CreateBookRequest) (*models.Book, error) {
	book := models.Book{
		Title:  utils.SanitizeString(req.Title),
		Author: utils.SanitizeString(req.Author),
	}

	if err := s.db.Create(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

// GetAllBooks retrieves all books
func (s *BookService) GetAllBooks() ([]models.Book, error) {
	var books []models.Book
	if err := s.db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

// GetBookByID retrieves a book by ID
func (s *BookService) GetBookByID(id string) (*models.Book, error) {
	var book models.Book

	// Validate ID format first
	if !utils.ValidateID(id) {
		return nil, errors.New("invalid book ID format")
	}

	// Convert string ID to uint
	bookID, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return nil, errors.New("invalid book ID")
	}

	if err := s.db.First(&book, uint(bookID)).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}

	return &book, nil
}

// UpdateBook updates an existing book
func (s *BookService) UpdateBook(id string, req models.UpdateBookRequest) (*models.Book, error) {
	book, err := s.GetBookByID(id)
	if err != nil {
		return nil, err
	}

	book.Title = utils.SanitizeString(req.Title)
	book.Author = utils.SanitizeString(req.Author)

	if err := s.db.Save(book).Error; err != nil {
		return nil, err
	}

	return book, nil
}

// DeleteBook deletes a book by ID
func (s *BookService) DeleteBook(id string) error {
	book, err := s.GetBookByID(id)
	if err != nil {
		return err
	}

	if err := s.db.Delete(book).Error; err != nil {
		return err
	}

	return nil
}
