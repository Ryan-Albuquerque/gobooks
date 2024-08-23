package main

import (
	"database/sql"
	"gobooks/internal/service"
	"gobooks/internal/web"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./books.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	bookService := service.NewBookService(db)

	bookHandler := web.NewBookHandler(bookService)

	router := http.NewServeMux()

	router.HandleFunc("GET /books", bookHandler.GetBooks)
	router.HandleFunc("POST /books", bookHandler.CreateBook)
	router.HandleFunc("PUT /books/{id}", bookHandler.UpdateBook)
	router.HandleFunc("DELETE /books/{id}", bookHandler.DeleteBook)
	router.HandleFunc("GET /books/{id}", bookHandler.GetBookById)

	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
