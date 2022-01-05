package book

import (
	"encoding/json"
	"net/http"

	"github.com/iwanjunaid/hexarch/internal/core/domain"
)

type CreateBookRequest struct {
	ISBN  string `json:"isbn"`
	Title string `json:"title"`
}

type CreateBookPresenter struct {
	ID    string `json:"id"`
	ISBN  string `json:"isbn"`
	Title string `json:"title"`
}

type CreateBookDTO struct{}

func (r *CreateBookDTO) TransformIn(req *http.Request) (*domain.Book, error) {
	var request CreateBookRequest

	err := json.NewDecoder(req.Body).Decode(&request)

	if err != nil {
		return nil, err
	}

	book := &domain.Book{
		ISBN:  request.ISBN,
		Title: request.Title,
	}

	return book, nil
}

func (p *CreateBookDTO) TransformOut(book *domain.Book) *CreateBookPresenter {
	return &CreateBookPresenter{
		ID:    book.ID,
		ISBN:  book.ISBN,
		Title: book.Title,
	}
}
