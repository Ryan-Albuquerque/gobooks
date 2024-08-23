package service

import (
	"database/sql"
	"gobooks/internal/entity"
)

type BookService struct {
	db *sql.DB
}

func NewBookService(db *sql.DB) *BookService {
	return &BookService{db}
}

type BookEntity = entity.Book

func (s *BookService) CreateBook(book *BookEntity) error {
	query := "INSERT INTO books (title, author, genre) VALUES (?, ?, ?)"
	res, err := s.db.Exec(query, book.Title, book.Author, book.Genre)
	if err != nil {
		return err
	}
	lastInsetredID, err := res.LastInsertId()
	if err != nil {
		return err
	}
	book.ID = int(lastInsetredID)

	return nil
}

func (s *BookService) GetBooks() ([]BookEntity, error) {
	query := "SELECT * FROM books"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []BookEntity{}
	for rows.Next() {
		book := BookEntity{}
		err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (s *BookService) UpdateBook(book *BookEntity) error {
	query := "UPDATE books SET title = ?, author = ?, genre = ? WHERE id = ?"
	_, err := s.db.Exec(query, book.Title, book.Author, book.Genre, book.ID)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) DeleteBook(id int) error {
	query := "DELETE FROM books WHERE id = ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (s *BookService) GetBookById(id int) (BookEntity, error) {
	query := "SELECT * FROM books WHERE id = ?"
	row := s.db.QueryRow(query, id)

	book := BookEntity{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Genre)
	if err != nil {
		return BookEntity{}, err
	}

	return book, nil
}
