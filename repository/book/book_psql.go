package bookRepository

import (
	"database/sql"
	"log"

	"github.com/siwonpawel/bookish-gopher/models"
)

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type BookRepository struct{}

func (c BookRepository) GetBooks(db *sql.DB) ([]models.Book, error) {
	var book models.Book
	var books []models.Book

	rows, err := db.Query("select * from books")
	if err != nil {
		return books, err
	}
	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)
		books = append(books, book)
	}

	if err != nil {
		return []models.Book{}, err
	}

	return books, err
}

func (c BookRepository) GetBookById(db *sql.DB, id int) (models.Book, error) {
	var book models.Book

	rows := db.QueryRow("select * from books where id = $1", id)
	err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Year)

	return book, err
}

func (c BookRepository) AddBook(db *sql.DB, book models.Book) (int, error) {

	var bookID int
	err := db.QueryRow("insert into books (title, author, year) values($1, $2, $3) RETURNING id;", book.Title, book.Author, book.Year).Scan(&bookID)

	return bookID, err
}

func (c BookRepository) UpdateBook(db *sql.DB, book models.Book) (int64, error) {
	result, err := db.Exec("update books set title = $1, author = $2, year = $3 where id = $4", book.Title, book.Author, book.Year, book.ID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil

}

func (c BookRepository) RemoveBook(db *sql.DB, bookID int64) (int64, error) {
	result, err := db.Exec("delete from books where id = $1", bookID)

	if err != nil {
		return 0, err
	}

	rowsUpdated, err := result.RowsAffected()

	if err != nil {
		return 0, err
	}

	return rowsUpdated, nil

}
