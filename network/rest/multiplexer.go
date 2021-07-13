package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Multiplexer interface {
	http.Handler
}

type multiplexer struct {
	path   string
	router httprouter.Router
}

func (m *multiplexer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.router.ServeHTTP(w, r)
}
