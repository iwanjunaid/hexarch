package handler

import (
	"net/http"

	"gitlab.sicepat.tech/platform/golib/response"
)

func Health(r *http.Request) *response.JSONResponse {
	return response.NewJSONResponse().APIStatusSuccess()
}
