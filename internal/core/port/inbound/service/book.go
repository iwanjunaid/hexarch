package service

import (
	"context"

	"github.com/iwanjunaid/hexarch/internal/core/domain"
)

type BookService interface {
	CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, error)
	GetBookByISBN(ctx context.Context, isbn string) (*domain.Book, error)
}
