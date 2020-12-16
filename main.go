package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/subosito/gotenv"

	"github.com/siwonpawel/bookish-gopher/controllers"
	"github.com/siwonpawel/bookish-gopher/driver"
)

var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	gotenv.Load()
}

func main() {

	db = driver.ConnectDB()

	cntlr := controllers.Controllers{}

	router := mux.NewRouter()
	router.HandleFunc("/books", cntlr.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", cntlr.GetBook(db)).Methods("GET")
	router.HandleFunc("/books", cntlr.AddBook(db)).Methods("POST")
	router.HandleFunc("/books", cntlr.UpdateBook(db)).Methods("PUT")
	router.HandleFunc("/books/{id}", cntlr.RemoveBook(db)).Methods("DELETE")

	fmt.Println("Your server is running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
