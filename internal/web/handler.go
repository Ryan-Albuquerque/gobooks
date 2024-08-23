package web

import (
	"encoding/json"
	"gobooks/internal/service"
	"net/http"
	"strconv"
)

type BookHandler struct {
	service *service.BookService
}

func NewBookHandler(service *service.BookService) *BookHandler {
	return &BookHandler{service: service}
}

func (h *BookHandler) GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := h.service.GetBooks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func (h *BookHandler) CreateBook(w http.ResponseWriter, r *http.Request) {
	book := service.BookEntity{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.CreateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) UpdateBook(w http.ResponseWriter, r *http.Request) {
	book := service.BookEntity{}
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.service.UpdateBook(&book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func (h *BookHandler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	err := h.service.DeleteBook(convertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func convertToInt(id string) int {
	i, err := strconv.Atoi(id)
	if err != nil {
		return 0
	}
	return i
}

func (h *BookHandler) GetBookById(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	book, err := h.service.GetBookById(convertToInt(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
