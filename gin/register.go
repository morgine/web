package gin

import (
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Name       string
	Comment    string
	In         interface{}
	Out        interface{}
	HandleFunc gin.HandlerFunc
	Err        CodeError
}

type CodeError struct {
	Code int
	Err  string
}

type Engine struct {
	Handlers []Handler
}

type Router interface {
	Use(handlers ...*Handler) Router
	Handle(method, path string, handler *Handler)
	HandleFunc(method, path string, handlerFunc gin.HandlerFunc)
}

type filterEngine struct {
	gin.IRouter
}

func (f *filterEngine) Use(handlers ...*Handler) Router {
	for _, handler := range handlers {
		f.IRouter.Use()
	}
}

func (f *filterEngine) Handle(method, path string, handler *Handler) {
	panic("implement me")
}

func (f *filterEngine) HandleFunc(method, path string, handlerFunc gin.HandlerFunc) {
	panic("implement me")
}

func handlerFunctions(handlers []*Handler) []gin.HandlerFunc {

}
