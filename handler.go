package web

type Handler interface {
	Handle(ctx *Context) error
}

type HandlerFunc func(ctx *Context) error

func (h HandlerFunc) Handle(ctx *Context) error {
	return h(ctx)
}
