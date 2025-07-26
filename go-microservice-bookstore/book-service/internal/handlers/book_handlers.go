package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pratikchigani/book-service/internal/models"
	"github.com/pratikchigani/book-service/internal/utils"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := models.GetAllBooks()
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch books")
		return
	}
	utils.RespondJSON(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	book, err := models.GetBookByID(id)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "Book not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to fetch book")
		}
		return
	}

	utils.RespondJSON(w, http.StatusOK, book)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var newBook models.Book
	if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := utils.ValidateBook(newBook); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err := models.GetBookByID(newBook.ID)
	if err == nil {
		utils.RespondError(w, http.StatusConflict, "Book with this ID already exists")
		return
	} else if !errors.Is(err, models.ErrNotFound) {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to check for existing book")
		return
	}

	if err := models.CreateBook(newBook); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create book")
		return
	}

	utils.RespondJSON(w, http.StatusCreated, newBook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	var updatedBook models.Book

	if err := json.NewDecoder(r.Body).Decode(&updatedBook); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	if err := utils.ValidateBook(updatedBook); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := models.UpdateBook(id, updatedBook); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "Book not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to update book")
		}
		return
	}

	utils.RespondJSON(w, http.StatusOK, updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	if err := models.DeleteBook(id); err != nil {
		if errors.Is(err, models.ErrNotFound) {
			utils.RespondError(w, http.StatusNotFound, "Book not found")
		} else {
			utils.RespondError(w, http.StatusInternalServerError, "Failed to delete book")
		}
		return
	}

	// No body for 204 No Content
	w.WriteHeader(http.StatusNoContent)
}
