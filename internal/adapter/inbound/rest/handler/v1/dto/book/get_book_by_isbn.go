package book

import (
	"github.com/iwanjunaid/hexarch/internal/core/domain"
)

type GetBookByISBNPresenter struct {
	ID    string `json:"id"`
	ISBN  string `json:"isbn"`
	Title string `json:"title"`
}

type GetBookByISBNDTO struct{}

func (p *GetBookByISBNDTO) TransformOut(book *domain.Book) *GetBookByISBNPresenter {
	return &GetBookByISBNPresenter{
		ID:    book.ID,
		ISBN:  book.ISBN,
		Title: book.Title,
	}
}
