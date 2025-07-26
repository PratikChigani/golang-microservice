package models

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/pratikchigani/book-service/internal/db"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

func GetAllBooks() ([]Book, error) {
	log.Println("[DB] Fetching all books")

	rows, err := db.DB.Query(`SELECT id, title, author FROM books`)
	if err != nil {
		return nil, fmt.Errorf("query all books: %w", err)
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var b Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author); err != nil {
			return nil, fmt.Errorf("scan book: %w", err)
		}
		books = append(books, b)
	}

	if books == nil {
		books = make([]Book, 0)
	}

	return books, nil
}

func GetBookByID(id string) (*Book, error) {
	var b Book
	err := db.DB.QueryRow(`SELECT id, title, author FROM books WHERE id = ?`, id).Scan(&b.ID, &b.Title, &b.Author)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, fmt.Errorf("get book by id [%s]: %w", id, err)
	}
	return &b, nil
}

func CreateBook(b Book) error {
	log.Printf("[DB] Attempting to create book: %+v\n", b)

	// Ensure book doesn't already exist
	if _, err := GetBookByID(b.ID); err == nil {
		return ErrConflict
	} else if !errors.Is(err, ErrNotFound) {
		return fmt.Errorf("check existence: %w", err)
	}

	_, err := db.DB.Exec(`INSERT INTO books (id, title, author) VALUES (?, ?, ?)`, b.ID, b.Title, b.Author)
	if err != nil {
		return fmt.Errorf("insert book: %w", err)
	}

	log.Println("[DB] Book created successfully")
	return nil
}

func UpdateBook(id string, b Book) error {
	res, err := db.DB.Exec(`UPDATE books SET title = ?, author = ? WHERE id = ?`, b.Title, b.Author, id)
	if err != nil {
		return fmt.Errorf("update book [%s]: %w", id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected for update: %w", err)
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	log.Printf("[DB] Book [%s] updated", id)
	return nil
}

func DeleteBook(id string) error {
	res, err := db.DB.Exec(`DELETE FROM books WHERE id = ?`, id)
	if err != nil {
		return fmt.Errorf("delete book [%s]: %w", id, err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("rows affected for delete: %w", err)
	}
	if rowsAffected == 0 {
		return ErrNotFound
	}

	log.Printf("[DB] Book [%s] deleted", id)
	return nil
}
