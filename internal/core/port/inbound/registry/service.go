package registry

import (
	"github.com/iwanjunaid/hexarch/internal/core/port/inbound/service"
)

type ServiceRegistry interface {
	GetBookService() service.BookService
}
