package memdb

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/iwanjunaid/hexarch/internal/core/domain"
)

var books map[string]*domain.Book

func init() {
	books = make(map[string]*domain.Book)
}

type BookRepository struct{}

func NewBookRepository() *BookRepository {
	return &BookRepository{}
}

func (r *BookRepository) Create(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	_, ok := books[book.ISBN]

	if ok {
		return nil, errors.New("book already exists")
	}

	id, err := uuid.NewRandom()

	if err != nil {
		return nil, errors.New("generate id failed")
	}

	book.ID = id.String()
	books[book.ISBN] = book

	return book, nil
}

func (r *BookRepository) GetBookByISBN(ctx context.Context, isbn string) (*domain.Book, error) {
	book, ok := books[isbn]

	if !ok {
		return nil, errors.New("book not found")
	}

	return book, nil
}
