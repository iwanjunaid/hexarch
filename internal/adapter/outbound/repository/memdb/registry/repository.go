package registry

import (
	"github.com/iwanjunaid/hexarch/internal/adapter/outbound/repository/memdb"
	"github.com/iwanjunaid/hexarch/internal/core/port/outbound/registry"
	"github.com/iwanjunaid/hexarch/internal/core/port/outbound/repository"
)

type RepositoryRegistry struct {
	bookRepository repository.BookRepository
}

func NewRepositoryRegistry() registry.RepositoryRegistry {
	repo := &RepositoryRegistry{
		bookRepository: memdb.NewBookRepository(),
	}

	return repo
}

func (r *RepositoryRegistry) GetBookRepository() repository.BookRepository {
	return r.bookRepository
}
