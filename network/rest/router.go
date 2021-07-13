package rest

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Router interface {
	Use(middlewares ...Middleware)

	GET(relative string, handler Handler)
	// HEAD(relative string, handler Handler)
	// POST(relative string, handler Handler)
	// PUT(relative string, handler Handler)
	// PATCH(relative string, handler Handler)
	// DELETE(relative string, handler Handler)
	// CONNECT(relative string, handler Handler)
	// OPTIONS(relative string, handler Handler)
	// TRACE(relative string, handler Handler)

	// Any(relative string, handler Handler)
	//
	// Method(method string, relative string, handler Handler)
	// Methods(methods []string, relative string, handler Handler)

	Group(relative string) Router
}

type iRouter struct {
	path    string
	middlewares []Middleware

	handler     Handler
	handlers    map[string]Handler

	mux *httprouter.Router
}

func (r *iRouter) Use(middlewares ...Middleware) {
	r.mux.GET()
	for i, j := 0, len(middlewares); i < j; i++ {
		r.middlewares = append(r.middlewares, middlewares[i])
	}
}

func (r *iRouter) GET(relative string, handler Handler) {
	// r.relative = relative
	r.handler = handler
}

func (r *iRouter) Group(relative string) Router {
	return &iRouter{
		// relative: r.relative + relative,
	}
}

type Matcher func(r *http.Request) (bool, error)

type iRoute struct {
	matcher []Matcher
}