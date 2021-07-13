package rest


type Handler interface {
	Handle(ctx Context) Response
}

type HandlerFunc func(ctx Context) Response

func (f HandlerFunc) Handle(ctx Context) Response  {
	return f(ctx)
}