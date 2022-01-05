package service

import (
	"context"

	"github.com/iwanjunaid/hexarch/internal/core/domain"
	"github.com/iwanjunaid/hexarch/internal/core/port/outbound/registry"
)

type BookService struct {
	repositoryRegistry registry.RepositoryRegistry
}

func NewBookService(repositoryRegistry registry.RepositoryRegistry) *BookService {
	return &BookService{
		repositoryRegistry: repositoryRegistry,
	}
}

func (s *BookService) CreateBook(ctx context.Context, book *domain.Book) (*domain.Book, error) {
	repo := s.repositoryRegistry.GetBookRepository()
	book, err := repo.Create(ctx, book)

	if err != nil {
		return nil, err
	}

	return book, nil
}

func (s *BookService) GetBookByISBN(ctx context.Context, isbn string) (*domain.Book, error) {
	repo := s.repositoryRegistry.GetBookRepository()
	book, err := repo.GetBookByISBN(ctx, isbn)

	if err != nil {
		return nil, err
	}

	return book, nil
}
