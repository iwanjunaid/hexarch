package registry

import (
	"github.com/iwanjunaid/hexarch/internal/core/port/outbound/repository"
)

type RepositoryRegistry interface {
	GetBookRepository() repository.BookRepository
}
