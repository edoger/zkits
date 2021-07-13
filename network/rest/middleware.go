package rest

type Middleware interface {
	Handle(ctx Context, next func(Context) Response) Response
}

type MiddlewareFunc func(ctx Context, next func(Context) Response) Response

func (f MiddlewareFunc) Handle(ctx Context, next func(Context) Response) Response {
	return f(ctx, next)
}
