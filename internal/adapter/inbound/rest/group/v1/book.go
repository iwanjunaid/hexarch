package v1

import (
	hv1 "github.com/iwanjunaid/hexarch/internal/adapter/inbound/rest/handler/v1"
	"gitlab.sicepat.tech/platform/golib/router"
)

func NewGroupBookV1(r *router.MyRouter, handler *hv1.Handler) {
	r.POST("/books", handler.Create)
	r.GET("/books/:isbn", handler.GetBookByISBN)
}
