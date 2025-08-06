package models

import (
	"time"

	"gorm.io/gorm"
)

// Book represents a book entity
type Book struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null" validate:"required,min=1,max=255"`
	Author    string         `json:"author" gorm:"not null" validate:"required,min=1,max=255"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// CreateBookRequest represents the request payload for creating a book
type CreateBookRequest struct {
	Title  string `json:"title" validate:"required,min=1,max=255" binding:"required"`
	Author string `json:"author" validate:"required,min=1,max=255" binding:"required"`
}

// UpdateBookRequest represents the request payload for updating a book
type UpdateBookRequest struct {
	Title  string `json:"title" validate:"required,min=1,max=255" binding:"required"`
	Author string `json:"author" validate:"required,min=1,max=255" binding:"required"`
}
