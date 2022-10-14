package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/nvanonim/go-bookstore/pkg/models"
	"github.com/nvanonim/go-bookstore/pkg/utils"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := utils.ParseBody(r, &book)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.CreateBook()
	utils.RespondWithJson(w, http.StatusCreated, book)
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	if books == nil {
		utils.RespondWithError(w, http.StatusNotFound, "No books found")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}
	book := models.GetBookById(bookId)
	if book == nil {
		utils.RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}
	utils.RespondWithJson(w, http.StatusOK, book)
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Cek apakah IDnya ada di parameter
	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	// Cek apakah buku dengan ID tsb ada
	book := models.GetBookById(bookId)
	if book == nil {
		utils.RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	// Cek apakah request body valid/cocok dengan structnya
	err = utils.ParseBody(r, &book)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// Update buku
	book.UpdateBook()
	utils.RespondWithJson(w, http.StatusOK, book)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId, err := strconv.ParseInt(vars["bookId"], 0, 0)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid book ID")
		return
	}

	book := models.GetBookById(bookId)
	if book == nil {
		utils.RespondWithError(w, http.StatusNotFound, "Book not found")
		return
	}

	models.DeleteBookById(bookId)
	utils.RespondWithJson(w, http.StatusOK, book)
}
