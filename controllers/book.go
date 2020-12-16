package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/siwonpawel/bookish-gopher/models"
	bookRepository "github.com/siwonpawel/bookish-gopher/repository/book"
	"github.com/siwonpawel/bookish-gopher/utils"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Controllers struct {
}

func (c Controllers) GetBooks(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		bookRepo := bookRepository.BookRepository{}

		books, err := bookRepo.GetBooks(db)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: "Server error"})
		} else {
			json.NewEncoder(w).Encode(books)
		}
	}
}

func (c Controllers) GetBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		params := mux.Vars(r)
		id, err := strconv.Atoi(params["id"])

		if err != nil {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Wrong identifier, need to be a number"})
			return
		}

		book, err := bookRepository.BookRepository{}.GetBookById(db, id)
		if err != nil {

			if err == sql.ErrNoRows {
				utils.SendError(w, http.StatusNotFound, models.Error{Message: "Not found"})
				return
			}

			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: "Server error"})
			return

		}

		utils.SendSuccess(w, book)
	}
}

func (c Controllers) AddBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int

		json.NewDecoder(r.Body).Decode(&book)
		if book.Author == "" || book.Title == "" || book.Year == "" {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Enter missing fields"})
			return
		}

		bookID, err := bookRepository.BookRepository{}.AddBook(db, book)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: "Server error"})
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controllers) UpdateBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var book models.Book
		var bookID int64

		json.NewDecoder(r.Body).Decode(&book)
		if book.Author == "" || book.Title == "" || book.Year == "" {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Enter missing fields"})
			return
		}

		bookID, err := bookRepository.BookRepository{}.UpdateBook(db, book)

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: "Server error"})
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, bookID)
	}
}

func (c Controllers) RemoveBook(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		bookID, err := strconv.Atoi(params["id"])
		if err != nil {
			utils.SendError(w, http.StatusBadRequest, models.Error{Message: "Wrong identifier, need to be a number"})
		}

		rowsDeleted, err := bookRepository.BookRepository{}.RemoveBook(db, int64(bookID))

		if err != nil {
			utils.SendError(w, http.StatusInternalServerError, models.Error{Message: "Server error"})
		}

		w.Header().Set("Content-Type", "text/plain")
		utils.SendSuccess(w, rowsDeleted)
	}
}
