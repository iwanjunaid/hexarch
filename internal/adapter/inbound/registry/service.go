package registry

import (
	iservice "github.com/iwanjunaid/hexarch/internal/core/port/inbound/service"
	"github.com/iwanjunaid/hexarch/internal/core/port/outbound/registry"
	"github.com/iwanjunaid/hexarch/internal/core/service"
)

type ServiceRegistry struct {
	repositoryRegistry registry.RepositoryRegistry
	bookService        iservice.BookService
}

func NewServiceRegistry(repositoryRegistry registry.RepositoryRegistry) *ServiceRegistry {
	return &ServiceRegistry{
		repositoryRegistry: repositoryRegistry,
		bookService:        service.NewBookService(repositoryRegistry),
	}
}

func (r *ServiceRegistry) GetBookService() iservice.BookService {
	return r.bookService
}
