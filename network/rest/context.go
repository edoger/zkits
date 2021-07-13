package rest

import (
	"net/http"
)

type Context interface {
	Request() *http.Request

	IsRPC() bool
	IsREST() bool
}
