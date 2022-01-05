package v1

import (
	"errors"
	"net/http"

	dto "github.com/iwanjunaid/hexarch/internal/adapter/inbound/rest/handler/v1/dto/book"
	"gitlab.sicepat.tech/platform/golib/response"
	"gitlab.sicepat.tech/platform/golib/router"
)

func (h *Handler) Create(r *http.Request) *response.JSONResponse {
	var createBookDTO dto.CreateBookDTO

	resp := response.NewJSONResponse()
	book, err := createBookDTO.TransformIn(r)

	if err != nil {
		errMessage := errors.New("bad request")
		resp.SetError(response.ErrBadRequest)
		resp.SetMessage(errMessage.Error())
		resp.SetLog("error", err)
		return resp
	}

	service := h.serviceRegistry.GetBookService()
	newBook, err := service.CreateBook(r.Context(), book)

	if err != nil {
		resp.SetError(response.ErrInternalServerError)
		resp.SetMessage(err.Error())
		resp.SetLog("error", err)

		return resp
	}

	return resp.SetData(createBookDTO.TransformOut(newBook)).APIStatusAccepted()
}

func (h *Handler) GetBookByISBN(r *http.Request) *response.JSONResponse {
	isbn := router.GetHttpParam(r.Context(), "isbn")
	service := h.serviceRegistry.GetBookService()
	resp := response.NewJSONResponse()
	book, err := service.GetBookByISBN(r.Context(), isbn)

	if err != nil {
		resp.SetError(response.ErrInternalServerError)
		resp.SetMessage(err.Error())
		resp.SetLog("error", err)

		return resp
	}

	var getBookByISBNDTO dto.GetBookByISBNDTO

	presBook := getBookByISBNDTO.TransformOut(book)

	return resp.SetData(presBook).APIStatusSuccess()
}
