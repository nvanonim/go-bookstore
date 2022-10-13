package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nvanonim/go-bookstore/pkg/models"
	"github.com/nvanonim/go-bookstore/pkg/utils"
	"net/http"
	"strconv"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	book.CreateBook()
	utils.RespondWithJson(w, http.StatusCreated, book)
}
