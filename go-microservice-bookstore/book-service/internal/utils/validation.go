package utils

import (
	"errors"
	"strings"

	"github.com/pratikchigani/book-service/internal/models"
)

func ValidateBook(b models.Book) error {
	if strings.TrimSpace(b.ID) == "" {
		return errors.New("book ID is required")
	}
	if strings.TrimSpace(b.Title) == "" {
		return errors.New("book title is required")
	}
	if strings.TrimSpace(b.Author) == "" {
		return errors.New("book author is required")
	}
	return nil
}
